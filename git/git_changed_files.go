package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func ChangedFiles(repo *git.Repository) {
	head, _ := repo.Head()
	hash := head.Hash()

	fmt.Println(hash)

	commit, commitErr := repo.CommitObject(hash)

	w, _ := repo.Worktree()
	stat, _ := w.Status()

	if commitErr != nil {
		fmt.Println(commitErr.Error())
	} else {
		fileItr, _ := commit.Files()

		_ = fileItr.ForEach(func(file *object.File) error {
			fStat := stat.File(file.Name)
			fmt.Printf("%s - %+v\n", file.Name, string(fStat.Worktree))
			return nil
		})
	}
}
