package test

import (
	"github.com/golang/sync/errgroup"
	"testing"
	"time"
)

type ABC struct {
	CBA int
}

func TestNormal(t *testing.T) {
	var (
		abcs = make(map[int]*ABC)
		g    errgroup.Group
		err  error
	)
	for i := 0; i < 10; i++ {
		abcs[i] = &ABC{CBA: i}
	}
	g.Go(func() (err error) {
		abcs[1].CBA++
		time.Sleep(4 * time.Second)
		return
	})
	g.Go(func() (err error) {
		abcs[2].CBA++
		time.Sleep(2 * time.Second)
		return
	})
	if err = g.Wait(); err != nil {
		t.Log(err)
	}
	t.Log(abcs[1])
}
