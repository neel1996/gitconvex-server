package global

import "log"

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
	log.Printf("INFO : %v\n", logger.Message)
}

func (logger *Logger) LogWarning() {
	log.Printf("WARNING: %v\n", logger.Message)
}

func (logger *Logger) LogError() {
	log.Printf("ERROR: %v\n", logger.Message)
}
