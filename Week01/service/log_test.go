package service

import "testing"

func TestInitLog(t *testing.T) {
	InitLog()
	logrus.Error()
	LogPhone.Info()
	LogWarn.Warn()
}
