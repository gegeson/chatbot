package log

import (
	"fmt"
	glog "log"
	"os"
)

type logger struct {
	logger *glog.Logger
}

type Logger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

func New() Logger {
	return &logger{
		logger: glog.New(os.Stdout, "", glog.Ldate|glog.Ltime),
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	_format := fmt.Sprintf("[INFO] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	_format := fmt.Sprintf("[WARN] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	_format := fmt.Sprintf("[ERROR] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	_format := fmt.Sprintf("[DEBUG] %s", format)
	l.logger.Printf(_format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	_format := fmt.Sprintf("[Fatal] %s", format)
	l.logger.Fatalf(_format, args...)
}
