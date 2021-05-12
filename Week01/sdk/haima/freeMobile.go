package haima

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func FreeMobile(pid string) (err error) {

	url := "http://api.hmyzm.cool:88/yhapi.ashx?act=setRel&token=" + token + "&pid=" + pid
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

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return errors.New("free mobile done")
}
