package adapters

import (
	"context"
	"errors"
	"time"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/shortUrl"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type shortUrlMongoRepository struct {
	client *mongo.Client
}

func NewShortUrlMongoRepository(client *mongo.Client) shortUrlMongoRepository {
	return shortUrlMongoRepository{
		client: client,
	}
}

type ShortUrlModel struct {
	ID          string    `bson:"_id"`
	OriginalUrl string    `bson:"original_url"`
	UserID      string    `bson:"user_id"`
	CreatedAt   time.Time `bson:"createdAt"`
}

func (r shortUrlMongoRepository) marshalShortUrl(u *shortUrl.ShortUrl) ShortUrlModel {
	return ShortUrlModel{
		ID:          u.ID(),
		OriginalUrl: u.OriginalUrl(),
		UserID:      u.UserID(),
		CreatedAt:   u.CreatedAt(),
	}
}
func (r shortUrlMongoRepository) unmarshalShortUrl(um ShortUrlModel) *shortUrl.ShortUrl {
	return shortUrl.UnmarshalFromDatabase(um.ID, um.OriginalUrl, um.UserID, um.CreatedAt)
}

func (r shortUrlMongoRepository) shortUrlCollection() *mongo.Collection {
	return r.client.Database(viper.GetString("mongo.db")).Collection("shortUrl")
}

func (r shortUrlMongoRepository) ShorteningUrl(ctx context.Context, url *shortUrl.ShortUrl) error {
	model := r.marshalShortUrl(url)
	_, err := r.shortUrlCollection().InsertOne(ctx, model)
	return err
}

func (r shortUrlMongoRepository) GetOriginalLink(ctx context.Context, id string) (string, error) {
	filter := bson.M{"_id": id}
	var um ShortUrlModel
	err := r.shortUrlCollection().FindOne(ctx, filter).Decode(&um)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return "", appError.ErrorNotFound{
			Key: "shortUrl",
			Err: err,
		}
	} else if err != nil {
		return "", err
	}

	return um.OriginalUrl, nil
}
