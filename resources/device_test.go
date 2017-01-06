package resources

import (
    "testing"
    "github.com/supple/gorest/core"
    "github.com/stretchr/testify/assert"
    "github.com/supple/gorest/tests"
    "github.com/supple/gorest/model"
)

func TestDeviceRP_Update(t *testing.T) {
    d := model.Device{}

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
    d := &model.Device{}
    d.AppId = "xo"
    d.CustomerName = cn
    err = dRp.Create(d)
    assert.Equal(t, core.ErrorFrom(core.ErrNotFound, "Customer not found"), err, "#1")

    // create customer
    CreateCustomer(tests.TEST_CUSTOMER)

    // create device on non existing app
    err = dRp.Create(d)
    assert.Equal(t, core.ErrorFrom(core.ErrNotFound, "App id not set in api key"), err, "#3")

    // create app and device
    app, err := CreateAndroidApp(cc, "", "")
    cc.AppId = app.Id
    d.AppId = app.Id
    err = dRp.Create(d)
    if (err != nil) {
        core.Log(err.Error())
    }
    assert.True(t, err == nil)
}
