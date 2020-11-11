package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields
type Logger struct {
	logger   *logrus.Logger
	withFunc bool
	withLine bool
}

type entry struct {
	entry *logrus.Entry
}

func NewLogger(logger *logrus.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) SetFunc(ok bool) {
	l.withFunc = true
}

func (l *Logger) SetLine(ok bool) {
	l.withLine = true
}

// func (l *Logger) WithField(key string, value interface{}) *entry {
// 	var entry entry
// 	entry.entry = l.logger.WithField(key, value)
// 	return &entry
// }

// func (l *Logger) WithError(err error) *entry {
// 	var entry entry
// 	entry.entry = l.logger.WithError(err)
// 	return &entry
// }

func (l *Logger) wrap() *entry {
	entry := newEntry(l)
	if l.withFunc {
		entry.WithFunc(2)
	}
	if l.withLine {
		entry.WithLine(2)
	}
	return entry
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "debug":
		l.logger.SetLevel(logrus.DebugLevel)
	case "info":
		l.logger.SetLevel(logrus.InfoLevel)
	case "warn":
		l.logger.SetLevel(logrus.WarnLevel)
	case "error":
		l.logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		// Calls os.Exit(1) after logging
		l.logger.SetLevel(logrus.FatalLevel)
	case "panic":
		// Calls panic() after logging
		l.logger.SetLevel(logrus.PanicLevel)
	default:
		l.logger.SetLevel(logrus.InfoLevel)
	}
}

func (l *Logger) GetOutput() io.Writer {
	return l.logger.Out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.logger.SetFormatter(formatter)
}

func newEntry(l *Logger) *entry {
	var entry entry
	entry.entry = l.logger.WithFields(nil)
	return &entry
}

func (entry *entry) WithFields(fields Fields) *entry {
	entry.entry = entry.entry.WithFields(logrus.Fields(fields))
	return entry
}

func (entry *entry) WithField(key string, value interface{}) *entry {
	return entry.WithFields(Fields{key: value})
}

func (entry *entry) WithError(err error) *entry {
	entry.entry = entry.entry.WithError(err)
	return entry
}

func (entry *entry) WithFunc(skip int) *entry {
	fn := GetCallerFuncName(skip + 1)
	if fn == "" {
		return entry
	}
	return entry.WithField("fn", fn)
}

func (entry *entry) WithLine(skip int) *entry {
	line := GetCallerLine(skip + 1)
	if line == "" {
		return entry
	}
	return entry.WithField("ln", line)
}

func (entry *entry) Debug(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Debug(args...)
	}
}

func (entry *entry) Print(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Print(args...)
	}
}

func (entry *entry) Info(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Info(args...)
	}
}

func (entry *entry) Warn(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warn(args...)
	}
}

func (entry *entry) Warning(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warning(args...)
	}
}

func (entry *entry) Error(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Error(args...)
	}
}

func (entry *entry) Fatal(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Fatal(args...)
	}
}

func (entry *entry) Panic(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Panic(args...)
	}
}

func (entry *entry) Debugf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Debugf(format, args...)
	}
}

func (entry *entry) Infof(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Infof(format, args...)
	}
}

func (entry *entry) Printf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Printf(format, args...)
	}
}

func (entry *entry) Warnf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warnf(format, args...)
	}
}

func (entry *entry) Warningf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warningf(format, args...)
	}
}

func (entry *entry) Errorf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Errorf(format, args...)
	}
}

func (entry *entry) Fatalf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Fatalf(format, args...)
	}
}

func (entry *entry) Panicf(format string, args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Panicf(format, args...)
	}
}

func (entry *entry) Debugln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Debugln(args...)
	}
}

func (entry *entry) Infoln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Infoln(args...)
	}
}

func (entry *entry) Println(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Println(args...)
	}
}

func (entry *entry) Warnln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warnln(args...)
	}
}

func (entry *entry) Warningln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Warningln(args...)
	}
}

func (entry *entry) Errorln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Errorln(args...)
	}
}

func (entry *entry) Fatalln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Fatalln(args...)
	}
}

func (entry *entry) Panicln(args ...interface{}) {
	if entry.entry != nil {
		entry.entry.Panicln(args...)
	}
}
