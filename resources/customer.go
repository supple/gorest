package resources

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	s "github.com/supple/gorest/storage"
	lc "github.com/supple/gorest/utils"
)

type Customer struct {
    Base `bson:",inline"`
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

func (rp *CustomerRP) Create(db *s.MongoDB, model *Customer) error {
    model.Hash = lc.RandString(8)
    return rp.gt.Insert(db, model)
}

func (rp *CustomerRP) Update(db *s.MongoDB, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *CustomerRP) FindOne(db *s.MongoDB, id string) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindById(db, id, result)

	return result, err
}

func (rp *CustomerRP) FindOneByName(db *s.MongoDB, customerName string) (*Customer, error) {
	result := &Customer{}
    conditions := bson.M{"name": customerName}
	err := rp.gt.FindOneBy(db, conditions, result)
	return result, err
}

func (rp *CustomerRP) FindOneBy(db *s.MongoDB, conditions bson.M) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindOneBy(db, conditions, result)
	return result, err
}

func (rp *CustomerRP) Delete(db *s.MongoDB, id string) (error) {
	err := rp.gt.Remove(db, id)
	return err
}

func (rp CustomerRP) CollectionName() string {
	return "Customer"
}

func CreateCustomer(db *s.MongoDB, name string) (*Customer, error) {
    cRp := NewCustomerRP()
    c := &Customer{}
    c.Name = name
    err := cRp.Create(db, c)

    return c, err
}

func (rp *CustomerRP) Install(db *s.MongoDB) error {
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
