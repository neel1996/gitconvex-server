package utils

import (
	"encoding/json"
	"github.com/neel1996/gitconvex-server/global"
	"io/ioutil"
	"os"
)

type EnvConfig struct {
	DataBaseFile string `json:"databaseFile"`
	Port         string `json:"port"`
}

func localLogger(message string, status string) {
	logger := &global.Logger{}
	logger.Log(message, status)
}

func EnvConfigValidator() error {
	cwd, _ := os.Getwd()
	fileString := cwd + "/env_config.json"
	_, openErr := os.Open(fileString)

	if openErr != nil {
		localLogger(openErr.Error(), global.StatusError)
		return openErr
	}
	return nil
}

func EnvConfigFileReader() *EnvConfig {
	cwd, _ := os.Getwd()
	fileString := cwd + "/env_config.json"
	envFile, err := os.Open(fileString)

	var envConfig *EnvConfig

	if err != nil {
		localLogger(err.Error(), global.StatusError)
		return nil
	} else {
		if fileContent, openErr := ioutil.ReadAll(envFile); openErr == nil {
			unMarshallErr := json.Unmarshal(fileContent, &envConfig)
			if unMarshallErr == nil {
				return envConfig
			} else {
				localLogger(unMarshallErr.Error(), global.StatusError)
				return nil
			}
		}
	}
	return nil
}

func EnvConfigFileGenerator() error {
	cwd, _ := os.Getwd()
	fileString := cwd + "/env_config.json"

	envContent, _ := json.MarshalIndent(&EnvConfig{
		DataBaseFile: cwd + "/database/repo_datastore.json",
		Port:         "9001",
	}, "", " ")

	return ioutil.WriteFile(fileString, envContent, 0755)
}
