package config

var (
	// The Logger is a pointer to Logger struct
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	// Initializa logger
	logger = NewLogger(prefix)
	return logger
}