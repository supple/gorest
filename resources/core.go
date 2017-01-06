package resources

import (
	"github.com/supple/gorest/storage"
    "unicode/utf8"
    "bytes"
    "unicode"
    "reflect"
    "fmt"
    "io"
    "encoding/json"
)

const CUSTOMER_NAME_FIELD string = "customerName"

const (
    OS_ANDRIOD = "android"
    OS_IOS = "ios"
)

type Repository interface {
	Create(db *storage.MongoDB, model interface{}) (error)
	Update(db *storage.MongoDB, id string, model interface{})
	FindOne(id string) (interface{}, error)
	CollectionName() string
}

// ACL object
type AccessTo struct {
    Resource string
    Action string
}

// Make first letter capital
func ucfirst(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    buf := &bytes.Buffer{}
    buf.WriteRune(unicode.ToUpper(r))
    buf.WriteString(s[size:])
    return buf.String()
}

// Create map from json string
func MapFromJson(data io.Reader) (*map[string]interface{}, error) {
    obj := make(map[string]interface{})
    decoder := json.NewDecoder(data)
    if err := decoder.Decode(obj); err != nil {
        return nil, err
    }

    return &obj, nil
}

// Update model properties from map
// Struct has to be passed by ref
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
