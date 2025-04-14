package amebo

import "fmt"

type LoggerOptions struct {
	// Log level: debug, info, warn, error, fatal
	LogLevel string
	// Where to output logs; console, file
	LogOutput string
	// Log output format: text, json
	LogFormat string
	// Log file prefix (only used when LogOutput is file)
	LogFilePrefix string
	// Log directory (only used when LogOutput is file)
	LogDirectory string
	// Max file size in MB (only used when LogOutput is file)
	MaxFileSize int
}

func NewLoggerOptions() *LoggerOptions {
	return &LoggerOptions{}
}

func (o *LoggerOptions) SetLogLevel(level string) *LoggerOptions {
	o.LogLevel = level
	return o
}

func (o *LoggerOptions) SetLogOutput(output string) *LoggerOptions {
	o.LogOutput = output
	return o
}

func (o *LoggerOptions) SetLogFormat(format string) *LoggerOptions {
	o.LogFormat = format
	return o
}

func (o *LoggerOptions) SetLogFilePrefix(prefix string) *LoggerOptions {
	o.LogFilePrefix = prefix
	return o
}

func (o *LoggerOptions) SetLogDirectory(directory string) *LoggerOptions {
	o.LogDirectory = directory
	return o
}

func (o *LoggerOptions) SetMaxFileSize(size int) *LoggerOptions {
	o.MaxFileSize = size
	return o
}

func (o *LoggerOptions) Validate() error {
	if o.LogLevel == "" {
		return fmt.Errorf("log level is required, ensure LogLevel is set")
	}

	if o.LogOutput == "" {
		o.LogOutput = "console"
	}

	if o.LogFormat == "" {
		o.LogFormat = "text"
	}

	if o.LogDirectory == "" {
		o.LogDirectory = "logs"
	}

	if o.LogFilePrefix == "" {
		o.LogFilePrefix = "amebo"
	}

	if o.MaxFileSize <= 0 {
		o.MaxFileSize = 50
	}

	return nil
}
