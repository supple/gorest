package resources

import (
	"encoding/json"
	"fmt"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/model"
	"github.com/supple/gorest/storage"
	"gopkg.in/mgo.v2/bson"
	"io"
)

const REPO_APP = "crm"

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

func (rp *AppRP) Create(model *model.App) error {
	model.CustomerName = rp.cc.CustomerName
	err := rp.ConstraintsValidation(model)
	if err != nil {
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

func (rp *AppRP) Update(id string, model *model.App) error {
	model.UpdatedAt = core.GetJodaTime()
	err := rp.gt.Update(id, model)
	return err
}

func (rp *AppRP) FindOne(id string) (*model.App, error) {
	result := &model.App{}
	err := rp.gt.FindById(id, result)
	return result, err
}

func (rp *AppRP) FindOneBy(conditions bson.M) (*model.App, error) {
	result := &model.App{}
	err := rp.gt.FindOneBy(conditions, result)
	return result, err
}

func (rp *AppRP) Delete(id string) error {
	err := rp.gt.Remove(id)
	return err
}

func (rp *AppRP) ConstraintsValidation(model *model.App) error {
	var err error
	csRp := NewCustomerRP(rp.cc)
	_, err = csRp.FindOneByName(model.CustomerName)
	if err == core.ErrNotFound {
		return core.ErrorFrom(core.ErrNotFound, "Customer not found")
	}
	if err != nil {
		return err
	}

	return err
}

func (rp AppRP) CollectionName() string {
	return "App"
}

func AppFromJson(data io.Reader) (*model.App, error) {
	var e []core.ValidationError
	var decoded interface{}

	decoder := json.NewDecoder(data)

	// check if string is valid json object
	if err := decoder.Decode(&decoded); err != nil {
		e = append(e, core.NewError("", "Invalid JSON object"))
		return nil, core.ValidationErrors{Errors: e}
	}

	// validate object fields, check if types are ok
	e = ValidateAppStruct(decoded)
	if len(e) > 0 {
		return nil, core.ValidationErrors{Errors: e}
	}

	obj := &model.App{}
	e = UpdateModel(obj, decoded.(map[string]interface{}))
	if len(e) > 0 {
		return nil, core.ValidationErrors{Errors: e}
	}

	fmt.Printf("%#v\n", obj)

	return obj, nil
}

func ValidateAppStruct(decoded interface{}) []core.ValidationError {
	var errors []core.ValidationError

	// check if json is array
	switch decoded.(type) {
	case []interface{}:
		return append(errors, core.NewError("", "Invalid data. JSON is array"))
	}

	m, ok := decoded.(map[string]interface{})
	if !ok {
		return append(errors, core.NewError("", "Invalid JSON object"))
	}

	// todo: remove
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv, fmt.Sprintf("%T", v))
		case int:
			fmt.Println(k, "is int", vv, fmt.Sprintf("%T", v))
		case []interface{}:
			fmt.Println(k, "is an array:", fmt.Sprintf("%T", v))
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv, fmt.Sprintf("%T", v))
		}
	}

	return errors
}

func ValidateApp(m *model.App) []*core.ValidationError {
	var errors []*core.ValidationError

	// os required - one of
	switch m.Os {
	case OS_ANDRIOD, OS_IOS: // ok
	default:
		errors = append(errors, &core.ValidationError{Field: "os", Message: "Invalid os value, expected: `ios` or `android`"})
	}

	// name required
	if len(m.Name) == 0 {
		errors = append(errors, &core.ValidationError{Field: "name", Message: "Name cannot be empty"})
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

func CreateAndroidApp(cc *core.CustomerContext, gcmToken string, id string) (*model.App, error) {
	aRp := NewAppRP(cc)

	app := &model.App{}
	app.Id = id
	app.GcmToken = gcmToken
	app.Os = OS_ANDRIOD

	err := aRp.Create(app)
	return app, err
}
