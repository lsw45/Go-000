package jukun

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func ForgetPass(mobile, code, token string) (err error) {

	url := "https://h5api.jukunwang.com/api/user/public/paypssforget"
	method := "POST"
	raw := fmt.Sprintf("username=%s&pay_pass=%S&verification_code=%s", mobile, PayPasswd, code)
	payload := strings.NewReader(raw)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("xx-token", token)

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

	if resp.Code != 0 && resp.Msg != "支付密码已修改!" {
		return errors.New("修改支付密码失败:" + resp.Msg)
	}
	return nil
}
