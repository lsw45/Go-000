package main

import (
	"fmt"
	"github.com/jin-Register/sdk/haima"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/jukun"
	"os"
	"sync"
)

var count = 1

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/jukun")

	var mut sync.Mutex
	var juKunStr = jukun.NewJukun(haima.IidJukun, mut)

	for i := 0; i < count; i++ {
		service.Start(juKunStr)
	}

	fmt.Println("本次批量注册任务完成")
}
