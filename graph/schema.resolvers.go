package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/neel1996/gitconvex-server/api"
	"github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/graph/generated"
	"github.com/neel1996/gitconvex-server/graph/model"
)

func (r *mutationResolver) AddRepo(ctx context.Context, repoName string, repoPath string, cloneSwitch bool, repoURL *string, initSwitch bool, authOption string, userName *string, password *string) (*model.AddRepoParams, error) {
	return api.AddRepo(model.NewRepoInputs{
		RepoName:    repoName,
		RepoPath:    repoPath,
		CloneSwitch: cloneSwitch,
		RepoURL:     repoURL,
		InitSwitch:  initSwitch,
		AuthOption:  authOption,
		UserName:    userName,
		Password:    password,
	}), nil
}

func (r *mutationResolver) AddBranch(ctx context.Context, repoID string, branchName string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.AddBranch(repo.GitRepo, branchName), nil
}

func (r *mutationResolver) CheckoutBranch(ctx context.Context, repoID string, branchName string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.CheckoutBranch(repo.GitRepo, branchName), nil
}

func (r *mutationResolver) DeleteBranch(ctx context.Context, repoID string, branchName string, forceFlag bool) (*model.BranchDeleteStatus, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.DeleteBranch(repo.GitRepo, branchName, forceFlag), nil
}

func (r *mutationResolver) AddRemote(ctx context.Context, repoID string, remoteName string, remoteURL string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.AddRemote(repo.GitRepo, remoteName, remoteURL), nil
}

func (r *mutationResolver) FetchFromRemote(ctx context.Context, repoID string, remoteURL *string, remoteBranch *string) (*model.FetchResult, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.FetchFromRemote(repo.GitRepo, *remoteURL, *remoteBranch), nil
}

func (r *mutationResolver) PullFromRemote(ctx context.Context, repoID string, remoteURL *string, remoteBranch *string) (*model.PullResult, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.PullFromRemote(repo.GitRepo, *remoteURL, *remoteBranch), nil
}

func (r *mutationResolver) StageItem(ctx context.Context, repoID string, item string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.StageItem(repo.GitRepo, item), nil
}

func (r *mutationResolver) RemoveStagedItem(ctx context.Context, repoID string, item string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan
	return git.RemoveItem(repo.RepoPath, item), nil
}

func (r *mutationResolver) StageAllItems(ctx context.Context, repoID string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.StageAllItems(repo.GitRepo), nil
}

func (r *mutationResolver) CommitChanges(ctx context.Context, repoID string, commitMessage string) (string, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.CommitChanges(repo.GitRepo, commitMessage), nil
}

func (r *queryResolver) HealthCheck(ctx context.Context) (*model.HealthCheckParams, error) {
	return api.HealthCheckApi(), nil
}

func (r *queryResolver) FetchRepo(ctx context.Context) (*model.FetchRepoParams, error) {
	return api.FetchRepo(), nil
}

func (r *queryResolver) GitRepoStatus(ctx context.Context, repoID string) (*model.GitRepoStatusResults, error) {
	return api.RepoStatus(repoID), nil
}

func (r *queryResolver) GitFolderContent(ctx context.Context, repoID string, directoryName *string) (*model.GitFolderContentResults, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	tmp := &model.GitFolderContentResults{}
	tmp.FileBasedCommits = nil
	tmp.TrackedFiles = nil

	return git.ListFiles(repo.GitRepo, repo.RepoPath, *directoryName), nil
}

func (r *queryResolver) GitCommitLogs(ctx context.Context, repoID string, skipLimit int) (*model.GitCommitLogResults, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.CommitLogs(repo.GitRepo, skipLimit), nil
}

func (r *queryResolver) GitCommitFiles(ctx context.Context, repoID string, commitHash string) ([]*model.GitCommitFileResult, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.CommitFileList(repo.GitRepo, commitHash), nil
}

func (r *queryResolver) SearchCommitLogs(ctx context.Context, repoID string, searchType string, searchKey string) ([]*model.GitCommits, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.SearchCommitLogs(repo.GitRepo, searchType, searchKey), nil
}

func (r *queryResolver) CodeFileDetails(ctx context.Context, repoID string, fileName string) (*model.CodeFileType, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return api.CodeFileView(repo.GitRepo, repo.RepoPath, fileName), nil
}

func (r *queryResolver) GitChanges(ctx context.Context, repoID string) (*model.GitChangeResults, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	return git.ChangedFiles(repo.GitRepo), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
