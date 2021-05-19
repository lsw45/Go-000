package fil

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"io/ioutil"
	"net/http"
	"strings"
)

type RegisterResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Register(name, idCard, mobile, code string) (err error) {

	url := "https://api.tcstar.cn/v3/reg"
	method := "POST"
	str := register + fmt.Sprintf(`"real_name":%s","id_card":%s","login_name":%s","code":%s"}`, name, idCard, mobile, code)
	payload := strings.NewReader(str)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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
		fmt.Println(err)
		return
	}

	if r.Code != 200 {
		service.LogPhone.Errorf("账号:%s,注册异常:%s", mobile, r.Msg)
	} else {
		service.LogPhone.Infof("新账号:%s,注册成功", mobile)
	}

	return nil
}
