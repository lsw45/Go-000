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

type CodeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SendCode(mobile, code string) (err error) {

	url := "http://xyb.wiifood.cn/home/utils/get_code.html"
	method := "POST"

	payload := `{"mobile":"` + mobile + `","type":"reg","code":"` + code + `"}`

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

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

	r := &CodeResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(body)
		return
	}

	if r.Code != 0 {
		service.LogPhone.Errorf("账号:%s,验证码发送异常:%s", mobile, r.Msg)
		return errors.New(r.Msg)
	}

	return nil
}
