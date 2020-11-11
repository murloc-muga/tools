package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var (
	defaultLogger = NewLogger(logrus.New())
)

func GetOutPut() io.Writer {
	return defaultLogger.GetOutput()
}

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func SetLevel(level string) {
	defaultLogger.SetLevel(level)
}

func Debug(args ...interface{}) {
	defaultLogger.wrap().Debug(args...)
}

func Print(args ...interface{}) {
	defaultLogger.wrap().Print(args...)
}

func Info(args ...interface{}) {
	defaultLogger.wrap().Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.wrap().Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.wrap().Error(args...)
}

// will exit process
func Fatal(args ...interface{}) {
	defaultLogger.wrap().Fatal(args...)
}

// will panic
func Panic(args ...interface{}) {
	defaultLogger.wrap().Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.wrap().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.wrap().Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	defaultLogger.wrap().Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.wrap().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.wrap().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.wrap().Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.wrap().Panicf(format, args...)
}

func Debugln(args ...interface{}) {
	defaultLogger.wrap().Debugln(args...)
}

func Println(args ...interface{}) {
	defaultLogger.wrap().Println(args...)

}

func Infoln(args ...interface{}) {
	defaultLogger.wrap().Infoln(args...)
}

func Warnln(args ...interface{}) {
	defaultLogger.wrap().Warnln(args...)
}

func Errorln(args ...interface{}) {
	defaultLogger.wrap().Errorln(args...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.wrap().Fatalln(args...)
}

func Panicln(args ...interface{}) {
	defaultLogger.wrap().Panicln(args...)
}

func DebugWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Debug(msg...)
}

func InfoWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Info(msg...)
}

func WarnWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Warn(msg...)
}

func ErrorWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Error(msg...)
}

func FatalWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Fatal(msg...)
}

func PanicWF(value Fields, msg ...interface{}) {
	defaultLogger.wrap().WithFields(value).Panic(msg...)
}

func WithFields(value Fields) *entry {
	return defaultLogger.wrap().WithFields(value)
}

func WithField(key string, value interface{}) *entry {
	return defaultLogger.wrap().WithField(key, value)
}

func WithError(err error) *entry {
	return defaultLogger.wrap().WithError(err)
}
