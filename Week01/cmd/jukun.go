package main

import (
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/jukun"
	"github.com/jin-Register/service/xiaobai"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

/*
1、获取手机号
2、金巨鲲发送验证码
3、获取验证码
4、提交注册信息
*/
func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir)

	mobile, err := xiaobai.GetMobile()
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(mobile) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = jukun.GenerateCode(mobile)
	if err != nil {
		logrus.Error(err)
		return
	}

	retry := 1
	code, err := xiaobai.GetCode(mobile)
	if err != nil || len(code) == 0 {
		// 获取验证码失败，重试5次
		for retry < 6 {
			time.Sleep(time.Second * time.Duration(retry))
			code1, err1 := xiaobai.GetCode(mobile)
			if err1 == nil {
				err = err1
				code = code1
				break
			}
			err = err1
			code = code1
			retry++
		}
	}
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(code) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = jukun.RegisterWithMobile(mobile, code)
	if err != nil {
		logrus.Error(err)
		return
	}
}
