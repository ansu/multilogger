package main

import (
	"context"
	"fmt"

	"github.com/ansu/multilogger/logger"
)

func main() {
	fmt.Println("Dummy")
	commonFields := map[string]interface{}{
		"userId":    "12345",
		"ipAddress": "192.168.0.1",
	}
	ctx := context.WithValue(context.Background(), "commonFields", commonFields)

	loggerWrapper := logger.NewLoggerWrapper("logrus", ctx)
	loggerWrapper.Info("This is an info log message.", commonFields)

	loggerWrapper = logger.NewLoggerWrapper("zap", ctx)
	loggerWrapper.Info("This is an info log message.", commonFields)

}
