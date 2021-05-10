package service

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog("D:\\workspace\\go_workspace\\Go-000\\Week01\\log\\zhonghe\\")
	logrus.Error("error")
	LogPhone.Info("info")
	LogPhone.Warn("warn")
}
