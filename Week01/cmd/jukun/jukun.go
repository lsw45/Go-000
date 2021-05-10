package main

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/jukun"
	"os"
	"sync"
)

/*
1、获取手机号
2、金巨鲲发送验证码
3、获取验证码
4、提交注册信息
*/
var debug = false
var count = 5
var sid = "163260"

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/jukun")

	var mut sync.Mutex
	var juKunStr = jukun.NewJukun(sid, "", "", mut)

	for i := 0; i < count; i++ {
		service.Start(juKunStr)
	}

	fmt.Println("本次批量注册任务完成")
}
