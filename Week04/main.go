package main

import (
	"fmt"
	"plugin"
)

func main() {
	/*	p := pingo.NewPlugin("tcp", "./discover/discover.exe")
		// 启动插件
		p.Start()
		// 使用完插件后停止它
		defer p.Stop()
		var resp string
		// 从先前创建的对象调用函数
		if err := p.Call("Has.Discovery", nil, &resp); err != nil {
			log.Print(err)
		} else {
			log.Print(resp)
		}*/
	d, err := plugin.Open("discover/discovery.so")
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := d.Lookup("Discovery")
	if err != nil {
		fmt.Println(err)
		return
	}
	ok, has := s.(func())
	if has {
		ok()
	} else {
		fmt.Println("plugin is error")
	}
}
