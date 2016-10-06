package resources

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	s "github.com/supple/gorest/storage"
	lc "github.com/supple/gorest/utils"
	"fmt"
	"errors"
)

type ApiKey struct {
	Id           string `json:"id" bson:"_id"`
	Key          string        `json:"key" bson:"key"`
	CustomerName string        `json:"customerName" bson:"customerName" validate:"required"`
}

func (rp *ApiKeyRP) ConstraintsValidation(db *s.Mongo, model *ApiKey) (*Customer, error) {
	csRp := CustomerRP{}
	c, err := csRp.FindOneBy(db, bson.M{"name": model.CustomerName})
	if (c == nil) {
		return nil, ErrNotFound
	}

	return c, err
}



// ### -- ApiKey repo

type ApiKeyRP struct{}

func (rp *ApiKeyRP) Create(db *s.Mongo, model *ApiKey) error {
	var err error

	// validate
	if (len(model.CustomerName) == 0) {
		return errors.New("Customer name not set")
	}
	customer, err := rp.ConstraintsValidation(db, model)
	if (err != nil) {
		return err
	}

	// create ids if not set
	if (len(model.Id) == 0) {
		model.Id = lc.NewId()
	}
	if (len(model.Key) == 0) {
		model.Key = fmt.Sprintf("%s-%s", lc.RandString(24), customer.Hash)
	}

	err = db.Coll(rp.CollectionName()).Insert(model)

	return err
}

func (rp *ApiKeyRP) FindOne(id string) (*Customer, error) {
	return nil, nil
}

func (rp *ApiKeyRP) FindOneBy(db *s.Mongo, conditions bson.M) (*ApiKey, error) {
	result := &ApiKey{}
	err := db.Coll(rp.CollectionName()).Find(conditions).One(result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (rp *ApiKeyRP) Install(db *s.Mongo) error {
	var index mgo.Index
	var err error

	// key in api key must be unique
	index = mgo.Index{
		Key: []string{"key"},
		Unique: true,
		DropDups: false,
		Background: true, // See notes.
		Sparse: true,
	}
	err = db.Coll(rp.CollectionName()).EnsureIndex(index)

	return err
}

func (rp ApiKeyRP) CollectionName() string {
	return "ApiKeyRP"
}

