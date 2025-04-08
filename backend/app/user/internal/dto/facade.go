package dto

import "sync"

var (
	userCv   *UserConverter
	userOnce sync.Once
)

func NewUserConverter() *UserConverter {
	userOnce.Do(func() {
		userCv = new(UserConverter)
	})
	return userCv
}
