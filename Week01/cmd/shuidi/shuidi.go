package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var registerUrl = "https://shuidi.sdf2ggggg.cn/index.php/Home/Public/reg.html"
var param = url.Values{
	"password":    {"854043"},
	"invite_code": {"649230"},
}

//mobile=13515503075&password=854043&invite_code=649230
//userAgent=1; BJYADMIN=ac2i89jbq9p3pf35gg8rsu4hv5

type response struct {
	Code string `json:"code"`
}

var successResp = "1"

func main() {
	InitLog()

	var mobile string

	if len(mobile) == 0 {
		return
	}
	param.Add("mobile", mobile)

	resp, err := http.Post(registerUrl, "application/x-www-form-urlencoded", strings.NewReader(param.Encode()))
	if err != nil {
		logrus.Errorf("调用注册链接失败:%s,error:%s", mobile, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("读取失败:%s,error:%s", mobile, err)
		return
	}
	result := &response{}

	err = json.Unmarshal(body, result)
	if err != nil {
		logrus.Errorf("反序列化失败:%s,error:%s", mobile, err)
		return
	}

	if result.Code != successResp {
		logWarn
	} else {
		logPhone
	}

}

var logPhone *logrus.Logger
var logWarn *logrus.Logger

func InitLog() {
	log, _ := os.OpenFile("./log/error.txt", os.O_APPEND, 0666)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2021-04-29 00:31:43",
	})
	logrus.SetOutput(log)

	logPhone = logrus.New()
	phone, _ := os.OpenFile("./log/phone.txt", os.O_APPEND, 0666)
	logPhone.SetOutput(phone)

	logWarn = logrus.New()
	warn, _ := os.OpenFile("./log/warn.txt", os.O_APPEND, 0666)
	logWarn.SetOutput(warn)
}
