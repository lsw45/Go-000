package main

import (
	"context"
	"fmt"
	"github.com/golang/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		err  error
		stop = make(chan os.Signal, 1)
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	server := http.Server{Addr: ":8080", Handler: mux}

	g := errgroup.Group{}

	g.Go(func() (err error) {
		if err = server.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})

	signal.Notify(stop, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
	g.Go(func() error {
		c := <-stop
		fmt.Printf("Got signal:%v\n", c)
		if err = server.Shutdown(context.TODO()); err != nil {
			return err
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		log.Println(err)
	}
}
