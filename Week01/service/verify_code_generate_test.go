package service

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	log, _ := os.OpenFile("./log.txt", os.O_APPEND, 0666)
	phone, _ := os.OpenFile("./phone.txt", os.O_APPEND, 0666)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2021-04-29 00:31:43",
	})

	logrus.SetOutput(log)

	logPhone := logrus.New()
	logPhone.SetOutput(phone)

	err := GenerateCode("16725568365")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
