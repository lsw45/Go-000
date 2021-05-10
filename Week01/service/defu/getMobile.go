package defu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var token = `nqjOsBAnR2DTNXn59p+EYUh1/gsloEkFRM7hqRBIVqasKDAnXFCuRwhgwTApL2nAAMr3e738Wzf3ppNm72uoA10X3vO9/PSqMCXlDGxcHxbRrR1SdWqFucprwe4kzheJiVtJhxQwDtI/btODh6HuEAup2nB979+SjfHAw6CI19Y=`
var success = "ok"

type MobileResp struct {
	Message string        `json:"message"`
	Mobile  string        `json:"mobile"`
	One     string        `json:"1分钟内剩余取卡数:"`
	Data    []interface{} `json:"data"`
}

func GetMobile(sid string) (mobile string, err error) {
	url := "http://api.do889.com:81/api/get_mobile?token=" + fmt.Sprintf(token+"&project_id=%s", sid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
	m := &MobileResp{}
	err = json.Unmarshal(body, m)
	if err != nil {
		return "", errors.Wrap(err, "德芙获取手机号失败:反序列化失败")
	}

	if m.Message != success {
		return "", errors.Wrapf(err, "德芙获取手机号失败:%s", m.Message)
	}

	return m.Mobile, nil
}
