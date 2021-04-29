package service

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog()
	logrus.Error("error")
	LogPhone.Info("info")
	LogWarn.Warn("warn")
}
