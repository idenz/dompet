package finances

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	Create(payload *CreateRequest) (*Model, error)
	GetAll() (*[]Model, error)
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

func (s *Service) Create(payload *CreateRequest) (*Model, error) {

	finance := new(Model)
	bytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, err

	}

	err = json.Unmarshal(bytes, &finance)
	if err != nil {
		return nil, err
	}

	_, err = s.collection.InsertOne(s.ctx, finance)
	if err != nil {
		return nil, err
	}

	return finance, nil

}

func (s *Service) GetAll() (*[]Model, error) {
	return nil, nil
}
