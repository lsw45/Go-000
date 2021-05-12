package jukun

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type changeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func ChangePasswd(mobile, code string) (err error) {
	url := "https://h5api.jukunwang.com/api/user/public/passwordforget"
	method := "POST"

	raw := fmt.Sprintf("username=%s&user_pass=%s&verification_code=%s", mobile, Passwd, code)
	payload := strings.NewReader(raw)

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

	resp := changeResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	if resp.Code != 0 && resp.Msg != "登录密码已修改!" {
		return errors.New("修改登录密码失败:" + resp.Msg)
	}
	return nil
}
