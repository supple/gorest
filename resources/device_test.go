package resources

import (
    "testing"
    s "github.com/supple/gorest/storage"
    a "github.com/stretchr/testify/assert"
    "fmt"
)

func CreateCustomer(db *s.MongoDB, name string) (*Customer, error) {
    cRp := NewCustomerRP()
    c := &Customer{}
    c.Name = "marek"
    err := cRp.Create(db, c)

    return c, err
}

func CreateApiKey(db *s.MongoDB, c *Customer, apiKey string) (*ApiKey, error) {
    akRp := NewApiKeyRP()
    ak := &ApiKey{}
    ak.CustomerName = c.Name
    err := akRp.Create(db, ak)

    return ak, err
}

func CreateApp(db *s.MongoDB, c *Customer, os string) (*App, error) {
    aRp := NewAppRP()
    app := &App{}
    app.CustomerName = c.Name
    app.AppVersion = "1.2.0"
    app.Os = os
    err := aRp.Create(db, app)

    return app, err
}

func TestDeviceRP_Update(t *testing.T) {
    d := Device{}

    m := make(map[string]interface{})
    m["appToken"] = "xo"
    m["appVersion"] = "1.2"
    m["customerName"] = "xod"
    m["appId"] = "xos"

    UpdateModel(&d, m)
    fmt.Println(d.AppId)
    if (d.AppId != "xos") { t.Fatalf("Fail appId: %s", d.AppId) }
    if (d.CustomerName != "xod") { t.Fatalf("Fail customerName: %s", d.CustomerName) }
    if (d.AppToken != "xo") { t.Fatalf("Fail appToken: %s", d.AppToken) }
}

func TestDeviceRP_Create(t *testing.T) {
    var err error

    db := s.GetInstance("entities")
    dRp := NewDeviceRP()

    // prepare
    s.DropCollection(db, dRp.CollectionName())

    cn := "marek"

    d := &Device{}
    d.AppId = "xo"
    d.CustomerName = cn
    err = dRp.Create(db, d)

    a.True(t, err.Error() == (&ErrObjectNotFound{"Customer", d.CustomerName}).Error())

    // create customer
    c, err := CreateCustomer(db, cn)
    err = dRp.Create(db, d)
    fmt.Println(err.Error())
    a.True(t, err.Error() == (&ErrObjectNotFound{"App", d.AppId}).Error())

    // create app
    app, err := CreateApp(db, c, "android")
    fmt.Println(app.Id)
    d.AppId = app.Id
    err = dRp.Create(db, d)
    //fmt.Println(err.Error())
    a.True(t, err == nil)
}


