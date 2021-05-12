package jukun

import (
	"github.com/jin-Register/sdk/haima"
	"github.com/jin-Register/sdk/jukun"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type JuKun struct {
	ProjectId  string
	UserName   string
	Secret     string
	Code       string
	Mut        sync.Mutex
	CodeRelate string
}

func NewJukun(projectId string, mut sync.Mutex) *JuKun {
	return &JuKun{
		ProjectId: projectId,
		Mut:       mut,
	}
}

func (j *JuKun) Register() (err error) {

	j.UserName, j.CodeRelate, err = haima.GetMobile(j.ProjectId)
	if err != nil {
		return
	}

	if len(j.UserName) == 0 {
		err = errors.New("no mobile")
		logrus.Error(err.Error())
		return
	}

	exit, err := jukun.UserExit(j.UserName)
	if err != nil {
		return
	}

	if exit {
		err = j.GetCodeAndChangePasswd()
	} else {
		err = j.GetCodeAndRegister()
	}

	return
}

// 新注册用户，设置云动码
func (j *JuKun) GetCodeAndRegister() (err error) {

	err = jukun.GenerateCode(j.UserName, j.Mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", j.UserName)
	}
	service.LogPhone.Info("开始注册巨鲲账号:" + j.UserName)

	j.Code, err = haima.GetCode(j.UserName, j.CodeRelate)
	if err != nil {
		return errors.Wrapf(err, "getCode error mobile:%s", j.UserName)
	}

	err = jukun.RegisterWithMobile(j.UserName, j.Code)
	if err != nil {
		return errors.Wrapf(err, "RegisterWithMobile error mobile:%s,code:%s", j.UserName, j.Code)
	}

	token, err := jukun.Login(j.UserName)
	if err != nil {
		return errors.Wrapf(err, "Login error mobile:%s", j.UserName)
	}

	j.Secret, err = jukun.BindSecret(j.UserName, j.Code, token)
	if err != nil {
		return errors.Wrapf(err, "BindSecret error mobile:%s,code:%s", j.UserName, j.Code)
	}

	return nil
}

// 已存在用户，改登陆密码，改交易密码，重新提交云动码
func (j *JuKun) GetCodeAndChangePasswd() (err error) {

	err = jukun.GenerateCode(j.UserName, j.Mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", j.UserName)
	}

	j.Code, err = haima.GetCode(j.UserName, j.CodeRelate)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", j.UserName)
	}

	err = jukun.ChangePasswd(j.UserName, j.Code)
	if err != nil {
		return errors.Wrapf(err, "ChangePasswd error mobile:%s,code:%s", j.UserName, j.Code)
	}

	token, err := jukun.Login(j.UserName)
	if err != nil {
		return errors.Wrapf(err, "Login error mobile:%s", j.UserName)
	}

	err = jukun.ForgetPass(j.UserName, j.Code, token)
	if err != nil {
		return errors.Wrapf(err, "ForgetPass error mobile:%s,code:%s", j.UserName, j.Code)
	}

	j.Secret, err = jukun.BindSecret(j.UserName, j.Code, token)
	if err != nil {
		return errors.Wrapf(err, "BindSecret error mobile:%s,code:%s", j.UserName, j.Code)
	}

	return nil
}
