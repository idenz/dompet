package finances

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
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

	result := make([]Model, 0)

	cursor, err := s.collection.Find(s.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(s.ctx) {

		finance := Model{}
		err := cursor.Decode(&finance)
		if err != nil {
			return nil, err
		}

		result = append(result, finance)

	}

	return &result, nil
}
