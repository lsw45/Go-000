package defu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type MobileResp struct {
	Message string        `json:"message"`
	Mobile  string        `json:"mobile"`
	One     string        `json:"1分钟内剩余取卡数:"`
	Data    []interface{} `json:"data"`
}

func GetMobile(sid, operator string) (mobile string, err error) {
	url := "http://api.do889.com:81/api/get_mobile?token=" + token + "&project_id=" + sid
	if len(operator) == 1 {
		url += "&operator=" + operator
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := &MobileResp{}
	err = json.Unmarshal(body, m)
	if err != nil {
		return "", errors.Wrap(err, "德芙获取手机号失败:反序列化失败")
	}

	if m.Message != success {
		return "", errors.New("德芙获取手机号失败:" + m.Message)
	}

	return m.Mobile, nil
}
