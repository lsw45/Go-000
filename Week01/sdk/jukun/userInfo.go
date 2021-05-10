package jukun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type LoginResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	} `json:"data"`
}

type User struct {
	ID           int    `json:"id"`
	UserStatus   int    `json:"user_status"`
	UserNickname string `json:"user_nickname"`
	Mobile       string `json:"mobile"`
	Fish         string `json:"fish"`
	Bean         string `json:"bean"`
	Invcode      string `json:"invcode"`
}

func Login(username string) (token string, err error) {
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", username)
	_ = writer.WriteField("password", Passwd)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, UserInfoUrl, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
	r := &LoginResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return "", errors.Wrap(err, "反序列化失败")
	}

	if r.Code != Success {
		service.LogPhone.Errorf("账号:%s,登陆异常:%s", username, r.Msg)
		return "", errors.New(r.Msg)
	}
	return r.Data.Token, nil
}
