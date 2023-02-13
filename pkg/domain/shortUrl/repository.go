package shortUrl

import "context"

type Repository interface {
	ShorteningUrl(ctx context.Context, url *ShortUrl) error
	GetOriginalLink(ctx context.Context, id string) (string, error)
}
