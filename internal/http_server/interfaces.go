package httpserver

// Logger ...
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
}
