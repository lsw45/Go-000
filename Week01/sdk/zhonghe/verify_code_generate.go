package zhonghe

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

var username = url.Values{"command": {"getVcode"}}

type GenerateResp struct {
	Ret bool   `json:"ret"`
	Msg string `json:"msg"`
}

/*
{
    "ret":true,
    "msg":"发送成功"
}
*/

func GenerateCode(mobile string, lock sync.Mutex) error {
	lock.Lock()
	defer lock.Unlock()
	time.Sleep(2 * time.Second)

	username.Add("mobile", mobile)

	resp, err := http.Post(GenerateCodeUrl, "application/x-www-form-urlencoded", strings.NewReader(username.Encode()))
	if err != nil {
		return errors.Wrap(err, "众和发送验证码失败提示:获取验证码链接失效")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "众和发送验证码失败提示:应答解析失败")
	}

	g := &GenerateResp{}
	err = json.Unmarshal(body, g)
	if err != nil {
		return errors.Wrap(err, "众和发送验证码失败提示:反序列化失败")
	}

	if ((g.Msg != "发送成功") && (g.Msg != "发送频繁，请稍后再试")) || g.Ret != Success {
		return errors.New("众和发送验证码失败提示:" + g.Msg + "--" + mobile)
	}

	return nil
}
