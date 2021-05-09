package main

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/defu"
	"github.com/jin-Register/service/jukun"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

/*
1、获取手机号
2、金巨鲲发送验证码
3、获取验证码
4、提交注册信息
*/
var debug = false
var count = 5
var sid = "163260"

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/jukun")

	var mut sync.Mutex

	for i := 0; i < count; i++ {
		register(mut)
	}

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

	err = GetCodeAndBindSecret(mobile, jukun.Passwd, mut)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

func GetCodeAndRegister(mobile string, mut sync.Mutex) (err error) {

	err = jukun.GenerateCode(mobile, mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", mobile)
	}
	service.LogPhone.Info("开始注册巨鲲账号:" + mobile)

	code, err := defu.GetCode(mobile, sid)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", mobile)
	}

	err = jukun.RegisterWithMobile(mobile, code)
	if err != nil {
		return errors.Wrapf(err, "RegisterWithMobile error mobile:%s,code:%s", mobile, code)
	}

	return nil
}

func GetCodeAndBindSecret(mobile, passwd string, mut sync.Mutex) (err error) {

	token, err := jukun.Login(mobile, passwd)
	if err != nil {
		return errors.Wrapf(err, "Login error mobile:%s", mobile)
	}

	err = jukun.GenerateCode(mobile, mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", mobile)
	}

	code, err := defu.GetCode(mobile, sid)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", mobile)
	}

	_, err = jukun.BindSecret(token, code, mobile)
	if err != nil {
		return errors.Wrapf(err, "BindSecret error mobile:%s,code:%s", mobile, code)
	}

	return nil
}
