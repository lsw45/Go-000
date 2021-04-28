package main

import (
	"github.com/jin-Register/model"
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
	log, _ := os.OpenFile("./log.txt", os.O_APPEND, 0666)
	logrus.SetOutput(log)

	mobile, err := model.GetMobile()
	if err != nil {
		logrus.Error(err)
		return
	}

	err = model.GenerateCode(mobile)
	if err != nil {
		logrus.Error(err)
		return
	}

	retry := 1
	code, err := model.GetCode(mobile)
	if err != nil {
		// 获取验证码失败，重试5次
		for retry < 6 {
			time.Sleep(time.Second * time.Duration(retry))
			code1, err1 := model.GetCode(mobile)
			if err1 == nil {
				err = err1
				code = code1
				break
			}
			err = err1
			retry++
		}
	}
	if err != nil {
		logrus.Error(err)
		return
	}

	err = model.RegisterWithMobile(mobile, code)
	if err != nil {
		logrus.Error(err)
		return
	}
}
