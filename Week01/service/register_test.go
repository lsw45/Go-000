package service

import "testing"

var MobileAndCode = map[string]string{
	"17056449656": "",
	"17036663982": "435373",
	"16725568365": "",
}

func TestRegisterWithMobile(t *testing.T) {
	err := RegisterWithMobile("17036663982", MobileAndCode["17036663982"])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestGetMobile(t *testing.T) {
	mobile, err := GetMobile()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mobile)
}
