package users

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(*UserRegisterRequest) (*UserModel, error)
	GetUser(*string) (*UserModel, error)
}

type UserService struct {
	collection *mongo.Collection
	ctx        context.Context
}

/** Constructors */
func NewService(_collection *mongo.Collection, _ctx context.Context) IUserService {
	return &UserService{
		collection: _collection,
		ctx:        _ctx,
	}
}

func (u *UserService) CreateUser(payload *UserRegisterRequest) (*UserModel, error) {

	user := new(UserModel)
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

func (u *UserService) GetUser(user *string) (*UserModel, error) {
	return nil, nil
}
