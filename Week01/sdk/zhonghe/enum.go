package zhonghe

import "net/url"

var Success = true

var RegisterUrl = "http://5g.hhml.cn/ajax/operate"

var GenerateCodeUrl = "http://5g.hhml.cn/ajax/obtain"

var registerReq = url.Values{"command": {"register"}}

var field = `{"mobile":"%s","vcode":"%s","password":"zhonghe123456","repassword":"zhonghe123456","up_code":"80d6af37a0"}`

// 亚军的邀请码：80d6af37a0
//361f428132
