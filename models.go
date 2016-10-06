package main

import (
    "sync"
    "math/rand"
    "reflect"
    "fmt"
    "bytes"
    "unicode"
    "unicode/utf8"
)

var (
    mu sync.RWMutex
)

func init() {
    fmt.Println("Models init")
}

// Helper random string generator
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
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

type AppService struct {
    Storage *MemStorage
}

// Device object
type Device struct {
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Version int32 `json:"version,omitempty"`
    Id string `json:"id,omitempty"`
}

func NewDevice(id string, name string) *Device {
    d := &Device{Name: name}
    d.Id = id
    return d
}

func (c *Device) Update(data map[string]interface{})  {
    updateModel(c, data)
}


