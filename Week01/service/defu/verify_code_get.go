package defu

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CodeResp struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    []struct {
		ProjectID   string `json:"project_id"`
		Modle       string `json:"modle"`
		Phone       string `json:"phone"`
		ProjectType string `json:"project_type"`
	} `json:"data"`
}

var TimeOutErr = errors.New("get code timeout")

func GetCode(mobile, sid string) (code string, err error) {
	time.Sleep(3 * time.Second)

	var retry = 1
	var timeout = time.After(20 * time.Second)

	for err != nil || len(code) == 0 {
		select {
		case <-timeout:
			service.LogPhone.Errorf("德芙验证码获取失败,mobile:%s,retry:%d,error:%s", mobile, retry, err)
			return "", TimeOutErr
		default:
			retry++
			code, err = getCode(mobile, sid)
			if err == nil && len(code) > 0 {
				return code, nil
			}
		}

		// 间隔1s获取一次
		time.Sleep(time.Second * 1)
	}
	return "", TimeOutErr
}

func getCode(mobile, sid string) (code string, err error) {
	urls := "http://api.do889.com:81/api/get_message"
	method := "POST"

	reqBody := url.QueryEscape(fmt.Sprintf("token="+token+";phone_num=%s;project_type=1;project_id=%s", mobile, sid))
	payload := strings.NewReader(reqBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, urls, payload)

	if err != nil {
		logrus.Error(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	a := &CodeResp{}
	err = json.Unmarshal(body, a)
	if err != nil {
		return "", errors.Wrap(err, "德芙获取验证码失败:反序列化失败")
	}

	if a.Message != success {
		return "", errors.New("德芙获取验证码失败:" + a.Message)
	}

	return a.Code, nil
}
