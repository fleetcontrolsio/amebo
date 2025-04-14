# Amebo

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A flexible, structured logging library for Go applications built on top of [zap](https://github.com/uber-go/zap) with simplified configuration and multiple output options.

## Features

- Simple yet powerful logging interface
- Support for both console and file-based logging
- Multiple output formats (JSON/text)
- Log rotation for file-based logging
- Configurable log levels
- Built-in mock for testing
- Fluent API for configuration

## Installation

```bash
go get github.com/fleetcontrolsio/amebo
```

## Quick Start

```go
package main

import (
	"github.com/fleetcontrolsio/amebo"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create logger with default console output
	options := amebo.NewLoggerOptions().
		SetLogLevel("INFO").
		SetLogFormat("json")

	logger, err := amebo.NewLogger(options)
	if err != nil {
		panic(err)
	}

	// Log messages
	logger.Info("Application started")
	logger.Debug("Debug message", zapcore.Field{Key: "value", Type: zapcore.StringType, String: "test"})
	logger.Error("Something went wrong", zapcore.Field{Key: "error_code", Type: zapcore.Int64Type, Integer: 500})
}
```

## Configuration Options

### Console Logging

```go
options := amebo.NewLoggerOptions().
	SetLogLevel("INFO").      // DEBUG, INFO, WARN, ERROR, FATAL
	SetLogOutput("console").  // Output to stdout
	SetLogFormat("json")      // json or text format
```

### File Logging

```go
options := amebo.NewLoggerOptions().
	SetLogLevel("INFO").
	SetLogOutput("file").
	SetLogFormat("json").
	SetLogDirectory("logs").
	SetLogFilePrefix("myapp").
	SetMaxFileSize(100)  // In MB
```

## Log Levels

- `DEBUG`: Detailed information, typically useful for debugging
- `INFO`: General operational information
- `WARN`: Warning events that might cause issues
- `ERROR`: Error events that might still allow the application to continue running
- `FATAL`: Severe error events that cause the application to terminate

## Testing with Mock Logger

```go
import (
	"testing"
	
	"github.com/fleetcontrolsio/amebo"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"
)

func TestWithLogger(t *testing.T) {
	mockLogger := amebo.NewMockLogger()
	
	// Set expectations
	mockLogger.On("Info", "test message", mock.Anything).Return()
	
	// Use mock logger in your code
	mockLogger.Info("test message", zapcore.Field{Key: "test", Type: zapcore.StringType, String: "value"})
	
	// Verify expectations were met
	mockLogger.AssertExpectations(t)
}
```

## License

MIT

## Dependencies

- [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging
- [lumberjack](https://github.com/natefinch/lumberjack) - Log rotation
- [testify](https://github.com/stretchr/testify) - Testing toolkit