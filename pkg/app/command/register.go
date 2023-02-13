package command

import (
	"context"
	"fmt"
	"log"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/user"
)

type RegisterHandler struct {
	userRepo user.Repository
	logger   *log.Logger
}

func NewRegisterHandler(ur user.Repository, logger *log.Logger) RegisterHandler {
	handler := RegisterHandler{
		userRepo: ur,
		logger:   logger,
	}
	if handler.userRepo == nil {
		defer recover()
		panic("nil user repository")
	}

	return handler
}

func (h RegisterHandler) Handle(ctx context.Context, name string) error {
	u, err := user.NewUser(name)
	if err != nil {
		return err
	}

	err = h.userRepo.CreateUser(ctx, u)
	if err != nil {
		if _, has := err.(appError.AppError); !has {
			err = fmt.Errorf("error creating user: %w", err)
			h.logger.Println(err)
			return err
		}
		return err
	}

	return nil
}
