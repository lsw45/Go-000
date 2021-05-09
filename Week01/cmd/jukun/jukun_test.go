package main

import (
	"sync"
	"testing"
)

func TestGetCode(t *testing.T) {
	debug = true
	var lock sync.Mutex
	register(lock)
}
