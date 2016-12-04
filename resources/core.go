package resources

import (
	s "github.com/supple/gorest/storage"
    "unicode/utf8"
    "bytes"
    "unicode"
    "reflect"
    "fmt"
)

const CUSTOMER_NAME_FIELD string = "customerName"

type CustomerBased struct {
    Id           string `json:"id" bson:"_id"`
    CustomerName string `json:"customerName" bson:"customerName" validate:"required"`
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

// c model to be updated
func UpdateModel(c interface{}, data map[string]interface{}) {
    for k, v := range  data {
        // public field name in struct
        fieldName := ucfirst(k)

        //vDst := reflect.ValueOf(c).Elem().FieldByName(fieldName)
        vDst := reflect.Indirect(reflect.ValueOf(c)).FieldByName(fieldName)
        //fmt.Printf("%s %b \n", fieldName, vDst.CanSet())
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