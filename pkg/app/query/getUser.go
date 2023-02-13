package query

import (
	"context"
	"fmt"
	"log"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/user"
)

type GetUserHandler struct {
	readModel GetUserReadModel
	logger    *log.Logger
}

type GetUserReadModel interface {
	GetUser(ctx context.Context, name string) (*user.User, error)
}

func NewGetUserHandler(readModel GetUserReadModel, logger *log.Logger) GetUserHandler {
	if readModel == nil {
		defer recover()
		panic("nil readModel")
	}

	return GetUserHandler{
		readModel: readModel,
		logger:    logger,
	}
}

func (h GetUserHandler) Handle(ctx context.Context, name string) (*user.User, error) {
	u, err := h.readModel.GetUser(ctx, name)
	if err != nil {
		if _, has := err.(appError.AppError); !has {
			err = fmt.Errorf("error getting user: %w", err)
			h.logger.Println(err)
			return nil, err
		}
		return nil, err
	}

	return u, nil
}
