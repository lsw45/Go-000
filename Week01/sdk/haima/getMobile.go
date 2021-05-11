package haima

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetMobile(Iid string) (mobile, pid string, err error) {
	url := "http://api.hmyzm.cool:88/yhapi.ashx?act=getPhone&token=" + token + "&iid=" + Iid
	method := "GET"

	payload := strings.NewReader("")

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

	//成功返回：1|pid|提取时间|串口号|手机号|运营商|归属地
	result := strings.Split(string(body), "|")

	if result[0] != success {
		return "", "", errors.New("海马获取手机号失败:" + getMobileErr[result[1]])
	}

	return result[4], result[1], nil
}
