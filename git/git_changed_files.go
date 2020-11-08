package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

// ChangedFiles returns the list of changes from the target
// The function organizes the tracked, untracked and staged files in separate slices and returns the struct *model.GitChangeResults

func ChangedFiles(repo *git.Repository) *model.GitChangeResults {
	UnPushedCommits(repo)

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

	w, _ := repo.Worktree()
	stat, _ := w.Status()

	statLines := strings.Split(stat.String(), "\n")

	var statusIndicator string
	var filePath string

	for _, statEntry := range statLines {
		var stagedFile string
		if len(statEntry) == 0 {
			continue
		}
		statEntry := strings.TrimSpace(statEntry)

		if strings.Contains(statEntry, " ") {
			splitEntry := strings.Split(statEntry, " ")
			statusIndicator = splitEntry[0]
			filePath = strings.TrimSpace(strings.Join(splitEntry[1:], " "))

			stagedStat := stat.File(filePath).Staging
			if string(stagedStat) == "M" {
				stagedFile = filePath
				stagedFiles = append(stagedFiles, &stagedFile)
			}

			switch statusIndicator {
			case "?", "??":
				logger.Log(fmt.Sprintf("Untracked entry -> %v", filePath), global.StatusInfo)
				changeStr := filePath
				unTrackedFiles = append(unTrackedFiles, &changeStr)
				break
			case "M":
				logger.Log(fmt.Sprintf("Modified entry - %s -> %v", statusIndicator, filePath), global.StatusInfo)
				if filePath != stagedFile {
					changeStr := "M," + filePath
					modifiedFiles = append(modifiedFiles, &changeStr)
				}
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

	// Loop to iterate and append untracked and deleted staged files to staged item list
	for _, entry := range newStagedItems {
		stagedFiles = append(stagedFiles, &entry)
	}

	// Conditional logic to remove duplicate items from staged files list
	var refMap = make(map[string]bool)
	var refinedStagedFiles []*string

	for _, entry := range stagedFiles {
		if mapEntry := refMap[*entry]; !mapEntry {
			refMap[*entry] = true
			refinedStagedFiles = append(refinedStagedFiles, entry)
		}
	}

	return &model.GitChangeResults{
		GitUntrackedFiles: unTrackedFiles,
		GitChangedFiles:   modifiedFiles,
		GitStagedFiles:    refinedStagedFiles,
	}
}
