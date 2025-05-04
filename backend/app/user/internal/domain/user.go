package domain

import "context"

type User struct {
	Id       uint64
	Username string
	Nickname string
	Password string
	Email    string
}

type UserRepo interface {
	Save(context.Context, *User) error
	Get(context.Context, uint64) (*User, error)
	Update(context.Context, *User, func(context.Context, *User) (*User, error)) error
	GetAll(context.Context) ([]*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
}

