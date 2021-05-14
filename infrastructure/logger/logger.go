package logger

// A Logger belong to the usecases layer.
type LoggerItf interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
