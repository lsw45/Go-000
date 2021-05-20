package fil

import (
	"github.com/jin-Register/sdk/defu"
	"github.com/jin-Register/sdk/fil"
	"github.com/jin-Register/sdk/idcard"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type FilCoin struct {
	ProjectId string
	Name      string
	IdCard    string
	Mobile    string
	Code      string
	Mut       sync.Mutex
}

func NewFilCoin(projectId string, mut sync.Mutex) *FilCoin {
	return &FilCoin{
		ProjectId: projectId,
		Mut:       mut,
	}
}

func (f *FilCoin) Register(client http.Client) (err error) {
	f.Mobile, err = defu.GetMobile(f.ProjectId, "1")
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(f.Mobile) == 0 {
		logrus.Error("no mobile")
		return
	}

	service.LogPhone.Info("开始注册fil账号:" + f.Mobile)

	err = f.GetCodeAndRegister(client)
	if err != nil {
		return
	}

	return
}

func (f *FilCoin) GetCodeAndRegister(client http.Client) (err error) {

	err = fil.SendCode(f.Mobile, f.Mut)
	if err != nil {
		return errors.Wrapf(err, "SendCode error mobile:%s", f.Mobile)
	}

	f.Code, err = defu.GetCode(f.Mobile, f.ProjectId)
	if err != nil {
		return errors.Wrapf(err, "getCode error mobile:%s", f.Mobile)
	}

	userInfo, err := idcard.GetUserInfo()
	if err != nil {
		return err
	}

	f.Name = userInfo.Name
	f.IdCard = userInfo.ID

	err = fil.Register(f.Name, f.IdCard, f.Mobile, f.Code)
	if err != nil {
		return errors.Wrapf(err, "register error mobile:%s,code:%s", f.Mobile, f.Code)
	}

	return nil
}
