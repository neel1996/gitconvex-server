package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
)

func RemoteData(repo *git.Repository) []string {
	logger := global.Logger{}

	remote, _ := repo.Remotes()
	remoteURL := func() []string {
		var rUrl []string
		for _, i := range remote {
			rUrl = append(rUrl, i.Config().URLs...)
		}
		return rUrl
	}()

	logger.Log(fmt.Sprintf("Available remotes in repo : \n%v", remoteURL), global.StatusInfo)
	return remoteURL
}
