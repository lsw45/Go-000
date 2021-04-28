package model

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	log, _ := os.OpenFile("./log.txt", os.O_APPEND, 0666)
	logrus.SetOutput(log)
	err := GenerateCode("16725568365")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
