package git

import (
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

type FetchInterface interface {
	FetchFromRemote() *model.FetchResult
}

type FetchStruct struct {
	Repo         *git2go.Repository
	RemoteName   string
	RepoPath     string
	RemoteURL    string
	RemoteBranch string
	RepoName     string
	AuthOption   string
	UserName     string
	Password     string
	SSHKeyPath   string
}

// FetchFromRemote performs a git fetch for the supplied remote and branch (e.g. `git fetch origin main`)
// If the remoteBranch is empty, then a fetch is performed with no branch name (similar to `git fetch`)
func (f FetchStruct) FetchFromRemote() *model.FetchResult {
	var targetRefPsec string
	repo := f.Repo
	remoteURL := f.RemoteURL
	remoteBranch := f.RemoteBranch

	// Pick the first available remote if no remote is selected
	if remoteURL == "" {
		remotes, _ := repo.Remotes.List()
		if len(remotes) > 0 {
			firstRemote, _ := repo.Remotes.Lookup(remotes[0])
			remoteURL = firstRemote.Url()
		}
	}

	var remoteDataObject RemoteDataInterface
	remoteDataObject = RemoteDataStruct{
		Repo:      repo,
		RemoteURL: remoteURL,
	}

	remoteName := remoteDataObject.GetRemoteName()
	if remoteBranch == "" {
		targetRefPsec = ""
	} else {
		targetRefPsec = "refs/heads/" + remoteBranch
	}
	targetRemote, _ := repo.Remotes.Lookup(remoteName)

	if targetRemote == nil {
		logger.Log("Target remote is unavailable", global.StatusError)
		return &model.FetchResult{
			Status:       global.FetchFromRemoteError,
			FetchedItems: nil,
		}
	}

	var remoteCallbackObject RemoteCallbackInterface
	remoteCallbackObject = &RemoteCallbackStruct{
		RepoName:   f.RepoName,
		UserName:   f.UserName,
		Password:   f.Password,
		SSHKeyPath: f.SSHKeyPath,
		AuthOption: f.AuthOption,
	}

	fetchOption := &git2go.FetchOptions{
		RemoteCallbacks: remoteCallbackObject.RemoteCallbackSelector(),
	}

	err := targetRemote.Fetch([]string{targetRefPsec}, fetchOption, "")

	if err != nil {
		logger.Log("Fetch failed - "+err.Error(), global.StatusError)
		return &model.FetchResult{
			Status:       global.FetchFromRemoteError,
			FetchedItems: nil,
		}
	} else {
		msg := "Changes fetched from remote " + remoteName
		logger.Log(msg, global.StatusInfo)
		return &model.FetchResult{
			Status:       global.FetchFromRemoteSuccess,
			FetchedItems: []*string{&msg},
		}
	}
}
