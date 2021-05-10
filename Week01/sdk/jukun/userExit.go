package jukun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type exitResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func UserExit(mobile string) (exits bool, err error) {

	url := "https://h5api.jukunwang.com/api/user/public/activeuser"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("username=%s&old_pass=85404344&new_pass=55555544&new_pay_pass=22222244&que_id=3&answer=x", mobile))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	resp := exitResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	if resp.Code == 0 && resp.Msg == "用户不存在!" {
		return false, nil
	}
	return true, nil
}
