package git

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"go/types"
	"io"
)

func PullFromRemote(repo *git.Repository, remoteURL string, remoteBranch string) *model.PullResult {
	var pullErr error
	logger := global.Logger{}
	remoteName := GetRemoteName(repo, remoteURL)
	w, _ := repo.Worktree()
	b := new(bytes.Buffer)

	refName := fmt.Sprintf("refs/heads/%s", remoteBranch)

	ref, refErr := repo.Storer.Reference(plumbing.ReferenceName(refName))

	if refErr != nil {
		fmt.Println(refErr.Error())
		pullErr = types.Error{Msg: "branch reference does not exist"}
	} else {
		pullErr = w.Pull(&git.PullOptions{
			RemoteName:    remoteName,
			ReferenceName: ref.Name(),
			Progress: sideband.Progress(func(f io.Writer) io.Writer {
				return f
			}(b)),
			SingleBranch: true,
		})
	}

	logger.Log(b.String(), global.StatusInfo)

	if pullErr != nil {
		if pullErr.Error() == git.NoErrAlreadyUpToDate.Error() {
			logger.Log(pullErr.Error(), global.StatusWarning)
			msg := "No changes to pull from " + remoteName
			return &model.PullResult{
				Status:      "NEW CHANGES ABSENT",
				PulledItems: []*string{&msg},
			}
		} else {
			logger.Log(pullErr.Error(), global.StatusError)
			return &model.PullResult{
				Status:      "PULL ERROR",
				PulledItems: nil,
			}
		}
	} else {
		logger.Log("New items pulled from remote", global.StatusInfo)
		msg := "New Items Pulled from remote " + remoteName
		return &model.PullResult{
			Status:      "PULL SUCCESS",
			PulledItems: []*string{&msg},
		}
	}
}
