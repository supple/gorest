package resources

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
)

type App struct {
    Id       string `json:"id" bson:"_id"`
    AppId    string        `json:"name" bson:"name" `
    GcmToken string        `json:"gcmToken" bson:"gcmToken"`
    AppVersion int32       `json:"appVersion" bson:"appVersion"`
    CustomerName string `json:"customerName" bson:"customerName" validate:"required"`
}

// ### -- App repository

type AppRP struct{
    gt *Gateway
}

func NewAppRP() *AppRP {
    rp := &AppRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *AppRP) Create(db *s.Mongo, model *App) error {
    return rp.gt.Insert(db, model)
}

func (rp *AppRP) Update(db *s.Mongo, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *AppRP) FindOne(db *s.Mongo, id string) (*App, error) {
    result := &App{}
    err := rp.gt.FindById(db, id, result)

    return result, err
}

func (rp *AppRP) FindOneBy(db *s.Mongo, conditions bson.M) (*App, error) {
    result := &App{}
    err := rp.gt.FindOneBy(db, conditions, result)
    return result, err
}

func (rp *AppRP) Delete(db *s.Mongo, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}

func (rp AppRP) CollectionName() string {
    return "App"
}
