package service

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Platform interface {
	Register(client http.Client) (err error)
}

func Start(platform Platform, client http.Client) (err error) {
	err = platform.Register(client)
	if err != nil {
		logrus.Error(err)
		return err
	}
	AccountLog.Infof("新账号：%+v", platform)
	/*if kun, ok := platform.(*jukun.JuKun); ok {
		err = kun.Register()
		if err != nil {
			logrus.Error(err)
			return err
		}
		LogPhone.Infof("注册完成：%+v", kun)
	}

	if zhong, ok := platform.(*zhonghe.Zhonghe); ok {
		err = zhong.Register()
		if err != nil {
			logrus.Error(err)
			return err
		}
		LogPhone.Infof("注册完成：%+v", zhong)
	}*/

	return nil
}
