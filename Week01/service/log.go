package service

import (
	"github.com/sirupsen/logrus"
	"os"
)

var LogPhone = logrus.New()
var AccountLog = logrus.New()

func InitLog(dir string) {
	log, _ := os.OpenFile(dir+"/error.txt", os.O_APPEND, 0666)
	logrus.SetOutput(log)

	phone, _ := os.OpenFile(dir+"/phone.txt", os.O_APPEND, 0666)
	LogPhone.SetOutput(phone)

	account, _ := os.OpenFile(dir+"/account.txt", os.O_APPEND, 0666)
	AccountLog.SetOutput(account)
}
