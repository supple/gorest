package resources

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	s "github.com/supple/gorest/storage"
	lc "github.com/supple/gorest/utils"
)

type Customer struct {
    Id string `json:"id" bson:"_id"`
	Hash string `json:"hash" bson:"hash"`
	Name string `json:"name" bson:"name"`
}

// ### -- Customer repo

type CustomerRP struct{
    gt *Gateway
}

func NewCustomerRP() *CustomerRP {
    rp := &CustomerRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *CustomerRP) Create(db *s.Mongo, model *Customer) error {
    model.Hash = lc.RandString(8)
    return rp.gt.Insert(db, model)
}

func (rp *CustomerRP) Update(db *s.Mongo, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *CustomerRP) FindOne(db *s.Mongo, id string) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindById(db, id, result)

	return result, err
}

func (rp *CustomerRP) FindOneBy(db *s.Mongo, conditions bson.M) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindOneBy(db, conditions, result)
	return result, err
}

func (rp *CustomerRP) Delete(db *s.Mongo, id string) (error) {
	err := rp.gt.Remove(db, id)
	return err
}

func (rp CustomerRP) CollectionName() string {
	return "Customer"
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