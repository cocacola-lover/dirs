package weblogger

import "log"

type WebLogger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewWebLogger() WebLogger {
	weblogger := newWebsocketMessenger()

	go weblogger.broadcastLogs()

	return WebLogger{
		Info:    log.New(weblogger, "[INFO]:", log.Ldate|log.Ltime),
		Warning: log.New(weblogger, "[WARNING]:", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(weblogger, "[ERROR]:", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
