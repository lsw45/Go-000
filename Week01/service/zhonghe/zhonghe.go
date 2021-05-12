package zhonghe

import (
	"context"
	"github.com/jin-Register/sdk/defu"
	zhonghe2 "github.com/jin-Register/sdk/zhonghe"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
	"time"
)

type Zhonghe struct {
	ProjectId string
	UserName  string
	Code      string
	Mut       sync.Mutex
}

func NewZhonghe(projectId string, mut sync.Mutex) *Zhonghe {
	return &Zhonghe{
		ProjectId: projectId,
		Mut:       mut,
	}
}

func (z *Zhonghe) Register1() (err error) {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(context context.Context, netw, addr string) (net.Conn, error) {
				//本地地址  ipaddr是本地外网IP
				lAddr, err := net.ResolveTCPAddr("tcp", "")
				if err != nil {
					return nil, err
				}
				//被请求的地址
				rAddr, err := net.ResolveTCPAddr(netw, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(netw, lAddr, rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(35 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}

	return z.Register(client)
}

func (z *Zhonghe) Register(client http.Client) (err error) {
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

	err = z.GetCodeAndRegister(client)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}

func (z *Zhonghe) GetCodeAndRegister(client http.Client) (err error) {

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
