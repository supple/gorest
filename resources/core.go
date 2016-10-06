package resources

import (
	"errors"
	s "github.com/supple/gorest/storage"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	Create(db *s.Mongo, model interface{}) (error)
	Update(db *s.Mongo, id string, model interface{})
	FindOne(id string) (interface{}, error)
	CollectionName() string
}

type AccessTo struct {
    Resource string
    Action string
}

type CustomerContext struct {
    ApiKey string
    CustomerName string
}