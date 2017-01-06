package resources

import (
    "testing"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/model"
    "github.com/supple/gorest/storage"
    a "github.com/stretchr/testify/assert"
    "fmt"
)

func TestAppRP_Update(t *testing.T) {
    d := model.App{}

    m := make(map[string]interface{})
    m["os"] = OS_ANDRIOD
    m["gcmToken"] = "test"
    m["name"] = "Android app"
    m["customerName"] = "c"

    UpdateModel(&d, m)

    if (d.Name != "Android app") { t.Fatalf("Fail name: %s", d.Name) }
    if (d.CustomerName != "c") { t.Fatalf("Fail customerName: %s", d.CustomerName) }
    if (d.Os != OS_ANDRIOD) { t.Fatalf("Fail os: %s", d.Os) }
    if (d.GcmToken != "test") { t.Fatalf("Fail gcmToken: %s", d.GcmToken) }
}

func TestAppRP_Create(t *testing.T) {
    var err error
    var cn = "customer_test"
    cc := &core.CustomerContext{CustomerName:cn}

    db := storage.GetInstance("crm")
    storage.DropDatabase(db)

    appRp := NewAppRP(cc)

    // prepare, drop device collection
    storage.DropCollection(db, appRp.CollectionName())

    // create device on non existing customer
    app := &model.App{}
    app.Name = "xo"
    app.Os = OS_ANDRIOD
    app.CustomerName = cn
    err = appRp.Create(app)
    a.Equal(t, core.ErrorFrom(core.ErrNotFound, "Customer not found"), err, "#1")

    // create customer
    customer, err := CreateCustomer(cn)
    fmt.Println(customer)
    a.Equal(t, nil, err)
    a.Equal(t, customer.CustomerName, cn, "#2")

    // create app and device
    appNew, err := CreateAndroidApp(cc, "", "")
    a.True(t, err == nil)
    a.Equal(t, appNew.Os, OS_ANDRIOD)
}
