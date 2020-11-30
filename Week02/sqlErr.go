package Week02

import (
	"github.com/pkg/errors"
	"log"
)

//

type Row struct {
	id   int
	name string
}

var errNoRow = errors.New("no row")

func dao() (list []Row, err error) {

	return nil, nil
}

func main() {
	list, err := dao()
	if err != nil {

	}
	log.Printf("result are %+v", list)
}
