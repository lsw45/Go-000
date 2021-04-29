package service

import (
	"github.com/sirupsen/logrus"
	"os"
)

var LogPhone *logrus.Logger
var LogWarn *logrus.Logger

func InitLog() {
	log, _ := os.OpenFile("./log/error.txt", os.O_APPEND, 0666)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2021-04-29 00:31:43",
	})
	logrus.SetOutput(log)

	LogPhone = logrus.New()
	phone, _ := os.OpenFile("./log/phone.txt", os.O_APPEND, 0666)
	LogPhone.SetOutput(phone)

	LogWarn = logrus.New()
	warn, _ := os.OpenFile("./log/warn.txt", os.O_APPEND, 0666)
	LogWarn.SetOutput(warn)
}
