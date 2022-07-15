package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	log = logger.Sugar()
}

func Info(msg string, args ...interface{}) {
	log.Infow(msg, args)
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalw(msg, args)
}

func Error(msg string, err error) {
	log.Errorw(msg, "error", err)
}
