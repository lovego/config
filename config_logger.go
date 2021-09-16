package config

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/lovego/alarm"
	"github.com/lovego/fs"
	loggerPkg "github.com/lovego/logger"
	"github.com/lovego/mailer"
)

const isRun = "isRun"

var theMailer = getMailer()
var theAlarm = getAlarm()
var theLogger, theHttpLogger *loggerPkg.Logger

func Mailer() *mailer.Mailer {
	return theMailer
}
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
		if os.Getenv(isRun) != `` {
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
	if os.Getenv(isRun) != `` {
		logger.SetLevel(loggerPkg.Debug)
	} else {
		logger.Set("project", Name())
		logger.Set("env", Env().String())
	}
	return logger
}

func Protect(fn func()) {
	defer Logger().Recover()
	fn()
}

func getMailer() *mailer.Mailer {
	m, err := mailer.New(theConfig.Mailer)
	if err != nil {
		log.Panic(err)
	}
	return m
}
func getAlarm() *alarm.Alarm {
	return alarm.New(alarm.MailSender{
		Receivers: Keepers(),
		Mailer:    Mailer(),
	}, nil, alarm.SetPrefix(DeployName()))
}
