package model

var GetCodeUrl = "http://api.do889.com:81/api/get_message?token=" +
	"MOFnMF7LMkIJwLBBWezpz4D2oqU4VFISw1kIPKzFcQdc8hA2HZpVSOiYDIEQVy924rhIbKVGwxvA9JHuGSGIy+Hz1pl4YODb3Ihngf996Br2+PTJMFns9Kyr8Ck6JD0RlGkno//dwyxrObOXUAqrUIjfDP8XJtbhiqXXGYPGKVg=&project_id=12734&phone_num=%s"

type Autogenerated struct {
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

/*
{
    "message": "短信还未到达,请继续获取",
    "data": []
}
*/
