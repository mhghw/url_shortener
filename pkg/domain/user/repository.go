package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, u *User) error
	GetUser(ctx context.Context, name string) (*User, error)
}
