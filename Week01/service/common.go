package service

import "sync"

type Application interface {
	RegisterWithMobile(mobile string, code string) error
	GenerateCode(mobile string, lock sync.Mutex) error
}
