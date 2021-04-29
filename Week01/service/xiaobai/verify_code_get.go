package xiaobai

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var GetCodeUrl = "http://api.xiaobai188.com:188/api/getMsg?token=" + token + "&sid=163260&phone="

type CodeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Content string `json:"content"`
		Code    string `json:"code"`
	} `json:"data"`
}

/*
{
    "code": 20000,
    "msg": "获取验证码成功~",
    "data": {
        "content": "【金巨鲲】验证码811631，切勿告知他人！",
        "code": "811631"
    }
}
*/
func GetCode(mobile string) (code string, err error) {
	resp, err := http.Get(GetCodeUrl + mobile)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "receive code fail：应答解析失败")
	}

	a := &CodeResp{}

	err = json.Unmarshal(body, a)
	if err != nil {
		return "", errors.Wrap(err, "receive code fail：反序列化失败")
	}
	if a.Code == Success {
		return "", errors.Wrap(err, "receive code fail：验证码不正确")
	}
	return a.Data.Code, nil
}
