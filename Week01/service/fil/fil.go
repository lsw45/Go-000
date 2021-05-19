package fil

import (
	"bufio"
	"github.com/jin-Register/sdk/defu"
	"github.com/jin-Register/sdk/fil"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
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
	f.Mobile, err = defu.GetMobile(f.ProjectId)
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
		logrus.Error(err)
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

	f.Name, f.IdCard, err = getIdCard()
	if err != nil {
		return err
	}

	err = fil.Register(f.Name, f.IdCard, f.Mobile, f.Code)
	if err != nil {
		return errors.Wrapf(err, "register error mobile:%s,code:%s", f.Mobile, f.Code)
	}

	return nil
}

func getIdCard() (string, string, error) {
	file, err := os.Open("./idCard.txt")
	if err != nil {
		return "", "", errors.Wrap(err, "idCard文件打开出错")
	}
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(file)

	//遇到\n结束读取
	b, err := buf.ReadBytes('\n')
	if err != nil {
		if err == io.EOF {
			return "", "", errors.New("idCard.txt is empty")
		}
		return "", "", err
	}

	idCard := strings.Split(string(b), "-")
	if len(idCard) < 2 {
		return "", "", errors.Wrap(errors.New("idCard.txt is wrong"), string(b))
	}
	return idCard[0], idCard[1], nil
}
