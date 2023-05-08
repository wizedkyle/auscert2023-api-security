package utils

import "go.uber.org/zap"

var Logger *zap.Logger

// GenerateLogger
// Creates an instance of Zap logger
func GenerateLogger() {
	l, err := zap.NewProduction()
	if err != nil {
		return
	}
	Logger = l
}

func WriteErrorLog(err *Error) {
	Logger.Error(err.ExternalError.Message,
		zap.Error(err.InternalError),
		zap.String("id", err.ExternalError.Id),
		zap.String("message", err.ExternalError.Message),
		zap.Int("code", err.ExternalError.Code),
		zap.String("transactionId", err.ExternalError.TransactionId))
}
