package jukun

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	gotToken, err := Login("18132692625", "aa419312")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gotToken)
}
