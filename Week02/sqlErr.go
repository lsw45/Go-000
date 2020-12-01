package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Row struct {
	id   int
	name string
}

var (
	errNoRow  = errors.New("no row")
	errUpdate = errors.New("Update error")
	errSelect = errors.New("select error")
)

func daoSelect(id int) ([]Row, error) {
	return []Row{{1, "sr"}}, errors.New("id " + strconv.Itoa(id) + " " + errNoRow.Error())
}
func daoUpdate() (int64, error) {
	return 2, errors.New("id " + strconv.Itoa(2) + " " + errNoRow.Error())
}

// 为了更好的报错给调用层，dao层的select和update可以做个wrap给service层
func main() {
	list, err := daoSelect(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result are %+v", list)

	_, err = daoUpdate()
	if err != nil {
		log.Fatal(err)
	}
}
