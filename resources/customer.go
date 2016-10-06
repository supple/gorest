package resources

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	s "github.com/supple/gorest/storage"
	lc "github.com/supple/gorest/utils"
)

type Customer struct {
	Id   string `json:"id" bson:"_id"`
	Hash string `json:"hash" bson:"hash"`
	Name string `json:"name" bson:"name"`
}

// ### -- Customer repo

type CustomerRP struct{}

func (rp *CustomerRP) Create(db *s.Mongo, model *Customer) error {
	if (len(model.Id) == 0) {
		model.Id = lc.NewId()
	}
	model.Hash = lc.RandString(8)
	err := db.Coll(rp.CollectionName()).Insert(model)
	return err
}

func (rp *CustomerRP) FindOne(db *s.Mongo, id string) (*Customer, error) {
	result := &Customer{}
	q := bson.M{"_id": id}
	err := db.Coll(rp.CollectionName()).Find(q).One(result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (rp *CustomerRP) FindOneBy(db *s.Mongo, conditions bson.M) (*Customer, error) {
	result := &Customer{}
	err := db.Coll(rp.CollectionName()).Find(conditions).One(result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (rp *CustomerRP) Delete(db *s.Mongo, id string) (error) {
	q := bson.M{"_id": id}
	err := db.Coll(rp.CollectionName()).Remove(q)
	return err
}

func (rp CustomerRP) CollectionName() string {
	return "CustomerRP"
}

func (rp *CustomerRP) Install(db *s.Mongo) error {
	index := mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: false,
		Background: true, // See notes.
		Sparse: true,
	}
	err := db.Coll(rp.CollectionName()).EnsureIndex(index)
	return err
}