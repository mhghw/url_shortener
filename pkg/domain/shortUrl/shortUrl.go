package shortUrl

import (
	"fmt"
	"net/url"
	"time"
	appError "urlShortener/error"

	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
)

type ShortUrl struct {
	id          string
	originalUrl string
	userID      string
	createdAt   time.Time
}

func (u ShortUrl) ID() string {
	return u.id
}
func (u ShortUrl) OriginalUrl() string {
	return u.originalUrl
}
func (u ShortUrl) UserID() string {
	return u.userID
}
func (u ShortUrl) FullShortUrl() string {
	return fmt.Sprintf("%v/%v", viper.GetString("baseurl"), u.id)
}
func (u ShortUrl) CreatedAt() time.Time {
	return u.createdAt
}

func NewShortUrl(originalUrl, userID string) (*ShortUrl, error) {
	if _, err := url.ParseRequestURI(originalUrl); err != nil {
		return nil, appError.ErrorInvalidArguments{
			Key: "shortUrl",
			Err: fmt.Errorf("invalid original url: %w", err),
		}
	}

	return &ShortUrl{
		id:          shortid.MustGenerate(),
		originalUrl: originalUrl,
		userID:      userID,
		createdAt:   time.Now().UTC(),
	}, nil
}

func UnmarshalFromDatabase(
	id string,
	originalUrl string,
	userID string,
	createdAt time.Time,
) *ShortUrl {
	return &ShortUrl{
		id:          id,
		originalUrl: originalUrl,
		userID:      userID,
		createdAt:   createdAt,
	}
}
