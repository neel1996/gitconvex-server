package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

// ChangedFiles returns the list of changes from the target
// The function organizes the tracked, untracked and staged files in separate slices and returns the struct *model.GitChangeResults

func ChangedFiles(repo *git.Repository) *model.GitChangeResults {
	var hash plumbing.Hash
	var stagedFiles []*string
	var unTrackedFiles []*string
	var modifiedFiles []*string
	var newStagedItems []string

	logger := global.Logger{}
	head, headErr := repo.Head()

	if headErr != nil {
		logger.Log(headErr.Error(), global.StatusError)
	} else {
		hash = head.Hash()
	}

	logger.Log(fmt.Sprintf("Fetching latest commit object for -> %s", hash), global.StatusInfo)

	commit, commitErr := repo.CommitObject(hash)
	w, _ := repo.Worktree()
	stat, _ := w.Status()

	statLines := strings.Split(stat.String(), "\n")

	var statusIndicator string
	var filePath string

	if commitErr != nil {
		logger.Log(fmt.Sprintf("Unable to fetch commit at HEAD for %s --> %s", hash, commitErr.Error()), global.StatusError)
	} else {
		fileItr, _ := commit.Files()

		_ = fileItr.ForEach(func(file *object.File) error {
			stagedStat := string(stat.File(file.Name).Staging)

			if stagedStat == "M" {
				logger.Log(fmt.Sprintf("Staged entry -> %v", file.Name), global.StatusInfo)
				stagedFiles = append(stagedFiles, &file.Name)
			}
			return nil
		})
	}

	for _, statEntry := range statLines {
		if len(statEntry) == 0 {
			continue
		}
		statEntry := strings.TrimSpace(statEntry)

		if strings.Contains(statEntry, " ") {
			splitEntry := strings.Split(statEntry, " ")
			statusIndicator = splitEntry[0]
			filePath = strings.TrimSpace(strings.Join(splitEntry[1:], " "))

			// Conditional logic to filter out staged entries from modified file list
			var isStaged bool
			for _, stagedItem := range stagedFiles {
				if *stagedItem == filePath {
					isStaged = true
					break
				}
			}

			if isStaged {
				continue
			}

			switch statusIndicator {
			case "?", "??":
				logger.Log(fmt.Sprintf("Untracked entry -> %v", filePath), global.StatusInfo)
				changeStr := filePath
				unTrackedFiles = append(unTrackedFiles, &changeStr)
				break
			case "M":
				logger.Log(fmt.Sprintf("Modified entry - %s -> %v", statusIndicator, filePath), global.StatusInfo)
				changeStr := "M," + filePath
				modifiedFiles = append(modifiedFiles, &changeStr)
				break
			case "D":
				logger.Log(fmt.Sprintf("Removed entry - %s -> %v", statusIndicator, filePath), global.StatusInfo)
				changeStr := "D," + filePath
				modifiedFiles = append(modifiedFiles, &changeStr)
				break
			case "A":
				logger.Log(fmt.Sprintf("New Staged entry - %s -> %v", statusIndicator, filePath), global.StatusInfo)
				newStagedItems = append(newStagedItems, filePath)
				break
			case "AM":
				logger.Log(fmt.Sprintf("New Staged entry - %s -> %v", statusIndicator, filePath), global.StatusInfo)
				newStagedItems = append(newStagedItems, filePath)
				changeStr := "M," + filePath
				modifiedFiles = append(modifiedFiles, &changeStr)
				break
			}
		} else {
			logger.Log(fmt.Sprintf("Status indicator cannot be obtained for -> %s", statEntry), global.StatusError)
			break
		}
	}

	// Loop to iterate and append untracked staged files to staged item list
	for _, entry := range newStagedItems {
		stagedFiles = append(stagedFiles, &entry)
	}

	return &model.GitChangeResults{
		GitUntrackedFiles: unTrackedFiles,
		GitChangedFiles:   modifiedFiles,
		GitStagedFiles:    stagedFiles,
	}
}
