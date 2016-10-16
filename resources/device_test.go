package resources

import (
    "testing"
    "github.com/supple/gorest/core"
    s "github.com/supple/gorest/storage"
    a "github.com/stretchr/testify/assert"
    "fmt"
)

func TestDeviceRP_Update(t *testing.T) {
    d := Device{}

    m := make(map[string]interface{})
    m["appToken"] = "a"
    m["appVersion"] = "b"
    m["customerName"] = "c"
    m["appId"] = "d"

    UpdateModel(&d, m)
    fmt.Println(d.AppId)
    if (d.AppToken != "a") { t.Fatalf("Fail appToken: %s", d.AppToken) }
    if (d.AppVersion != "b") { t.Fatalf("Fail appVersion: %s", d.AppVersion) }
    if (d.CustomerName != "c") { t.Fatalf("Fail customerName: %s", d.CustomerName) }
    if (d.AppId != "d") { t.Fatalf("Fail appId: %s", d.AppId) }
}

func TestDeviceRP_Create(t *testing.T) {
    var err error
    var cn = "customer_test"
    cc := &core.CustomerContext{CustomerName:cn}

    db := s.GetInstance("entities")
    dRp := NewDeviceRP(cc)

    // prepare, drop device collection
    s.DropCollection(db, dRp.CollectionName())

    // create Device on non existing constrains
    d := &Device{}
    d.AppId = "xo"
    d.CustomerName = cn
    err = dRp.Create(db, d)
    a.True(t, err.Error() == (&ErrObjectNotFound{"Customer", d.CustomerName}).Error())

    // create customer and device with non existing app
    _, err = CreateCustomer(db, cn)
    err = dRp.Create(db, d)
    a.True(t, err.Error() == (&ErrObjectNotFound{"App", d.AppId}).Error())

    // create app and device
    app, err := CreateApp(db, cc, "android", "_")
    d.AppId = app.Id
    err = dRp.Create(db, d)
    a.True(t, err == nil)
}
