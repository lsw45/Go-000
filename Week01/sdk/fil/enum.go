package fil

var invitation_code = "Y690753"
var register = `{"invitation_code":"` + invitation_code + `","login_pwd":"filcoin123","login_repwd":"filcoin123","pay_pwd":"filcoin123","t":"Uc644165",`

type GetUser struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Name     string `json:"name"`
		Mobile   string `json:"mobile"`
		Age      int    `json:"age"`
		ID       string `json:"id"`
		Birthday string `json:"birthday"`
		Sex      int    `json:"sex"`
		AreaNam  string `json:"area_nam"`
	} `json:"data"`
}

//http://106.14.127.87:8011/getUser
