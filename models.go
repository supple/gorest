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

// Helper random string generator
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func updateModel(c interface{}, data map[string]interface{}) {

    for k, v := range  data {
        // public field name in struct
        fieldName := ucfirst(k)
        vDst := reflect.ValueOf(c).Elem().FieldByName(fieldName)
        vSrc := reflect.ValueOf(v)
        if !vDst.CanSet() {
            continue
        }
        if vDst.Type() != vSrc.Type() {
            switch tDst := vDst.Kind(); tDst {
            case reflect.Int32:
                switch tSrc := vSrc.Kind(); tSrc {
                case reflect.Int:
                    if vp, ok := v.(int); ok {
                        vDst.SetInt(int64(vp))
                    }
                case reflect.Float64:
                    if vp, ok := v.(float64); ok {
                        vDst.SetInt(int64(vp))
                    }
                }
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

func ucfirst(s string) string {
    r, size := utf8.DecodeRuneInString(s)
    buf := &bytes.Buffer{}
    buf.WriteRune(unicode.ToUpper(r))
    buf.WriteString(s[size:])
    return buf.String()
}

type Object struct {
    Id string `json:"id,omitempty"`
}

// Device object
type Device struct {
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Version int32 `json:"version,omitempty"`
    Object
}

func NewDevice(id string, name string) *Device {
    d := &Device{Name: name}
    d.Id = id
    return d
}

func (c *Device) Update(data map[string]interface{})  {
    updateModel(c, data)
}

func (c *Device) ToString() string {
    return c.Name
}

// Cache, id -> *Object
type MemStorage struct {
    objects map[string]*Device
}

func NewMemStorage(names []string) *MemStorage {
    ca := MemStorage{}
    ca.objects = make(map[string]*Device, len(names))
    for _, name := range names {
        id := randSeq(5)
        ca.objects[id] = NewDevice(id, name)
    }

    return &ca
}

// Get returns object, or nil if there's no one.
func (ca* MemStorage) Get(id string) *Device {
    mu.RLock()
    cmp := ca.objects[id]
    mu.RUnlock()

    return cmp
}

// Set object in storage
func (ca* MemStorage) Set(pCmp *Device) {
    mu.Lock()
    ca.objects[pCmp.Id] = pCmp
    mu.Unlock()
}

// Update object in storage
func (ca* MemStorage) Update(obj *Device) {
    mu.Lock()
    ca.objects[obj.Id] = obj
    mu.Unlock()
}

// Get object names with ids
func (ca* MemStorage) GetNames() map[string]string {
    names := make(map[string]string)
    mu.RLock()
    for _, obj := range ca.objects {
        names[obj.Id] = obj.ToString()
    }
    mu.RUnlock()
    return names
}