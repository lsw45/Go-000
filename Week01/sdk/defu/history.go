package defu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ResponseList struct {
	Message string `json:"message"`
	Count   string `json:"count"`
	Sum     string `json:"sum"`
	Data    []struct {
		ID          string `json:"id"`
		ProjectID   string `json:"project_id"`
		Modle       string `json:"modle"`
		Phone       string `json:"phone"`
		Money       string `json:"money"`
		Time        string `json:"time"`
		ProjectType string `json:"project_type"`
		Version     string `json:"version"`
		UserID      string `json:"user_id"`
	} `json:"data"`
}

var postUrl = "http://api.do889.com:81/api/get_expenditure"
var prefix = "验证码是"
var suffix = "。"

func QueryHistory(phone, sid string) (string, error) {
	token := url.Values{"token": {token}}
	resp, err := http.Post(postUrl, "application/x-www-form-urlencoded", strings.NewReader(token.Encode()))
	if err != nil {
		return "", errors.Wrap(err, "验证码列表:请求失败")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "验证码列表:应答解析失败")
	}

	m := &ResponseList{}
	err = json.Unmarshal(body, m)
	if err != nil {
		return "", errors.Wrap(err, "验证码列表:反序列化失败")
	}

	if m.Message != "ok" {
		return "", errors.Wrapf(err, "验证码列表:%s", m.Message)
	}

	for _, item := range m.Data {
		if item.Phone == phone && item.ProjectID == sid {
			fmt.Println(strings.Index(item.Modle, prefix))
			fmt.Println(strings.Index(item.Modle, suffix))
			return item.Modle[strings.Index(item.Modle, prefix):strings.Index(item.Modle, suffix)], nil
		}
	}
	return "", nil
}
