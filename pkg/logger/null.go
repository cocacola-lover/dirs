package logger

type nullLogger struct{}

func (t nullLogger) Println(v ...any)               {}
func (t nullLogger) Printf(format string, v ...any) {}

func NullLogger() (Logger, Logger, Logger) {
	return nullLogger{}, nullLogger{}, nullLogger{}
}
