package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
	"time"
)

func CommitOrganizer(commits []object.Commit) []*model.GitCommits {
	logger := global.Logger{}
	var commitList []*model.GitCommits
	for _, commit := range commits {
		if !commit.Hash.IsZero() {
			commitHash := commit.Hash.String()
			commitAuthor := strings.Split(commit.Author.String(), " ")[0]
			commitMessage := strings.Split(commit.Message, "\n")[0]
			commitFilesItr, err := commit.Files()
			commitFileCount := 0
			commitDate := ""
			commitRelativeTime := ""

			for _, cString := range strings.Split(commit.String(), "\n") {
				if strings.Contains(cString, "Date:") {
					str := strings.Split(cString, "Date:")[1]
					tempDate := strings.TrimSpace(str)

					if strings.Contains(tempDate, "+") {
						tempDate = strings.TrimSpace(strings.Split(tempDate, "+")[0])
					}

					cTime, convErr := time.Parse(time.ANSIC, tempDate)
					if convErr != nil {
						logger.Log(convErr.Error(), global.StatusError)
					} else {
						commitDate = cTime.String()
						t := time.Now()
						commitRelativeTime = strings.Split(t.Sub(cTime).String(), ".")[0]
					}
				}
			}

			if err != nil {
				logger.Log(err.Error(), global.StatusError)
			} else {
				_ = commitFilesItr.ForEach(func(file *object.File) error {

					return nil
				})
			}

			// Shortening commit hash
			commitHash = commitHash[0:7]

			commitList = append(commitList, &model.GitCommits{
				Hash:               &commitHash,
				Author:             &commitAuthor,
				CommitTime:         &commitDate,
				CommitMessage:      &commitMessage,
				CommitRelativeTime: &commitRelativeTime,
				CommitFilesCount:   &commitFileCount,
			})
		}
	}

	return commitList
}

func CommitLogs(repo *git.Repository, skipCount int) *model.GitCommitLogResults {
	var commitLogs []object.Commit
	allCommitChan := make(chan AllCommitData)
	go AllCommits(repo, allCommitChan)
	acc := <-allCommitChan
	totalCommits := acc.TotalCommits

	head, _ := repo.Head()
	commitItr, commitErr := repo.Log(&git.LogOptions{
		From:  head.Hash(),
		Order: git.LogOrderCommitterTime,
		All:   false,
	})

	if commitErr == nil {
		_ = commitItr.ForEach(func(commit *object.Commit) error {
			commitLogs = append(commitLogs, *commit)
			return nil
		})
	}

	commitLimit := skipCount + 10
	refinedCommits := CommitOrganizer(commitLogs[skipCount:commitLimit])
	return &model.GitCommitLogResults{
		TotalCommits: &totalCommits,
		Commits:      refinedCommits,
	}
}
