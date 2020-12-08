package main

import (
	"fmt"
	"github.com/golang/sync/errgroup"
	"log"
	"net/http"
)

func main() {
	var (
		err error
	)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello")
	})

	g := errgroup.Group{}

	g.Go(func() (err error) {
		if err = http.ListenAndServe(":8080", nil); err != nil {
			return err
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		log.Println(err)
	}
}
