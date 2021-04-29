package service

import (
	"github.com/sirupsen/logrus"
	"os"
)

var LogPhone = logrus.New()
var LogWarn = logrus.New()

func InitLog(dir string) {
	log, _ := os.OpenFile(dir+"/log/error.txt", os.O_APPEND, 0666)
	logrus.SetOutput(log)

	phone, _ := os.OpenFile(dir+"/log/phone.txt", os.O_APPEND, 0666)
	LogPhone.SetOutput(phone)

	warn, _ := os.OpenFile(dir+"/log/warn.txt", os.O_APPEND, 0666)
	LogWarn.SetOutput(warn)
}
