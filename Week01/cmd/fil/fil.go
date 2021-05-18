package main

import (
	"fmt"
	"github.com/jin-Register/sdk/defu"
	"github.com/jin-Register/service"
	"github.com/jin-Register/service/fil"
	"net/http"
	"os"
	"sync"
)

var count = 7

func main() {
	dir, _ := os.Getwd()
	service.InitLog(dir + "/log/zhonghe")

	var mut sync.Mutex

	var FilStr = fil.NewFilCoin(defu.IidFil, mut)

	for i := 0; i < count; i++ {
		service.Start(FilStr, http.Client{})
	}

	fmt.Println("本次批量注册任务完成")
}
