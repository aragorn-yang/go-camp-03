package domain

import "context"

// User 用户
type User struct {
	ID   int
	Name string
	City string
}

type IUserUsecase interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}

type IUserRepo interface {
	GetUserInfo(ctx context.Context, id int) (*User, error)
}
