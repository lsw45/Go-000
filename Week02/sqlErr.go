package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Row struct {
	id   int
	name string
}

var errNoRow = errors.New("no row")

func dao() (list []Row, err error) {
	err = mockErr()
	return []Row{{1, "sr"}}, err
}

func main() {
	list, err := dao()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result are %+v", list)
}

func mockErr() error {
	select {
	case <-time.After(1 * time.Second):
		return nil
	default:
		return errNoRow
	}
}
