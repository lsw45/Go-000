package jukun

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var GenerateCodeUrl = "https://h5api.jukunwang.com/api/user/Verification_Code/send"

var GenerateReq = map[string]string{"username": ""}

var GenerateSuccess = 1

type GenerateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

/*
{
    "code": 1,
    "msg": "验证码已经发送成功!",
    "data": ""
}
*/

func GenerateCode(mobile string) error {
	if len(mobile) == 0 {
		return errors.New("mobile is empty")
	}

	username := `{"username":` + mobile + `}`
	resp, err := http.Post(GenerateCodeUrl, "application/x-www-form-urlencoded", bytes.NewReader([]byte(username)))
	if err != nil {
		return errors.Wrap(err, "send code fail：获取验证码链接失效")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "send code fail：应答解析失败")
	}

	g := &GenerateResp{}
	err = json.Unmarshal(body, g)
	if err != nil {
		return errors.Wrap(err, "send code fail：反序列化失败")
	}

	if g.Code != 1 {
		return errors.Wrapf(err, "send code fail：%s", g.Msg)
	}
	return nil
}
