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
            t := vDst.Kind()
            switch t {
            case reflect.Int32:
                switch knd := vSrc.Kind(); knd {
                case reflect.Int:
                    if vp, okp  := v.(int); okp {
                        vDst.SetInt(int64(vp))
                    }
                case reflect.Float64:
                    if vp, okp := v.(float64); okp {
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
    ca *CacheArr
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

// Campaign object
type Device struct {
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Version int32 `json:"version,omitempty"`
    Object
}

func NewCampaign(id string, name string) *Device {
    d := &Device{Name: name}
    d.Id = id
    return d
}

func (c *Device) Update(data map[string]interface{})  {
    updateModel(c, data)
}

// Cache, campaign id -> *Device
type CacheArr struct {
    campaigns map[string]*Device
}

func NewCacheArr(names []string) *CacheArr {
    ca := CacheArr{}
    ca.campaigns = make(map[string]*Device, len(names))
    for _, name := range names {
        id := randSeq(5)
        ca.campaigns[id] = NewCampaign(id, name)
    }

    return &ca
}

// Get returns campaign, or nil if there's no one.
func (ca* CacheArr) Get(id string) *Device {
    mu.RLock()
    cmp := ca.campaigns[id]
    mu.RUnlock()

    return cmp
}

// Set campaign in cache
func (ca* CacheArr) Set(pCmp *Device) {
    mu.Lock()
    ca.campaigns[pCmp.Id] = pCmp
    mu.Unlock()
}

// Update campaign in cache
func (ca* CacheArr) Update(pCmp *Device) {
    mu.Lock()
    ca.campaigns[pCmp.Id] = pCmp
    mu.Unlock()
}

// Get campaign names with ids
func (ca* CacheArr) GetNames() map[string]string {
    names := make(map[string]string)
    mu.RLock()
    for _, cmp := range ca.campaigns {
        names[cmp.Id] = cmp.Name
    }
    mu.RUnlock()
    return names
}