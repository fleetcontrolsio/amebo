package amebo

// NewLogger creates a new logger based on the provided options
func NewLogger(options *LoggerOptions) (Logger, error) {
	err := options.Validate()
	if err != nil {
		return nil, err
	}

	return NewZapLogger(options)
}
