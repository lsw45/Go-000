package model

var RegisterUrl = "https://h5api.jukunwang.com/api/user/public/register"

var UserToken = map[string]string{"xx-token": ""}

var RegisterSuccess = 1

var RegisterReq = RegisterRequest{
	Password: 854043,
	PayPass:  854043,
	Invcode:  "54NNM",
	QueID:    3,
	Answer:   "狗",
}

type RegisterRequest struct {
	Username         int64  `json:"username"`
	Password         int    `json:"password"`
	PayPass          int    `json:"pay_pass"`
	VerificationCode string `json:"verification_code"`
	Invcode          string `json:"invcode"`
	QueID            int    `json:"que_id"`
	Answer           string `json:"answer"`
}

type RegisterResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

/*
{
    "code": 0,
    "msg": "此账号已存在!",
    "data": ""
}
{
    "code": 1,
    "msg": "验证码已经发送成功!",
    "data": ""
}
*/
