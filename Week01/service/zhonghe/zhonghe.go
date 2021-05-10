package zhonghe

import (
	zhonghe2 "github.com/jin-Register/api/zhonghe"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/defu"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
)

type Zhonghe struct {
	ProjectId string
	UserName  string
	Code      string
	Mut       sync.Mutex
}

func NewZhonghe(projectId, userName, verifyCode string, mut sync.Mutex) *Zhonghe {
	return &Zhonghe{
		ProjectId: projectId,
		UserName:  userName,
		Code:      verifyCode,
		Mut:       mut,
	}
}

func (z *Zhonghe) Register() (err error) {
	z.UserName, err = defu.GetMobile(z.ProjectId)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(z.UserName) == 0 {
		logrus.Error("no mobile")
		return
	}

	service.LogPhone.Info("开始注册众和账号:" + z.UserName)

	err = z.GetCodeAndRegister()
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

func (z *Zhonghe) GetCodeAndRegister() (err error) {

	err = zhonghe2.GenerateCode(z.UserName, z.Mut)
	if err != nil {
		return errors.Wrapf(err, "GenerateCode error mobile:%s", z.UserName)
	}

	code, err := defu.GetCode(z.UserName, z.ProjectId)
	if err != nil {
		return errors.Wrapf(err, " getCode error mobile:%s", z.UserName)
	}

	err = zhonghe2.RegisterWithMobile(z.UserName, code)
	if err != nil {
		return errors.Wrapf(err, "RegisterWithMobile error mobile:%s,code:%s", z.UserName, code)
	}

	return nil
}
