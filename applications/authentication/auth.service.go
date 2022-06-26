package authentication

import (
	"context"
	users "dompet/applications/user"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Login(AuthLoginRequest) (*Data, error)
}

type Service struct {
	userService users.IService
	collection  *mongo.Collection
	ctx         context.Context
}

/** Constructors */
func NewService(_userService users.IService, _collection *mongo.Collection, _ctx context.Context) IService {
	return &Service{
		collection:  _collection,
		ctx:         _ctx,
		userService: _userService,
	}
}

func (u *Service) Login(payload AuthLoginRequest) (*Data, error) {

	user, err := u.userService.FindByEmail(&payload.Email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return nil, errors.New("wrong password")
	}
	token, err := u.userService.GenerateToken(user)

	result := &Data{
		Token: token,
		Email: user.Email,
	}

	return result, nil
}
