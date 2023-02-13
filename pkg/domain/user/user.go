package user

import (
	"errors"
	appError "urlShortener/error"

	"github.com/google/uuid"
)

type User struct {
	id   string
	name string
}

func (u User) ID() string {
	return u.id
}
func (u User) Name() string {
	return u.name
}

func NewUser(name string) (*User, error) {
	if len(name) < 3 || len(name) > 20 {
		return nil, appError.ErrorInvalidArguments{
			Key: "user",
			Err: errors.New("name length must be greater than 2 and less than 20"),
		}
	}

	return &User{
		id:   uuid.New().String(),
		name: name,
	}, nil
}

func UnmarshalFromDatabase(
	id string,
	name string,
) *User {
	return &User{
		id,
		name,
	}
}
