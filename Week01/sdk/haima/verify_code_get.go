package haima

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
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

func GetCode(mobile, pid string) (code string, err error) {
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
			code, err = getCode(pid)
			if err == nil && len(code) > 0 {
				return code, nil
			}
		}

		// 间隔1s获取一次
		time.Sleep(time.Second * 1)
	}
	return "", TimeOutErr
}

func getCode(pid string) (code string, err error) {

	url := "http://api.hmyzm.cool:88/yhapi.ashx?act=getPhoneCode&token=" + token + "&pid=" + pid
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

	//成功返回：1|验证码数字|完整短信内容
	result := strings.Split(string(body), "|")

	if result[0] != success {
		return "", errors.New("海马获取验证码失败:" + getCodeErr[result[1]])
	}

	return result[1], nil
}
