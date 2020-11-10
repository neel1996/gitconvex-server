package api

import (
	"encoding/json"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"github.com/neel1996/gitconvex-server/utils"
	"io/ioutil"
	"os"
)

// GetSettingsData returns the data extracted from the env_config.json file
func GetSettingsData() *model.SettingsDataResults {
	envData := utils.EnvConfigFileReader()

	if envData != nil {
		return &model.SettingsDataResults{
			SettingsDatabasePath: envData.DataBaseFile,
			SettingsPortDetails:  envData.Port,
		}
	} else {
		return &model.SettingsDataResults{}
	}
}

// UpdatePortNumber updates the port number in the env_config.json file with the
// newly supplied port number
func UpdatePortNumber(newPort string) string {
	logger := global.Logger{}
	if utils.EnvConfigValidator() == nil {
		var newEnvData utils.EnvConfig
		envData := utils.EnvConfigFileReader()

		newEnvData.Port = newPort
		newEnvData.DataBaseFile = envData.DataBaseFile

		cwd, _ := os.Getwd()
		fileString := cwd + "/env_config.json"
		envContent, _ := json.MarshalIndent(&newEnvData, "", " ")
		writeErr := ioutil.WriteFile(fileString, envContent, 0755)

		if writeErr != nil {
			logger.Log(writeErr.Error(), global.StatusError)
			return "PORT_UPDATE_FAILED"
		} else {
			return "PORT_UPDATED"
		}
	} else {
		return "PORT_UPDATE_FAILED"
	}
}

func UpdateDBFilePath(newFilePath string) string {
	logger := global.Logger{}
	if utils.EnvConfigValidator() == nil {
		var newEnvData utils.EnvConfig
		envData := utils.EnvConfigFileReader()

		newEnvData.Port = envData.Port
		newEnvData.DataBaseFile = newFilePath

		cwd, _ := os.Getwd()
		fileString := cwd + "/env_config.json"
		envContent, _ := json.MarshalIndent(&newEnvData, "", " ")
		writeErr := ioutil.WriteFile(fileString, envContent, 0755)

		if writeErr != nil {
			logger.Log(writeErr.Error(), global.StatusError)
			return "DATAFILE_UPDATE_FAILED"
		} else {
			return "DATAFILE_UPDATE_SUCCESS"
		}
	} else {
		return "DATAFILE_UPDATE_FAILED"
	}
}
