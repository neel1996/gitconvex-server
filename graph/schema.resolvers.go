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

func (r *mutationResolver) AddRepo(ctx context.Context, repoName string, repoPath string, cloneSwitch bool, repoURL *string, initSwitch bool) (*model.AddRepoParams, error) {
	return api.AddRepo(repoName, repoPath, cloneSwitch, repoURL, initSwitch), nil
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

func (r *queryResolver) HealthCheck(ctx context.Context) (*model.HealthCheckParams, error) {
	return api.HealthCheckApi(), nil
}

func (r *queryResolver) FetchRepo(ctx context.Context) (*model.FetchRepoParams, error) {
	return api.FetchRepo(), nil
}

func (r *queryResolver) GitRepoStatus(ctx context.Context, repoID string) (*model.GitRepoStatusResults, error) {
	return api.RepoStatus(repoID), nil
}

func (r *queryResolver) GitFolderContent(ctx context.Context, repoID string) (*model.GitFolderContentResults, error) {
	repoChan := make(chan git.RepoDetails)
	go git.Repo(repoID, repoChan)
	repo := <-repoChan

	tmp := &model.GitFolderContentResults{}
	tmp.FileBasedCommits = nil
	tmp.TrackedFiles = nil

	return git.ListFiles(repo.GitRepo, repo.RepoPath), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
