package xingyuan

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type RegisterResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Register(mobile, smsCode, code string) (err error) {

	url := "http://xyb.wiifood.cn//home/login/reg/u/052471.html"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(register, mobile, smsCode, code))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

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

	r := &RegisterResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(body)
		return
	}

	if r.Code != 0 {
		service.LogPhone.Errorf("账号:%s,注册异常:%s", mobile, r.Msg)
		return errors.New(r.Msg)
	}

	return nil
}
