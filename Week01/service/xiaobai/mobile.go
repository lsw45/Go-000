package xiaobai

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var token = "69d22f4efff4879dc84415c34bdd2aba"
var Success = 20000

var GetMobileUrl = "http://api.xiaobai188.com:188/api/getPhone?token=" + token + "&sid=163260"

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

func GetMobile() (mobile string, err error) {
	resp, err := http.Get(GetMobileUrl)
	if err != nil {
		return "", errors.Wrap(err, "get mobile fail：获取手机号链接失效")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "get mobile fail：应答解析失败")
	}

	m := &MobileResp{}
	err = json.Unmarshal(body, m)
	if err != nil {
		return "", errors.Wrap(err, "get mobile fail：反序列化失败")
	}

	if m.Code != Success {
		return "", errors.Wrapf(err, "get mobile fail：%s", m.Msg)
	}
	return m.Data[0].Phone, nil
}
