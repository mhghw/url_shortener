package query

import (
	"context"
	"fmt"
	"log"
	appError "urlShortener/error"
)

type GetOriginalLinkHandler struct {
	readModel GetOriginalLinkReadModel
	logger    *log.Logger
}

type GetOriginalLinkReadModel interface {
	GetOriginalLink(ctx context.Context, id string) (string, error)
}

func NewGetOriginalLinkHandler(readModel GetOriginalLinkReadModel, logger *log.Logger) GetOriginalLinkHandler {
	if readModel == nil {
		defer recover()
		panic("nil readModel")
	}

	return GetOriginalLinkHandler{
		readModel: readModel,
		logger:    logger,
	}
}

func (h GetOriginalLinkHandler) Handle(ctx context.Context, id string) (string, error) {
	shortedLink, err := h.readModel.GetOriginalLink(ctx, id)
	if err != nil {
		if _, has := err.(appError.AppError); !has {
			err = fmt.Errorf("error getting shortUrl: %w", err)
			h.logger.Println(err)
			return "", err
		}
		return "", err
	}

	return shortedLink, nil
}
