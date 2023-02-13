package adapters

import (
	"context"
	"errors"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/user"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userMongoRepository struct {
	client *mongo.Client
}

func NewUserMongoRepository(client *mongo.Client) userMongoRepository {
	repo := userMongoRepository{
		client: client,
	}

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := repo.userCollection().Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		defer recover()
		panic(err)
	}

	return repo
}

type UserModel struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (r userMongoRepository) marshalUser(u *user.User) UserModel {
	return UserModel{
		ID:   u.ID(),
		Name: u.Name(),
	}
}
func (r userMongoRepository) unmarshalUser(um UserModel) *user.User {
	return user.UnmarshalFromDatabase(um.ID, um.Name)
}

func (r userMongoRepository) userCollection() *mongo.Collection {
	return r.client.Database(viper.GetString("mongo.db")).Collection("user")
}

func (r userMongoRepository) CreateUser(ctx context.Context, u *user.User) error {
	model := r.marshalUser(u)

	_, err := r.userCollection().InsertOne(ctx, model)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return appError.ErrorAlreadyExists{
				Key: "user",
				Err: err,
			}
		}
		return err
	}

	return nil
}

func (r userMongoRepository) GetUser(ctx context.Context, name string) (*user.User, error) {
	filter := bson.M{"name": name}
	model := new(UserModel)
	err := r.userCollection().FindOne(ctx, filter).Decode(model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, appError.ErrorNotFound{
				Key: "user",
				Err: err,
			}
		}
		return nil, err
	}

	return r.unmarshalUser(*model), nil
}
