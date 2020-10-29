package git

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"io"
)

func FetchFromRemote(repo *git.Repository, remoteURL string, remoteBranch string) *model.FetchResult {
	remoteName := GetRemoteName(repo, remoteURL)
	logger := global.Logger{}

	targetRefPsec := "refs/heads/" + remoteBranch + ":refs/remotes/" + remoteBranch
	b := new(bytes.Buffer)
	var fetchErr error

	if remoteURL != "" && remoteBranch != "" {
		fetchErr = repo.Fetch(&git.FetchOptions{
			RemoteName: remoteName,
			RefSpecs:   []config.RefSpec{config.RefSpec(targetRefPsec)},
			Progress: sideband.Progress(func(f io.Writer) io.Writer {
				return f
			}(b)),
		})
	} else {
		fetchErr = repo.Fetch(&git.FetchOptions{
			RemoteName: git.DefaultRemoteName,
			Progress: sideband.Progress(func(f io.Writer) io.Writer {
				return f
			}(b)),
		})
	}

	if fetchErr != nil {
		if fetchErr.Error() == "already up-to-date" {
			logger.Log(fetchErr.Error(), global.StatusWarning)
			return &model.FetchResult{
				Status:       "NEW CHANGES ABSENT",
				FetchedItems: nil,
			}
		} else {
			logger.Log(fetchErr.Error(), global.StatusError)
			return &model.FetchResult{
				Status:       "FETCH ERROR",
				FetchedItems: nil,
			}
		}

	} else {
		logger.Log(b.String(), global.StatusInfo)
		msg := fmt.Sprintf("Changes fetched from %v", git.DefaultRemoteName)
		return &model.FetchResult{
			Status:       "CHANGES FETCHED FROM REMOTE",
			FetchedItems: []*string{&msg},
		}
	}

}