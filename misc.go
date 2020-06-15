package config

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/lovego/alarm"
	"github.com/lovego/fs"
	loggerPkg "github.com/lovego/logger"
)

var theLogger, theHttpLogger *loggerPkg.Logger

var theAlarm = alarm.New(alarm.MailSender{
	Receivers: Keepers(),
	Mailer:    Mailer(),
}, 0, 10*time.Second, time.Minute, alarm.SetPrefix(DeployName()))

func Alarm() *alarm.Alarm {
	return theAlarm
}

func Logger() *loggerPkg.Logger {
	if theLogger == nil {
		theLogger = NewLoggerFromWriter(os.Stderr)
	}
	return theLogger
}

func HttpLogger() *loggerPkg.Logger {
	if theHttpLogger == nil {
		if DevMode() {
			theHttpLogger = NewLoggerFromWriter(os.Stdout)
		} else {
			theHttpLogger = NewLogger("http.log")
		}
	}
	return theHttpLogger
}

func NewLogger(paths ...string) *loggerPkg.Logger {
	file, err := fs.NewLogFile(filepath.Join(
		append([]string{Root(), `log`}, paths...)...,
	))
	if err != nil {
		Logger().Fatal(err)
	}
	return NewLoggerFromWriter(file)
}

func NewLoggerFromWriter(writer io.Writer) *loggerPkg.Logger {
	logger := loggerPkg.New(writer)
	logger.SetAlarm(Alarm())
	if DevMode() {
		logger.SetLevel(loggerPkg.Debug)
	} else {
		logger.Set("project", Name())
		logger.Set("env", Env())
	}
	return logger
}

func Protect(fn func()) {
	defer Logger().Recover()
	fn()
}
