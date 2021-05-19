package fil

import (
	"encoding/json"
	"fmt"
	"github.com/jin-Register/service"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func SendCode(mobile string, mut sync.Mutex) (err error) {

	url := "https://api.tcstar.cn/v3/Sms/send_reg_sms"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"phone":%s,"invitation_code":"Y690753","t":"Uc644165"}`, mobile))

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

	r := &RegisterResp{}
	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(body)
		return
	}

	if r.Code != 200 || r.Msg != "发送成功!" {
		service.LogPhone.Errorf("账号:%s,验证码发送异常:%s", mobile, r.Msg)
		return errors.New(r.Msg)
	}

	return nil
}
