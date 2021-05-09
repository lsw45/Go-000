package main

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/defu"
	"github.com/jin-Register/service/zhonghe"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var debug = false
var count = 5
var sid = "65429"

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/zhonghe")

	var mut sync.Mutex

	register(mut)
	/*for i := 0; i < count; i++ {
	}*/

	fmt.Println("本次批量注册任务完成")
}

func register(mut sync.Mutex) {
	mobile, err := defu.GetMobile(sid)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(mobile) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = GetCodeAndRegister(mobile, mut)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

func GetCodeAndRegister(mobile string, mut sync.Mutex) (err error) {

	err = zhonghe.GenerateCode(mobile, mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", mobile)
	}

	service.LogPhone.Info("开始注册众和账号:" + mobile)

	code, err := defu.GetCode(mobile, sid)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", mobile)
	}

	err = zhonghe.RegisterWithMobile(mobile, code)
	if err != nil {
		return errors.Wrapf(err, "RegisterWithMobile error mobile:%s,code:%s", mobile, code)
	}

	return nil
}
