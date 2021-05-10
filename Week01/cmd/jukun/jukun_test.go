package main

import (
	"sync"
	"testing"
	"time"
)

func TestGetCode(t *testing.T) {
	debug = true
	var lock sync.Mutex
	register(lock)
	t.Parallel()
	time.Now()
}
