package main

import (
	"fmt"
	"runtime"
)

func init() {

}

func main() {
	api.RegisterApi()
	router.Load()
	runtime.GOMAXPROCS(runtime.NumCPU())

}
