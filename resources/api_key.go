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
    Key          string `json:"key" bson:"key"`
    CustomerName string `json:"customerName" bson:"customerName" validate:"required"`
}

// ### -- ApiKey repo

type ApiKeyRP struct {
    gt *Gateway
}

func NewApiKeyRP() *ApiKeyRP {
    rp := &ApiKeyRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *ApiKeyRP) Create(db *s.MongoDB, model *ApiKey) error {
    var err error

     // validate
    if (len(model.CustomerName) == 0) {
        return errors.New("Customer name not set")
    }
    customer, err := rp.ConstraintsValidation(db, model)
    if (err != nil) {
        return err
    }

    // create key if not set
    if (len(model.Key) == 0) {
        model.Key = fmt.Sprintf("%s-%s", lc.RandString(24), customer.Hash)
    }

    err = rp.gt.Insert(db, model)

    return err
}

func (rp *ApiKeyRP) FindOne(db *s.MongoDB, id string) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindById(db, id, result)
    return result, err
}

func (rp *ApiKeyRP) FindOneBy(db *s.MongoDB, conditions bson.M) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindOneBy(db, conditions, result)
    return result, err
}

func (rp *ApiKeyRP) Delete(db *s.MongoDB, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}

func (rp *ApiKeyRP) ConstraintsValidation(db *s.MongoDB, model *ApiKey) (*Customer, error) {
    csRp := NewCustomerRP()
    c, err := csRp.FindOneByName(db, model.CustomerName)
    if (c == nil) {
        return nil, ErrNotFound
    }

    return c, err
}

func (rp ApiKeyRP) CollectionName() string {
    return "ApiKey"
}


func (rp *ApiKeyRP) Install(db *s.MongoDB) error {
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

