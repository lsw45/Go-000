package jukun

import (
	"github.com/jin-Register/api/jukun"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/defu"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type JuKun struct {
	ProjectId string
	UserName  string
	Secret    string
	Code      string
	Mut       sync.Mutex
}

func NewJukun(projectId, userName, verifyCode string, mut sync.Mutex) *JuKun {
	return &JuKun{
		ProjectId: projectId,
		UserName:  userName,
		Code:      verifyCode,
		Mut:       mut,
	}
}

func (j *JuKun) Register() (err error) {
	j.UserName, err = defu.GetMobile(j.ProjectId)
	if err != nil {
		return
	}

	if len(j.UserName) == 0 {
		logrus.Error("no mobile")
		return
	}

	err = j.GetCodeAndRegister()
	if err != nil {
		return
	}

	err = j.GetCodeAndBindSecret()
	if err != nil {
		return
	}

	return
}

func (j *JuKun) GetCodeAndRegister() (err error) {

	err = jukun.GenerateCode(j.UserName, j.Mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", j.UserName)
	}
	service.LogPhone.Info("开始注册巨鲲账号:" + j.UserName)

	j.Code, err = defu.GetCode(j.UserName, j.ProjectId)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", j.UserName)
	}

	err = jukun.RegisterWithMobile(j.UserName, j.Code)
	if err != nil {
		return errors.Wrapf(err, "RegisterWithMobile error mobile:%s,code:%s", j.UserName, j.Code)
	}

	return nil
}

func (j *JuKun) GetCodeAndBindSecret() (err error) {

	token, err := jukun.Login(j.UserName)
	if err != nil {
		return errors.Wrapf(err, "Login error mobile:%s", j.UserName)
	}

	err = jukun.GenerateCode(j.UserName, j.Mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", j.UserName)
	}

	j.Code, err = defu.GetCode(j.UserName, j.ProjectId)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", j.UserName)
	}

	j.Secret, err = jukun.BindSecret(token, j.Code, j.UserName)
	if err != nil {
		return errors.Wrapf(err, "BindSecret error mobile:%s,code:%s", j.UserName, j.Code)
	}

	return nil
}
