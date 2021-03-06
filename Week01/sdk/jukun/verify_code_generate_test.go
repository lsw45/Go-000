package jukun

import (
	"github.com/jin-Register/service"
	"github.com/sirupsen/logrus"
	"sync"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	var lock sync.Mutex
	err := GenerateCode("1672556836", lock)
	service.InitLog("D:\\workspace\\Go-000\\Week01")
	if err != nil {
		logrus.Error(err)
		t.Fatal(err)
	}
	t.Log("success")
}
