package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type Message bool

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		return
	}
	defer conn.Close()

	msg := make(chan Message)

	go func() {
		for i := 0; i < 30; i++ {
			_, err := io.WriteString(conn, "hello")
			if err != nil {
				fmt.Println("write string failed", err)
				return
			}

			msg <- true
			if i == 25 {
				close(msg)
			}
		}
	}()

	buf := make([]byte, 10)

	go func() {
		for {
			_, ok := <-msg
			if !ok {
				return
			}

			count, err := conn.Read(buf)
			if err != nil {
				break
			}

			fmt.Println(string(buf[0:count]))
		}
	}()

	time.Sleep(20 * time.Second)
}
