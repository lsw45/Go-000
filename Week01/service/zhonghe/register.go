package zhonghe

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var RegisterUrl = "http://5g.hhml.cn/ajax/operate"

var RegisterSuccess = true

var registerReq = url.Values{
	"command": {"register"},
}

type RegisterResp struct {
	Ret bool   `json:"ret"`
	Msg string `json:"msg"`
}

/*
{
    "ret":false,
    "msg":"验证码错误"
}
*/
var field = `{"mobile":"%s","vcode":"%s","password":"zhonghe123456","repassword":"zhonghe123456","up_code":"361f428132"}`

func RegisterWithMobile(mobile string, code string) error {
	registerReq.Add("field", fmt.Sprintf(field, mobile, code))

	resp, err := http.Post(RegisterUrl, "application/x-www-form-urlencoded", strings.NewReader(registerReq.Encode()))
	if err != nil {
		return errors.Wrapf(err, "众和注册手机号失败:%s", mobile)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "众和注册手机号失败:%s", mobile)
	}

	r := &RegisterResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.Wrapf(err, "众和注册手机号失败:%s", mobile)
	}

	if r.Ret != RegisterSuccess {
		service.LogPhone.Errorf("账号:%s,注册异常:%s", mobile, r.Msg)
	} else {
		service.LogPhone.Infof("众和新账号:%s,注册成功", mobile)
	}

	return nil
}
