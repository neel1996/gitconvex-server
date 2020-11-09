package git

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/neel1996/gitconvex-server/global"
	"io"
)

func PushToRemote(repo *git.Repository, remoteName string, remoteBranch string) string {
	targetRefPsec := "refs/remotes/" + remoteName + "/" + remoteBranch + ":refs/heads/" + remoteBranch
	b := new(bytes.Buffer)
	sshAuth, _ := ssh.NewSSHAgentAuth("git")
	logger.Log(fmt.Sprintf("Pushing changes to remote -> %s : %s", remoteName, targetRefPsec), global.StatusInfo)

	err := repo.Push(&git.PushOptions{
		RemoteName: remoteName,
		Auth:       sshAuth,
		Progress: sideband.Progress(func(f io.Writer) io.Writer {
			return f
		}(b)),
	})

	if err != nil {
		logger.Log(fmt.Sprintf("Error occurred while pushing changes to -> %s : %s\n%s", remoteName, targetRefPsec, err.Error()), global.StatusError)
		return "PUSH_FAILED"
	} else {
		logger.Log(fmt.Sprintf("commits pushed to remote -> %s : %s\n%v", remoteName, targetRefPsec, b.String()), global.StatusInfo)
		return "PUSH_SUCCESS"
	}
}
