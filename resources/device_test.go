package resources

import (
    "testing"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/storage"
    a "github.com/stretchr/testify/assert"
    "fmt"
)

func init() {
    storage.SetInstance("crm", storage.NewMongoDB("192.168.1.106:27017", "crm_test"))
    db := storage.GetInstance("crm")
    storage.DropDatabase(db)
}

func TestDeviceRP_Update(t *testing.T) {
    d := Device{}

    m := make(map[string]interface{})
    m["appToken"] = "a"
    m["appVersion"] = "b"
    m["customerName"] = "c"
    m["appId"] = "d"

    UpdateModel(&d, m)

    if (d.AppToken != "a") { t.Fatalf("Fail appToken: %s", d.AppToken) }
    if (d.AppVersion != "b") { t.Fatalf("Fail appVersion: %s", d.AppVersion) }
    if (d.CustomerName != "c") { t.Fatalf("Fail customerName: %s", d.CustomerName) }
    if (d.AppId != "d") { t.Fatalf("Fail appId: %s", d.AppId) }
}

func TestDeviceRP_Create(t *testing.T) {
    var err error
    var cn = "customer_test"
    cc := &core.CustomerContext{CustomerName:cn}

    db := storage.GetInstance("crm")
    a.True(t, db != nil, "Database instance not found")
    if (db == nil) {
        return
    }
    dRp := NewDeviceRP(cc)

    // prepare, drop device collection
    storage.DropCollection(db, dRp.CollectionName())

    // create device on non existing customer
    d := &Device{}
    d.AppId = "xo"
    d.CustomerName = cn
    err = dRp.Create(d)
    a.Equal(t, err, (&core.ErrObjectNotFound{"Customer", ""}), "#1")

    // create customer
    customer, err := CreateCustomer(cn)
    fmt.Println(customer)
    a.Equal(t, nil, err)
    a.Equal(t, customer.CustomerName, "customer_test", "#2")

    // create device on non existing app
    err = dRp.Create(d)
    a.Equal(t, err, (&core.ErrObjectNotFound{"App", ""}), "#3")

    // create app and device
    app, err := CreateApp(cc, "android", "_")
    d.AppId = app.Id
    err = dRp.Create(d)
    a.True(t, err == nil)
}
