package main

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/jukun"
	"github.com/jin-Register/service/xiaobai"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

/*
1、获取手机号
2、金巨鲲发送验证码
3、获取验证码
4、提交注册信息
*/
var debug = false

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir)

	var lock sync.Mutex
	var count = 5

	for i := 0; i < count; i++ {
		register(lock)
	}

	fmt.Println("本次批量注册任务完成")
}

func register(lock sync.Mutex) {
	mobile, err := xiaobai.GetMobile()
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(mobile) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = jukun.GenerateCode(mobile, lock)
	if err != nil {
		logrus.Error(err)
		return
	}
	service.LogPhone.Info("开始注册账号:" + mobile)

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
			service.LogPhone.Errorf("小白验证码获取失败,mobile:%s,retry:%d", mobile, retry)
			return TimeOutErr
		default:
			retry++
			code, err = xiaobai.GetCode(mobile)
			if err == nil && len(code) > 0 {
				err = jukun.RegisterWithMobile(mobile, code)
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
