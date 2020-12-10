package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	BaseModel
	Username string `json:"username"`
	SayHello string `json:"sayHello"`
	Password string `json:"password"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

func NewUser(username, password string, base BaseModel) UserInfo {
	return UserInfo{
		BaseModel: base,
		Username:  username,
		Password:  password,
	}
}

func NewBase() BaseModel {
	return BaseModel{
		CreatedAt: time.Now(),
		DeletedAt: nil,
	}
}
