package api

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/neel1996/gitconvex-server/graph/model"
)

func getOs() string {
	return runtime.GOOS
}

func getGitVersion() string {
	gitPath, err := exec.LookPath("git")

	if err != nil {
		log.Printf("Git cannot be lovated \n %v", err)
		panic(err)
	}

	gitCmd := &exec.Cmd{
		Path: gitPath,
		Args: []string{gitPath, "version"},
	}

	gitVersion, err := gitCmd.Output()

	if err != nil {
		fmt.Printf("Git version could not be obtained \n %v", err)
	}

	return strings.Split(string(gitVersion), "\n")[0]
}

func HealthCheckApi() *model.HealthCheckParams {
	return &model.HealthCheckParams{
		Os:  getOs(),
		Git: getGitVersion(),
	}
}
