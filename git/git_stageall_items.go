package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
	"strings"
)

func StageAllItems(repo *git.Repository) string {
	var stageCounter int
	stageCounter = 0
	w, wErr := repo.Worktree()
	logger.Log(fmt.Sprintf("Staging all changes from repo"), global.StatusInfo)

	if wErr != nil {
		logger.Log(fmt.Sprintf("Error occurred while fetching worktree -> %s", wErr.Error()), global.StatusError)
		return "ALL_STAGE_FAILED"
	} else {
		status, statErr := w.Status()
		if statErr != nil {
			logger.Log(fmt.Sprintf("Error occurred while fetching repo status -> %s", statErr.Error()), global.StatusError)
		} else {
			statLines := strings.Split(status.String(), "\n")

			for _, statEntry := range statLines {
				if len(statEntry) == 0 {
					continue
				}
				statEntry := strings.TrimSpace(statEntry)
				if strings.Contains(statEntry, " ") {
					splitEntry := strings.Split(statEntry, " ")
					statusIndicator := splitEntry[0]
					fileItem := strings.TrimSpace(strings.Join(splitEntry[1:], " "))

					if statusIndicator == "A" {
						continue
					}

					itemHash, addErr := w.Add(fileItem)
					if addErr != nil {
						logger.Log(fmt.Sprintf("Error occurred while staging -> %s --> %s", fileItem, addErr.Error()), global.StatusError)
					} else {
						stageCounter++
						logger.Log(fmt.Sprintf("New item -> %s added to the repo worktree --> %s", fileItem, itemHash.String()), global.StatusInfo)
					}

				} else {
					logger.Log(fmt.Sprintf("Status indicator cannot be obtained for -> %s", statEntry), global.StatusError)
					break
				}
			}
		}
	}

	if stageCounter > 0 {
		return "ALL_STAGED"
	} else {
		return "ALL_STAGE_FAILED"
	}
}
