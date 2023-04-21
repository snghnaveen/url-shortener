package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func SetupLogger() {
	// Create a development logger
	devConfig := zap.NewDevelopmentConfig()
	devConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	devLogger, err := devConfig.Build()
	if err != nil {
		panic("Error creating development logger")
	}

	// Create a production logger
	prodConfig := zap.NewProductionConfig()
	prodConfig.EncoderConfig.TimeKey = "timestamp"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodLogger, err := prodConfig.Build()
	if err != nil {
		panic("Error creating production logger")
	}

	// Set the logger to use in the app
	if IsProdMode() {
		logger = prodLogger
	} else {
		logger = devLogger
	}

	// Use the logger to write some logs
	logger.Info("Starting up the app")

}

// CloseLoggerOnAppExit closes the loggers when the app exits
func CloseLoggerOnAppExit() {
	logger.Sync()
}

func Logger() *zap.Logger {
	return logger
}
