package users

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(*UserModel) error
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

func (u *UserService) CreateUser(user *UserModel) error {
	return nil
}

func (u *UserService) GetUser(user *string) (*UserModel, error) {
	return nil, nil
}
