package resources

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/core"
    s "github.com/supple/gorest/storage"
    lc "github.com/supple/gorest/utils"
    "fmt"
)

const API_KEY_FIELD string = "apiKey"

type ApiKey struct {
    CustomerBased `bson:",inline"`
    AppId  string `json:"appId" bson:"appId"`
    ApiKey string `json:"apiKey" bson:"apiKey"`
}

// ### -- ApiKey repo

type ApiKeyRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewApiKeyRP(cc *core.CustomerContext) *ApiKeyRP {
    rp := &ApiKeyRP{cc:cc}
    gt := core.NewGateway(rp.CollectionName(), cc)
    rp.gt = gt

    return rp
}

func (rp *ApiKeyRP) Create(db *s.MongoDB, model *ApiKey) error {
    var err error
    model.CustomerName = rp.cc.CustomerName

     // validate
    customer, err := rp.ConstraintsValidation(db, model)
    if (err != nil) {
        return &ErrObjectNotFound{"Customer", rp.cc.CustomerName}
    }
    fmt.Println("OK CUST is")
    // create key if not set
    if (len(model.ApiKey) == 0) {
        model.ApiKey = fmt.Sprintf("%s-%s", lc.RandString(24), customer.Hash)
    }
    err = rp.gt.Insert(db, model)
    fmt.Println(model)
    return err
}

func (rp *ApiKeyRP) FindOne(db *s.MongoDB, id string) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindById(db, id, result)
    return result, err
}

func (rp *ApiKeyRP) FindOneBy(db *s.MongoDB, conditions bson.M) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindInsecureOneBy(db, conditions, result)
    return result, err
}

func (rp *ApiKeyRP) Delete(db *s.MongoDB, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}

func (rp *ApiKeyRP) ConstraintsValidation(db *s.MongoDB, model *ApiKey) (*Customer, error) {
    csRp := NewCustomerRP(rp.cc)
    c, err := csRp.FindOneByName(db, model.CustomerName)
    if (c == nil) {
        return nil, ErrNotFound
    }

    return c, err
}

func (rp ApiKeyRP) CollectionName() string {
    return "ApiKey"
}

func CreateApiKey(db *s.MongoDB, cc *core.CustomerContext) (*ApiKey, error) {
    akRp := NewApiKeyRP(cc)
    ak := &ApiKey{}
    ak.CustomerName = cc.CustomerName
    err := akRp.Create(db, ak)

    return ak, err
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

