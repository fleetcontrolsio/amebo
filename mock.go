package amebo

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"
)

// MockLogger implements the Logger interface for testing
type MockLogger struct {
	mock.Mock
}

// NewMockLogger creates a new mock logger
func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

// Debug logs a debug level message
func (m *MockLogger) Debug(msg string, fields ...zapcore.Field) {
	m.Called(msg, fields)
}

// Info logs an info level message
func (m *MockLogger) Info(msg string, fields ...zapcore.Field) {
	m.Called(msg, fields)
}

// Warn logs a warning level message
func (m *MockLogger) Warn(msg string, fields ...zapcore.Field) {
	m.Called(msg, fields)
}

// Error logs an error level message
func (m *MockLogger) Error(msg string, fields ...zapcore.Field) {
	m.Called(msg, fields)
}

// Fatal logs a fatal level message
func (m *MockLogger) Fatal(msg string, fields ...zapcore.Field) {
	m.Called(msg, fields)
}
