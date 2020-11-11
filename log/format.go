package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	isShowFuncName bool
	isShowFile     bool
)

func ShowFuncName(ok bool) {
	defaultLogger.SetFunc(ok)
}

func ShowLine(ok bool) {
	defaultLogger.SetLine(ok)
}

func GetCallerFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		fmt.Println("failed to get caller func name, stack:", skip)
		return ""
	}
	f := runtime.FuncForPC(pc)
	name := f.Name()
	return name[strings.LastIndex(name, ".")+1:]
}

func GetCallerLine(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		fmt.Println("failed to get caller func name, stack:", skip)
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func SetFormatter(f logrus.Formatter) {
	defaultLogger.SetFormatter(f)
}
