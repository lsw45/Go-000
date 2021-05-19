package util

import (
	"github.com/jin-Register/constant"
	"math/rand"
	"strconv"
	"time"
)

type IdNumber struct {
	id          string
	area_id     int
	birth_year  int
	birth_month int
	birth_day   int
}

func NewIdNumber(id_number string) *IdNumber {
	id := &IdNumber{id: id_number}
	id.area_id, _ = strconv.Atoi(id.id[0:6])
	id.birth_year, _ = strconv.Atoi(id.id[6:10])
	id.birth_month, _ = strconv.Atoi(id.id[10:12])
	id.birth_day, _ = strconv.Atoi(id.id[12:14])
	return id
}

func GenerateId(sex int) string {
	//sex = 0表示女性，sex = 1表示男性
	rand.Seed(3470)
	m := rand.Int()
	i := 0
	id_number := ""
	for k, _ := range constant.AREA_INFO {
		i++
		if i == m {
			id_number = strconv.Itoa(k)
			break
		}
	}

	start, _ := time.ParseDuration("1960-01-01")
	end, _ := time.ParseDuration("2000-12-30")
	rand.Seed(int64(end - start))

	timeTemplate := "20060102" //其他类型
	id_number += time.Unix(int64(start)+rand.Int63(), 0).Format(timeTemplate)

	rand.Seed(89)
	id_number += strconv.Itoa(rand.Int() + 10)

}
