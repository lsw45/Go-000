package test

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(context context.Context, netw, addr string) (net.Conn, error) {
				//本地地址  ipaddr是本地外网IP
				lAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9006")
				if err != nil {
					return nil, err
				}
				//被请求的地址
				rAddr, err := net.ResolveTCPAddr(netw, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(netw, lAddr, rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(35 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
	resp, err := client.Post("http://5g.hhml.cn/ajax/obtain", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
}
