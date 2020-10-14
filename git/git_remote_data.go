package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
)

func RemoteData(repo *git.Repository) []string {
	var remoteList []string
	logger := global.Logger{}

	remote, _ := repo.Remotes()
	for _, r := range remote {
		remoteList = append(remoteList, r.String())
	}
	logger.Log(fmt.Sprintf("Available remotes in repo : \n%v\n", remoteList), global.StatusInfo)

	return remoteList
}
