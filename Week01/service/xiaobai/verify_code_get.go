package xiaobai

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var GetCodeUrl = "http://api.xiaobai188.com:188/api/getMsg?token=" + token + "&sid=%s&phone=%s"

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
        "content": "【金巨鲲】验证码811631,切勿告知他人！",
        "code": "811631"
    }
}
*/

func GetCode(mobile string, sid string) (code string, err error) {
	resp, err := http.Get(fmt.Sprintf(GetCodeUrl, sid, mobile))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "小白获取验证码失败:应答解析失败")
	}

	a := &CodeResp{}

	err = json.Unmarshal(body, a)
	if err != nil {
		return "", errors.Wrap(err, "小白获取验证码失败:反序列化失败")
	}
	if a.Code != Success && a.Msg != "获取验证码成功~" {
		return "", errors.New("小白获取验证码失败:" + a.Msg)
	}
	return a.Data.Code, nil
}
