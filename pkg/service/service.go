package service

import (
	"context"
	"log"
	"urlShortener/pkg/adapters"
	"urlShortener/pkg/app"
	"urlShortener/pkg/app/command"
	"urlShortener/pkg/app/query"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewApplication(ctx context.Context, logger *log.Logger) app.Application {
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := adapters.NewUserMongoRepository(mongoClient)
	shortUrlRepo := adapters.NewShortUrlMongoRepository(mongoClient)

	return app.Application{
		Commands: app.Command{
			CreateShortLink: command.NewCreateShortLinkHandler(userRepo, shortUrlRepo, logger),
			Register:        command.NewRegisterHandler(userRepo, logger),
		},

		Queries: app.Query{
			GetOriginalLink: query.NewGetOriginalLinkHandler(shortUrlRepo, logger),
			GetUser:         query.NewGetUserHandler(userRepo, logger),
		},
	}
}
