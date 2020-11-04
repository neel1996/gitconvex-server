package api

import (
	"bufio"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"os"
)

func CodeFileView(repo *git.Repository, targetFile string) *model.CodeFileType {
	var codeLines []*string
	var fileCommit string
	fileCommit = ""

	logger := global.Logger{}
	file, err := os.Open(targetFile)

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			codeLines = append(codeLines, &line)
		}
	}

	commitLog, commitErr := repo.Log(&git.LogOptions{
		From:     plumbing.Hash{},
		Order:    git.LogOrderDFSPost,
		FileName: &targetFile,
		All:      false,
	})

	if commitErr != nil {
		logger.Log(commitErr.Error(), global.StatusError)
	} else {
		nxt, _ := commitLog.Next()
		fileCommit = nxt.Message
	}

	return &model.CodeFileType{
		FileCommit: fileCommit,
		FileData:   codeLines,
	}
}
