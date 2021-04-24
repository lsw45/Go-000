package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func BenchmarkAtomic(b *testing.B) {
	cfg := Config{}
	go func() {
		i := 0
		for {
			i++
			cfg = Config{a: []int{i, i + 1, i + 2, i + 3}}
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				fmt.Printf("%+v\n", cfg)
			}
			wg.Add(-1)
		}()
	}
	wg.Wait()
}

func TestAtomic(t *testing.T) {
	var v atomic.Value
	go func() {
		i := 0
		for {
			if i > 100 {
				v.Store(&Config{a: []int{100}})
			}
			v.Store(&Config{a: []int{1}})
			i++
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 10; n++ {
		wg.Add(1)
		go func() {
			c := v.Load().(*Config)
			if c.a[0] == 100 {

			}
			wg.Add(-1)
		}()
	}
	wg.Wait()
}

func BenchmarkValueRead(b *testing.B) {
	var v atomic.Value
	v.Store(new(int))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := v.Load().(*int)
			if *x != 0 {
				b.Fatalf("wrong value: got %v, want 0", *x)
			}
		}
	})
}
