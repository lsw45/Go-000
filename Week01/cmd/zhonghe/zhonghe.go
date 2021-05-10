package main

import (
	"fmt"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/zhonghe"
	"os"
	"sync"
)

var debug = false
var count = 5
var sid = "65429"

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/zhonghe")

	var mut sync.Mutex

	var zhongStr = zhonghe.NewZhonghe(sid, "", "", mut)

	for i := 0; i < count; i++ {
		service.Start(zhongStr)
	}

	fmt.Println("本次批量注册任务完成")
}
