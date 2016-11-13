package resources

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
    "github.com/supple/gorest/core"
)

const REPO_APP = "crm"

type App struct {
    CustomerBased `bson:",inline"`
    AppId    string        `json:"name" bson:"name" `
    GcmToken string        `json:"gcmToken" bson:"gcmToken"`
    Os       string        `json:"os" bson:"os"`
}

// ### -- App repository

type AppRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewAppRP(cc *core.CustomerContext) *AppRP {
    rp := &AppRP{cc:cc}
    db := s.GetInstance(REPO_APP)
    gt := core.NewGateway(rp.CollectionName(), cc, db)
    rp.gt = gt
    return rp
}

func (rp *AppRP) Create(model *App) error {
    model.CustomerName = rp.cc.CustomerName
    return rp.gt.Insert(model)
}

func (rp *AppRP) Update(id string, model *map[string]interface{}) error {
    err := rp.gt.Update(id, model)
    return err
}

func (rp *AppRP) FindOne(id string) (*App, error) {
    result := &App{}
    err := rp.gt.FindById(id, result)
    return result, err
}

func (rp *AppRP) FindOneBy(conditions bson.M) (*App, error) {
    result := &App{}
    err := rp.gt.FindOneBy(conditions, result)
    return result, err
}

func (rp *AppRP) Delete(id string) (error) {
    err := rp.gt.Remove(id)
    return err
}

func (rp AppRP) CollectionName() string {
    return "App"
}

func CreateApp(db *s.MongoDB, cc *core.CustomerContext, os string, gcmToken string) (*App, error) {
    aRp := NewAppRP(cc)

    app := &App{}
    app.GcmToken = gcmToken
    app.Os = os

    err := aRp.Create(app)
    return app, err
}