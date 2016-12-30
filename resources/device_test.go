package resources

import (
    "testing"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/storage"
    a "github.com/stretchr/testify/assert"
    "fmt"
    "github.com/supple/gorest/tests"
)

func init() {
    // Init storage instances
    tests.GetStorage()
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
    var cn = tests.TEST_CUSTOMER
    cc := &core.CustomerContext{CustomerName: tests.TEST_CUSTOMER}

    dRp := NewDeviceRP(cc)

    // create device on non existing customer
    d := &Device{}
    d.AppId = "xo"
    d.CustomerName = cn
    err = dRp.Create(d)
    a.Equal(t, core.ErrorFrom(core.ErrNotFound, "Customer not found"), err, "#1")

    // create customer
    tests.CreateTestCustomer()

    // create device on non existing app
    err = dRp.Create(d)
    a.Equal(t, core.ErrorFrom(core.ErrNotFound, "App id not set in api key"), err, "#3")

    // create app and device
    app, err := CreateAndroidApp(cc, "", "")
    cc.AppId = app.Id
    d.AppId = app.Id
    err = dRp.Create(d)
    if (err != nil) {
        fmt.Println(err.Error())
    }
    a.True(t, err == nil)
}
