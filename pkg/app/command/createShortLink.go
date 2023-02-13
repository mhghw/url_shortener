package command

import (
	"context"
	"fmt"
	"log"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/shortUrl"
	"urlShortener/pkg/domain/user"
)

type CreateShortLink struct {
	OriginalUrl string
	Name        string
}

type CreateShortLinkHandler struct {
	userRepo     user.Repository
	shortUrlRepo shortUrl.Repository
	logger       *log.Logger
}

func NewCreateShortLinkHandler(
	ur user.Repository,
	sr shortUrl.Repository,
	logger *log.Logger,
) CreateShortLinkHandler {
	handler := CreateShortLinkHandler{
		userRepo:     ur,
		shortUrlRepo: sr,
		logger:       logger,
	}
	if handler.userRepo == nil {
		defer recover()
		panic("nil user repository")
	}
	if handler.shortUrlRepo == nil {
		defer recover()
		panic("nil shortUrl repository")
	}

	return handler
}

func (h CreateShortLinkHandler) Handle(ctx context.Context, input CreateShortLink) (*shortUrl.ShortUrl, error) {
	u, err := h.userRepo.GetUser(ctx, input.Name)
	if err != nil {
		if _, has := err.(appError.AppError); !has {
			err = fmt.Errorf("error getting user: %w", err)
			h.logger.Println(err)
			return nil, err
		}
		return nil, err
	}

	shortLink, err := shortUrl.NewShortUrl(input.OriginalUrl, u.ID())
	if err != nil {
		return nil, err
	}

	err = h.shortUrlRepo.ShorteningUrl(ctx, shortLink)
	if err != nil {
		if _, has := err.(appError.AppError); !has {
			err = fmt.Errorf("error shortening url: %w", err)
			h.logger.Println(err)
			return nil, err
		}
		return nil, err
	}

	return shortLink, nil
}
