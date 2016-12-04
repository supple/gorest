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
const API_KEY_REPO = "crm"

type ApiKey struct {
    CustomerBased `bson:",inline"`
    AppId  string `json:"appId" bson:"appId"`
    ApiKey string `json:"apiKey" bson:"apiKey"`
}

// ### -- ApiKey repo

type ApiKeyRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
    db *s.MongoDB
}

func NewApiKeyRP(cc *core.CustomerContext) *ApiKeyRP {
    rp := &ApiKeyRP{cc:cc}
    db := s.GetInstance(API_KEY_REPO)
    gt := core.NewGateway(rp.CollectionName(), cc, db)
    rp.gt = gt

    return rp
}

func (rp *ApiKeyRP) Create(model *ApiKey) error {
    var err error
    model.CustomerName = rp.cc.CustomerName

     // validate
    customer, err := rp.ConstraintsValidation(model)
    if (err != nil) {
        return err
    }

    fmt.Println("OK CUST is")
    // create key if not set
    if (len(model.ApiKey) == 0) {
        model.ApiKey = fmt.Sprintf("%s-%s", lc.RandString(24), customer.Hash)
    }
    err = rp.gt.Insert(model)
    fmt.Println(model)
    return err
}

func (rp *ApiKeyRP) FindOne(id string) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindById(id, result)
    return result, err
}

func (rp *ApiKeyRP) FindOneBy(conditions bson.M) (*ApiKey, error) {
    result := &ApiKey{}
    err := rp.gt.FindOneWithoutContextBy(conditions, result)

    return result, err
}

func (rp *ApiKeyRP) Delete(id string) (error) {
    err := rp.gt.Remove(id)
    return err
}

func (rp *ApiKeyRP) ConstraintsValidation(model *ApiKey) (*Customer, error) {
    csRp := NewCustomerRP(rp.cc)
    c, err := csRp.FindOneByName(model.CustomerName)
    if (err != nil) {
        return nil, err
    }

    return c, err
}

func (rp ApiKeyRP) CollectionName() string {
    return "ApiKey"
}

func Auth(apiKey string, accessTo AccessTo) (*core.CustomerContext, error) {
    var cc *core.CustomerContext

    akRp := NewApiKeyRP(cc)
    ak, err := akRp.FindOneBy(bson.M{API_KEY_FIELD: apiKey})
    // @todo: hasAccess(accessTo)
    if err == mgo.ErrNotFound {
        return nil, core.ErrInvalidApiKey
    }

    if (ak != nil) {
        cc = &core.CustomerContext{}
        // copy to customer context
        cc.ApiKey = ak.ApiKey
        cc.CustomerName = ak.CustomerName
        cc.AppId = ak.AppId
    }

    return cc, err
}

func CreateApiKey(cc *core.CustomerContext) (*ApiKey, error) {
    akRp := NewApiKeyRP(cc)
    ak := &ApiKey{}
    ak.CustomerName = cc.CustomerName
    err := akRp.Create(ak)

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

