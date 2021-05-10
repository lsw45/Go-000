package xiaobai

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var token = "da18d2400cac388187dbcd3b22a36d1d"
var Success = 20000

var GetMobileUrl = "http://api.xiaobai188.com:188/api/getPhone?token=" + token + "&sid="

type MobileResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Phone    string `json:"phone"`
		Province string `json:"province"`
		City     string `json:"city"`
		Server   string `json:"server"`
		Haoduan  string `json:"haoduan"`
	} `json:"data"`
}

/*
{
    "code": 20000,
    "msg": "Success",
    "data": [
        {
            "phone": "17041251765",
            "province": "河北",
            "city": "唐山",
            "server": "民生通讯",
            "haoduan": "1704125"
        }
    ]
}
*/

func GetMobile(sid string) (mobile string, err error) {
	resp, err := http.Get(GetMobileUrl + sid)
	if err != nil {
		return "", errors.Wrap(err, "小白获取手机号失败:获取手机号链接失效")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "小白获取手机号失败:应答解析失败")
	}

	m := &MobileResp{}
	err = json.Unmarshal(body, m)
	if err != nil {
		return "", errors.Wrap(err, "小白获取手机号失败:反序列化失败")
	}

	if m.Code != Success {
		return "", errors.Wrapf(err, "小白获取手机号失败:%s", m.Msg)
	}
	return m.Data[0].Phone, nil
}
