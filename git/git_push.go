package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/utils"
)

type PushInterface interface {
	PushToRemote() string
	windowsPush() string
}

type PushStruct struct {
	Repo         *git2go.Repository
	RemoteName   string
	RemoteBranch string
	RepoPath     string
}

// windowsPush is used for pushing changes using the git client if the platform is windows
// go-git push fails in windows due to SSH authentication error
func (p PushStruct) windowsPush() string {
	repoPath := p.RepoPath
	remoteName := p.RemoteName
	branch := p.RemoteBranch

	args := []string{"push", "-u", remoteName, branch}
	cmd := utils.GetGitClient(repoPath, args)
	cmdStr, cmdErr := cmd.Output()

	if cmdErr != nil {
		logger.Log(fmt.Sprintf("Push failed -> %s", cmdErr.Error()), global.StatusError)
		return global.PushToRemoteError
	} else {
		logger.Log(fmt.Sprintf("Changes pushed to remote -> %s", cmdStr), global.StatusInfo)
		return global.PushToRemoteSuccess
	}
}

// PushToRemote pushed the commits to the selected remote repository
// By default it will choose the current branch and pushes to the matching remote branch
func (p PushStruct) PushToRemote() string {
	repo := p.Repo
	remoteBranch := p.RemoteBranch
	remoteName := p.RemoteName

	targetRefPsec := "refs/heads/" + remoteBranch

	remote, remoteErr := repo.Remotes.Lookup(remoteName)
	if remoteErr != nil {
		logger.Log(remoteErr.Error(), global.StatusError)
		return global.PushToRemoteError
	}

	remoteCallbacks := git2go.RemoteCallbacks{
		CertificateCheckCallback: git2go.CertificateCheckCallback(func(cert *git2go.Certificate, valid bool, hostname string) git2go.ErrorCode {
			return 0
		}),
		CredentialsCallback: git2go.CredentialsCallback(func(url string, username_from_url string, allowed_types git2go.CredentialType) (*git2go.Credential, error) {
			return git2go.NewCredentialSSHKeyFromAgent(username_from_url)
		}),
	}

	pushOption := &git2go.PushOptions{
		RemoteCallbacks: remoteCallbacks,
	}

	err := remote.Push([]string{targetRefPsec}, pushOption)

	if err != nil {
		logger.Log(fmt.Sprintf("Error occurred while pushing changes to -> %s : %s\n%s", remoteName, targetRefPsec, err.Error()), global.StatusError)
		return global.PushToRemoteError
	} else {
		logger.Log(fmt.Sprintf("commits pushed to remote -> %s : %s", remoteName, targetRefPsec), global.StatusInfo)
		return global.PushToRemoteSuccess
	}
}
