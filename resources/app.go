package resources

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/supple/gorest/storage"
    "github.com/supple/gorest/core"
    "io"
    "encoding/json"
)

const REPO_APP = "crm"

type App struct {
    CustomerBased `bson:",inline"`
    Os       string `json:"os" bson:"os"`
    Name     string `json:"name" bson:"name"`
    GcmToken string `json:"gcmToken" bson:"gcmToken"`
    ApnsAuthKey string `json:"apnsAuthKey" bson:"apnsAuthKey"`
    ApnsTeamId string `json:"apnsTeamId" bson:"apnsTeamId"`
    ApnsKeyId string `json:"apnsKeyId" bson:"apnsKeyId"`
}


// ### -- App repository

type AppRP struct {
    gt *core.Gateway
    cc *core.CustomerContext
}

func NewAppRP(cc *core.CustomerContext) *AppRP {
    rp := &AppRP{cc: cc}
    db := storage.GetInstance(REPO_APP)
    d := core.ContextDecorator(cc)
    gt := core.NewGateway(rp.CollectionName(), d, db)
    rp.gt = gt
    return rp
}

func (rp *AppRP) Create(model *App) error {
    model.CustomerName = rp.cc.CustomerName
    err := rp.ConstraintsValidation(model)
    if (err != nil) {
        return err
    }
    model.SetBasicFields()

    return rp.gt.Insert(model)
}
//
//func (rp *AppRP) UpdateMap(id string, model *map[string]interface{}) error {
//    model["updatedAt"] = GetJodaTime()
//    err := rp.gt.Update(id, model)
//    return err
//}

func (rp *AppRP) Update(id string, model *App) error {
    model.UpdatedAt = GetJodaTime()
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

func (rp *AppRP) ConstraintsValidation(model *App) (error) {
    var err error
    csRp := NewCustomerRP(rp.cc)
    _, err = csRp.FindOneByName(model.CustomerName)
    if err == core.ErrNotFound {
        return core.ErrorFrom(core.ErrNotFound,  "Customer not found")
    }
    if (err != nil) {
        return err
    }

    return err
}

func (rp AppRP) CollectionName() string {
    return "App"
}


func AppFromJson(data io.Reader) (*App, error) {
    obj := &App{}
    decoder := json.NewDecoder(data)
    if err := decoder.Decode(obj); err != nil {
        return nil, err
    }

    // validate

    return obj, nil
}

type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

func (er ValidationError) Error() string {
    b, err := json.Marshal(er)
    if err != nil {
        return ""
    } else {
        return string(b)
    }
}

func ValidateApp(m *App) []*ValidationError {
    var errors []*ValidationError

    // os required - one of
    switch m.Os {
    case OS_ANDRIOD, OS_IOS: // ok
    default:
        errors = append(errors, &ValidationError{Field: "os", Message: "Invalid os value, expected: `ios` or `android`"})
    }

    // name required
    if len(m.Name) == 0 {
        errors = append(errors, &ValidationError{Field: "name", Message: "Name cannot be empty"})
    }

    // on update

    return errors
}
//
//func ValidateUpdateApp(m *map[string]interface{}) []*ValidationError {
//    var errors []*ValidationError
//
//    // os required - one of
//    switch m.Os {
//    case OS_ANDRIOD, OS_IOS: // ok
//    default:
//        errors = append(errors, &ValidationError{Field: "os", Message: "Invalid os value, expected: `ios` or `android`"})
//    }
//
//    // name required
//    if len(m.Name) == 0 {
//        errors = append(errors, &ValidationError{Field: "name", Message: "Name cannot be empty"})
//    }
//
//    // on update
//
//    return errors
//}

func CreateAndroidApp(cc *core.CustomerContext, gcmToken string, id string) (*App, error) {
    aRp := NewAppRP(cc)

    app := &App{}
    app.Id = id
    app.GcmToken = gcmToken
    app.Os = OS_ANDRIOD

    err := aRp.Create(app)
    return app, err
}