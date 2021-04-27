package model

var GetMobileUrl = "http://api.do889.com:81/api/get_mobile?token=MOFnMF7LMkIJwLBBWezpz4D2oqU4VFISw1kIPKzFcQdc8hA2HZpVSOiYDIEQVy924rhIbKVGwxvA9JHuGSGIy+Hz1pl4YODb3Ihngf996Br2+PTJMFns9Kyr8Ck6JD0RlGkno//dwyxrObOXUAqrUIjfDP8XJtbhiqXXGYPGKVg=&project_id=11428"

type MobileResp struct {
	Message string        `json:"message"`
	Mobile  string        `json:"mobile"`
	One     string        `json:"1分钟内剩余取卡数:"`
	Data    []interface{} `json:"-"`
}

var MobileSuccess = "ok"

/*
{
    "message": "ok",
    "mobile": "16725569389",
    "1分钟内剩余取卡数:": "299",
    "data": []
}
*/
