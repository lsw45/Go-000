package jukun

import (
	"encoding/json"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var RegisterUrl = "https://h5api.jukunwang.com/api/user/public/register"

var Success = 1
var Passwd = "jukun123456"
var registerReq = url.Values{
	"password": {Passwd},
	"pay_pass": {Passwd},
	"invcode":  {"54NNM"},
	"que_id":   {"3"},
	"answer":   {"狗"},
}

type RegisterRequest struct {
	Username         int64  `json:"username"`
	Password         int    `json:"password"`
	PayPass          int    `json:"pay_pass"`
	VerificationCode string `json:"verification_code"`
	Invcode          string `json:"invcode"`
	QueID            int    `json:"que_id"`
	Answer           string `json:"answer"`
}

type RegisterResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

/*
{
    "code": 0,
    "msg": "此账号已存在!",
    "data": ""
}
{
    "code": 1,
    "msg": "验证码已经发送成功!",
    "data": ""
}
*/

func RegisterWithMobile(mobile string, code string) error {
	registerReq.Add("username", mobile)
	registerReq.Add("verification_code", code)

	resp, err := http.Post(RegisterUrl, "application/x-www-form-urlencoded", strings.NewReader(registerReq.Encode()))
	if err != nil {
		return errors.Wrapf(err, "注册手机号失败:%s", mobile)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "注册手机号失败:%s", mobile)
	}

	r := &RegisterResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.Wrapf(err, "注册手机号失败:%s", mobile)
	}

	if r.Code != Success {
		service.LogPhone.Errorf("账号:%s,注册异常:%s", mobile, r.Msg)
	} else {
		service.LogPhone.Infof("新账号:%s,注册成功", mobile)
	}

	return nil
}
