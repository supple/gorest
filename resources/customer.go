package resources

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/core"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
)

const REPO_CUSTOMER = "crm"

type Customer struct {
    CustomerBased `bson:",inline"`
	Hash string `json:"hash" bson:"hash"`
}

// ### -- Customer repo

type CustomerRP struct{
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewCustomerRP(cc *core.CustomerContext) *CustomerRP {
    rp := &CustomerRP{cc:cc}
    db := s.GetInstance(REPO_CUSTOMER)
    gt := core.NewGateway(rp.CollectionName(), cc, db)
    rp.gt = gt

    return rp
}

func (rp *CustomerRP) Create(db *s.MongoDB, model *Customer) error {
    model.Hash = lc.RandString(8)
    return rp.gt.Insert(model)
}

func (rp *CustomerRP) Update(id string, model *map[string]interface{}) error {
    //err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    err := rp.gt.Update(id, model)
    return err
}

func (rp *CustomerRP) FindOne(id string) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindById(id, result)
	return result, err
}

func (rp *CustomerRP) FindOneByName(customerName string) (*Customer, error) {
	result := &Customer{}
    conditions := bson.M{CUSTOMER_NAME_FIELD: customerName}
	err := rp.gt.FindOneBy(conditions, result)
	return result, err
}

func (rp *CustomerRP) FindOneBy(conditions bson.M) (*Customer, error) {
	result := &Customer{}
    err := rp.gt.FindOneBy(conditions, result)
	return result, err
}

func (rp *CustomerRP) Delete(id string) (error) {
	err := rp.gt.Remove(id)
	return err
}

func (rp CustomerRP) CollectionName() string {
	return "Customer"
}

func CreateCustomer(db *s.MongoDB, name string) (*Customer, error) {
    cc := &core.CustomerContext{CustomerName: name}
    cRp := NewCustomerRP(cc)
    c := &Customer{}
    c.CustomerName = name
    err := cRp.Create(db, c)
    return c, err
}

func (rp *CustomerRP) Install(db *s.MongoDB) error {
	index := mgo.Index{
		Key: []string{CUSTOMER_NAME_FIELD},
		Unique: true,
		DropDups: false,
		Background: true, // See notes.
		Sparse: true,
	}
	err := db.Coll(rp.CollectionName()).EnsureIndex(index)
	return err
}
