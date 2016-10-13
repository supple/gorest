package resources

import (
    "gopkg.in/mgo.v2/bson"
    s "github.com/supple/gorest/storage"
)

type App struct {
    AppId    string        `json:"name" bson:"name" `
    GcmToken string        `json:"gcmToken" bson:"gcmToken"`
    AppVersion string       `json:"appVersion" bson:"appVersion"`
    Os string        `json:"os" bson:"os"`
    CustomerBased `bson:",inline"`
}

// ### -- App repository

type AppRP struct {
    gt *Gateway

}

func NewAppRP() *AppRP {
    rp := &AppRP{}
    gt := &Gateway{collectionName: rp.CollectionName()}
    rp.gt = gt

    return rp
}

func (rp *AppRP) Create(db *s.MongoDB, model *App) error {
    return rp.gt.Insert(db, model)
}

func (rp *AppRP) Update(db *s.MongoDB, id string, model *map[string]interface{}) error {
    err := db.Coll(rp.CollectionName()).Update(bson.M{"_id": id}, model)
    return err
}

func (rp *AppRP) FindOne(db *s.MongoDB, id string) (*App, error) {
    result := &App{}
    err := rp.gt.FindById(db, id, result)

    return result, err
}

func (rp *AppRP) FindOneBy(db *s.MongoDB, conditions bson.M) (*App, error) {
    result := &App{}
    err := rp.gt.FindOneBy(db, conditions, result)
    return result, err
}

func (rp *AppRP) Delete(db *s.MongoDB, id string) (error) {
    err := rp.gt.Remove(db, id)
    return err
}

func (rp AppRP) CollectionName() string {
    return "App"
}
