package model

var GenerateCodeUrl = "https://h5api.jukunwang.com/api/user/Verification_Code/send"

var GenerateReq = map[string]string{"username": ""}

var GenerateSuccess = 1

type GenerateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

/*
{
    "code": 1,
    "msg": "验证码已经发送成功!",
    "data": ""
}
*/
