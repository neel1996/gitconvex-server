package api

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

func getOs() string {
	return runtime.GOOS
}

func getGitVersion() string {
	gitPath, err := exec.LookPath("git")

	if err != nil {
		logger := global.Logger{Message: fmt.Sprintf("Git cannot be lovated \n %v", err)}
		logger.LogError()
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

	logger := global.Logger{Message: fmt.Sprintf("Obtained host information : %v -- %v", getOs(), getGitVersion())}
	logger.LogInfo()

	return &model.HealthCheckParams{
		Os:  getOs(),
		Git: getGitVersion(),
	}
}
