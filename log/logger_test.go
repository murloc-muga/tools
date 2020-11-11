package log_test

import (
	"errors"
	"testing"

	"github.com/murloc-muga/tools/log"
	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	log.ShowFuncName(true)
	log.ShowLine(true)
	log.Debug("this is debug log")
	log.Infof("this is a info %d %d", 123, 456)
	log.WarnWF(log.Fields{
		"1": 1,
		"2": 2,
	}, "this is a warn with fields log")
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "20060102 15:04:05",
	})

	log.Error("this is err info:", errors.New("format error"))
	log.WithField("1", "1").Print("dd ", 33)
	log.WithError(errors.New("this is demo error")).Error()
}
