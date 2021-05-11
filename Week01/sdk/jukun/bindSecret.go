package jukun

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var template = "GAXGK23XMZRHCYRZ"

type BindResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Title    string `json:"title"`
		Username string `json:"username"`
		Key      string `json:"key"`
	} `json:"data"`
}

func BindSecret(mobile, code, token string) (secret string, err error) {

	secret = "GAX" + getRandomString(13)

	url := "https://h5api.jukunwang.com/api/v1/user/bindgooglesecret"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("username=%s&user_pass=null&verification_code=%s&secret=%s&mobile=%s", mobile, code, secret, mobile))

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

	r := &BindResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return "", errors.Wrapf(err, "注册云动码失败:%s", mobile)
	}

	if r.Code != Success {
		service.LogPhone.Errorf("账号:%s,云动码注册异常:%s", mobile, r.Msg)
		return "", errors.New(r.Msg)
	} else {
		service.LogPhone.Infof("新账号:%s,云动码:%s,注册成功", mobile, secret)
	}
	return secret, nil
}

func getRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
