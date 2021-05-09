package defu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type CodeResp struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    []struct {
		ProjectID   string `json:"project_id"`
		Modle       string `json:"modle"`
		Phone       string `json:"phone"`
		ProjectType string `json:"project_type"`
	} `json:"data"`
}

func GetCode(mobile, sid string) (code string, err error) {
	url := "http://api.do889.com:81/api/get_message"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(token+";phone_num=%s;project_type=1;project_id=%s", mobile, sid))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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
	a := &CodeResp{}

	err = json.Unmarshal(body, a)
	if err != nil {
		return "", errors.Wrap(err, "德芙获取验证码失败:反序列化失败")
	}
	if a.Message != success {
		return "", errors.New("德芙获取验证码失败:" + a.Message)
	}
	return a.Code, nil
}
