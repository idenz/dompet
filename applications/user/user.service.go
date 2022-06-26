package users

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Create(*RegisterRequest) (*Model, error)
	GetProfile(*string) (*Model, error)
	FindByEmail(*string) (*Model, error)
	GenerateToken(*Model) (string, error)
}

type Service struct {
	collection *mongo.Collection
	ctx        context.Context
}

/** Constructors */
func NewService(_collection *mongo.Collection, _ctx context.Context) IService {
	return &Service{
		collection: _collection,
		ctx:        _ctx,
	}
}

func (u *Service) Create(payload *RegisterRequest) (*Model, error) {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)

	user := new(Model)
	bytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, err

	}

	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}

	_, err = u.collection.InsertOne(u.ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Service) FindByEmail(email *string) (*Model, error) {
	var user *Model

	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.collection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *Service) GenerateToken(user *Model) (string, error) {

	others := make(map[string]interface{}, 0)
	token, err := user.GenerateToken(others)
	if err != nil {
		return "", err
	}
	return token, err
}

func (u *Service) GetProfile(user *string) (*Model, error) {
	return nil, nil
}
