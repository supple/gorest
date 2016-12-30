package resources

import (
	s "github.com/supple/gorest/storage"
    "unicode/utf8"
    "bytes"
    "unicode"
    "reflect"
    "fmt"
    "io"
    "encoding/json"
    "time"
)

const CUSTOMER_NAME_FIELD string = "customerName"

const (
    OS_ANDRIOD = "android"
    OS_IOS = "ios"
)

type CustomerBased struct {
    Id           string `json:"id,omitempty" bson:"_id"`
    CustomerName string `json:"customerName" bson:"customerName,omitempty" validate:"required"`
    CreatedAt    string  `json:"createdAt" bson:"createdAt,omitempty"`
    UpdatedAt    string  `json:"updatedAt" bson:"updatedAt,omitempty"`
    DeletedAt    string  `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

func (model *CustomerBased) SetBasicFields() {
    model.CreatedAt = GetJodaTime()
    model.UpdatedAt = model.CreatedAt
}

type Repository interface {
	Create(db *s.MongoDB, model interface{}) (error)
	Update(db *s.MongoDB, id string, model interface{})
	FindOne(id string) (interface{}, error)
	CollectionName() string
}

type AccessTo struct {
    Resource string
    Action string
}

func ucfirst(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    buf := &bytes.Buffer{}
    buf.WriteRune(unicode.ToUpper(r))
    buf.WriteString(s[size:])
    return buf.String()
}


func MapFromJson(data io.Reader) (*map[string]interface{}, error) {
    obj := make(map[string]interface{})
    decoder := json.NewDecoder(data)
    if err := decoder.Decode(obj); err != nil {
        return nil, err
    }

    return &obj, nil
}


// Update model properties from map
// Model has to be passed by ref
func UpdateModel(model interface{}, data map[string]interface{}) {
    for k, v := range  data {
        // public field name in struct
        fieldName := ucfirst(k)
        vDst := reflect.Indirect(reflect.ValueOf(model)).FieldByName(fieldName)
        if !vDst.CanSet() {
            continue
        }
        vSrc := reflect.ValueOf(v)
        if vDst.Type() != vSrc.Type() {
            if vSrc.Type().ConvertibleTo(vDst.Type()) {
                vDst.Set(vSrc.Convert(vDst.Type()))
                fmt.Printf("SET fieldName: %s, %d, dt: %s,\n", k, v, vSrc.Kind())
            }
        } else {
            vDst.Set(vSrc)
        }
    }
}

func GetJodaTime() string {
    return time.Now().Format("2006-01-02T15:04:05.999Z")
}