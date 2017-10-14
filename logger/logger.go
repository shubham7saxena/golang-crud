package logger

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Debugrf(r *http.Request, format string, args ...interface{})
	Infof(format string, args ...interface{})
	Inforf(r *http.Request, format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Errorrf(r *http.Request, format string, args ...interface{})
}

type SurgeAppLogger struct {
	*logrus.Logger
}

var Log *SurgeAppLogger

func SetupLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}

	logrusVar := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log = &SurgeAppLogger{logrusVar}
	return nil
}

func (logger *SurgeAppLogger) Debugrf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Debugf(format, args...)
}

func (logger *SurgeAppLogger) Inforf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Infof(format, args...)
}

func (logger *SurgeAppLogger) Errorrf(r *http.Request, format string, args ...interface{}) {
	logger.httpRequestLogEntry(r).Errorf(format, args...)
}

func (logger *SurgeAppLogger) httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"RequestMethod": r.Method,
		"Host":          r.Host,
		"Path":          r.URL.Path,
		"Query":         r.URL.RawQuery,
	})
}
