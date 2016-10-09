package resources

import (
	"errors"
	s "github.com/supple/gorest/storage"
    "unicode/utf8"
    "bytes"
    "unicode"
    "reflect"
    "fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

//type Base struct {
//    Id string `json:"id" bson:"_id"`
//}

type Repository interface {
	Create(db *s.Mongo, model interface{}) (error)
	Update(db *s.Mongo, id string, model interface{})
	FindOne(id string) (interface{}, error)
	CollectionName() string
}



type AccessTo struct {
    Resource string
    Action string
}

type CustomerContext struct {
    ApiKey string
    CustomerName string
}

func ucfirst(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    buf := &bytes.Buffer{}
    buf.WriteRune(unicode.ToUpper(r))
    buf.WriteString(s[size:])
    return buf.String()
}

// c model to be updated
func updateModel(c interface{}, data map[string]interface{}) {
    for k, v := range  data {
        // public field name in struct
        fieldName := ucfirst(k)
        vDst := reflect.ValueOf(c).Elem().FieldByName(fieldName)
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