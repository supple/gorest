package resources

import (
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/model"
	"github.com/supple/gorest/storage"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const REPO_CUSTOMER = "crm"

// ### -- Customer repo

type CustomerRP struct {
	gt *core.Gateway
	cc *core.CustomerContext
}

func NewCustomerRP(cc *core.CustomerContext) *CustomerRP {
	rp := &CustomerRP{cc: cc}
	db := storage.GetInstance(REPO_CUSTOMER)
	d := core.ContextDecorator(cc)
	gt := core.NewGateway(rp.CollectionName(), d, db)
	rp.gt = gt

	return rp
}

func (rp *CustomerRP) Create(model *model.Customer) error {
	_, err := rp.FindOneByName(model.CustomerName)
	if err != nil {
		if err != core.ErrNotFound {
			return err
		}
		//_, ok := err.(*core.ErrObjectNotFound)
		//if (!ok) {
		//    return err
		//}
		model.SetBasicFields()
		model.Hash = core.RandString(8)
		return rp.gt.Insert(model)
	}

	return err
}

func (rp *CustomerRP) Update(id string, model *map[string]interface{}) error {
	//err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
	err := rp.gt.Update(id, model)
	return err
}

func (rp *CustomerRP) FindOne(id string) (*model.Customer, error) {
	result := &model.Customer{}
	err := rp.gt.FindById(id, result)
	return result, err
}

func (rp *CustomerRP) FindOneByName(customerName string) (*model.Customer, error) {
	result := &model.Customer{}
	conditions := bson.M{CUSTOMER_NAME_FIELD: customerName}
	err := rp.gt.FindOneBy(conditions, result)
	return result, err
}

func (rp *CustomerRP) FindOneBy(conditions bson.M) (*model.Customer, error) {
	result := &model.Customer{}
	err := rp.gt.FindOneBy(conditions, result)
	return result, err
}

func (rp *CustomerRP) Delete(id string) error {
	err := rp.gt.Remove(id)
	return err
}

func (rp CustomerRP) CollectionName() string {
	return "Customer"
}

func CreateCustomer(name string) (*model.Customer, error) {
	cc := &core.CustomerContext{CustomerName: name}
	cRp := NewCustomerRP(cc)
	c := &model.Customer{}
	c.CustomerName = name
	err := cRp.Create(c)
	if err != nil {
		return nil, err
	}
	return c, err
}

func (rp *CustomerRP) Install(db *storage.MongoDB) error {
	index := mgo.Index{
		Key:        []string{CUSTOMER_NAME_FIELD},
		Unique:     true,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := db.Coll(rp.CollectionName()).EnsureIndex(index)
	return err
}
