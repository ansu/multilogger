# Multi Logger
Multi logger is a Go package that provides a wrapper for both logrus and zap loggers, allowing you to easily switch between the two logging frameworks in your application.

## Installation
To use Dual Logger in your Go project, you can install it using go get:
```
go get github.com/ansu/multilogger
```

## Usage
To use Dual Logger in your code, you first need to create an instance of the logger wrapper:
```go
import (
 
  "github.com/ansu/multilogger"
)

// Create a new logrus logger
loggerWrapper := logger.NewLoggerWrapper("logrus", ctx)
//or 
loggerWrapper = logger.NewLoggerWrapper("zap", ctx)

```
Once you have created the logger wrapper, you can use it to log messages. Here's an example of how to log a message with fields using the logger wrapper:

```go
//With Common fields similar to logrus.WithFields
commonFields := map[string]interface{}{
		"userId":    "12345",
		"ipAddress": "192.168.0.1",
	}
	ctx := context.WithValue(context.Background(), "commonFields", commonFields)
  loggerWrapper.Info("This is an info log message.", commonFields)
//In the above example, the ctx variable is a context.Context instance that contains any additional contextual information you want to include in the log //message.

```

# License
Dual Logger is released under the MIT License. See LICENSE for details.
