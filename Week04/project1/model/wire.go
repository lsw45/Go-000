//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package model

import "github.com/google/wire"

func InitializeSpeaker(username, password string) UserInfo {
	wire.Build(NewUser, NewBase)
	return UserInfo{}
}
