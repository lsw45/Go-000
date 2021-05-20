package idcard

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Name     string `json:"name"`
		Mobile   string `json:"mobile"`
		Age      int    `json:"age"`
		ID       string `json:"id"`
		Birthday string `json:"birthday"`
		Sex      int    `json:"sex"`
		AreaNam  string `json:"area_nam"`
	} `json:"data"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
	ID       string `json:"id"`
	Birthday string `json:"birthday"`
	Sex      int    `json:"sex"`
	AreaNam  string `json:"area_nam"`
}

func GetUserInfo() (u UserInfo, err error) {
	url := "http://106.14.127.87:8011/getUser"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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

	r := &Response{}
	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if r.Code != 200 {
		return u, errors.New(r.Msg)
	}

	return r.Data, nil
}
