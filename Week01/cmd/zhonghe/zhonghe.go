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
	"time"
)

var debug = false
var count = 5
var sid = "65429"

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir)

	var lock sync.Mutex

	register(lock)
	/*for i := 0; i < count; i++ {
	}*/

	fmt.Println("本次批量注册任务完成")
}

func register(lock sync.Mutex) {
	mobile, err := defu.GetMobile(sid)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(mobile) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = zhonghe.GenerateCode(mobile, lock)
	if err != nil {
		logrus.Error(err)
		return
	}
	service.LogPhone.Info("开始注册众和账号:" + mobile)

	GetCodeAndRegister(mobile)

	return
}

var TimeOutErr = errors.New("get code timeout")

func GetCodeAndRegister(mobile string) (err error) {
	time.Sleep(3 * time.Second)

	var code = ""
	var retry = 1
	var timeout = time.After(20 * time.Second)

	for err != nil || len(code) == 0 {
		time.Sleep(time.Second * 1) // 每1s获取一次
		select {
		case <-timeout:
			service.LogPhone.Errorf("德芙验证码获取失败,mobile:%s,retry:%d", mobile, retry)
			return TimeOutErr
		default:
			retry++
			code, err = defu.GetCode(mobile, sid)
			if err == nil && len(code) > 0 {
				err = zhonghe.RegisterWithMobile(mobile, code)
				if err != nil {
					logrus.Error(err)
					return
				}
				return
			}
		}
	}
	return nil
}
