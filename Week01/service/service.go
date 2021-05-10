package service

import (
	"github.com/sirupsen/logrus"
)

type Platform interface {
	Register() (err error)
}

func Start(platform Platform) (err error) {
	err = platform.Register()
	if err != nil {
		logrus.Error(err)
		return err
	}
	LogPhone.Infof("注册完成：%+v", platform)

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
