package jukun

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
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
{
    "code": 0,
    "msg": "请稍后重试!",
    "data": ""
}
*/

func GenerateCode(mobile string, lock sync.Mutex) error {
	lock.Lock()
	defer lock.Unlock()
	time.Sleep(2 * time.Second)

	username := url.Values{"username": {mobile}}
	resp, err := http.Post(GenerateCodeUrl, "application/x-www-form-urlencoded", strings.NewReader(username.Encode()))
	if err != nil {
		return errors.Wrap(err, "巨鲲发送验证码失败提示:获取验证码链接失效")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "巨鲲发送验证码失败提示:应答解析失败")
	}

	g := &GenerateResp{}
	err = json.Unmarshal(body, g)
	if err != nil {
		return errors.Wrap(err, "巨鲲发送验证码失败提示:反序列化失败")
	}

	if (g.Msg != "验证码已经发送成功!") && (g.Msg != "验证码有效期两小时!请勿重复获取") {
		return errors.New("巨鲲发送验证码失败提示:" + g.Msg + "--" + mobile)
	}

	return nil
}
