package global

import ("log"
"github.com/TwinProduction/go-color"
)

type LoggerInterface interface {
	LogInfo()
	LogWarning()
	LogError()
}

type Logger struct {
	Message string
}

const (
	StatusError   = "StatusError"
	StatusInfo    = "StatusInfo"
	StatusWarning = "StatusWarning"
)

func (logger *Logger) Log(message string, status string) {
	logger.Message = message
	switch status {
		case StatusInfo:
			logger.LogInfo()
			break
		case StatusWarning:
			logger.LogWarning()
			break
		case StatusError:
			logger.LogError()
			break
		default:
			logger.LogInfo()
	}
}

func (logger *Logger) LogInfo() {
	log.Printf("%vINFO: %v%v\n", color.Cyan, color.Reset, logger.Message)
}

func (logger *Logger) LogWarning() {
	log.Printf("%vWARNING: %v%v\n", color.Yellow, color.Reset, logger.Message)
}

func (logger *Logger) LogError() {
	log.Printf("%vERROR: %v%v\n", color.Red, color.Reset, logger.Message)
}
