package resources

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/model"
	"github.com/supple/gorest/tests"
	"strings"
	"testing"
)

func TestAppRP_Update(t *testing.T) {
	d := model.App{}

	m := make(map[string]interface{})
	m["os"] = OS_ANDRIOD
	m["gcmToken"] = "test"
	m["name"] = "Android app"
	m["customerName"] = "c"

	UpdateModel(&d, m)

	if d.Name != "Android app" {
		t.Fatalf("Fail name: %s", d.Name)
	}
	if d.CustomerName != "c" {
		t.Fatalf("Fail customerName: %s", d.CustomerName)
	}
	if d.Os != OS_ANDRIOD {
		t.Fatalf("Fail os: %s", d.Os)
	}
	if d.GcmToken != "test" {
		t.Fatalf("Fail gcmToken: %s", d.GcmToken)
	}
}

func TestAppFromJson(t *testing.T) {
	var data *strings.Reader
	var obj *model.App
	var err error

	// err json as array
	data = strings.NewReader(`[{"name": 21, "os": "android", "createdAt": "test"}]`)
	_, err = AppFromJson(data)
	assert.Error(t, err)
	assert.Equal(t, "{\"errors\":[{\"field\":\"\",\"message\":\"Invalid data. JSON is array\"}]}", err.Error())

	// err invalid json
	data = strings.NewReader(`{"name": 21, "os": "android", "createdAt": "test", #wrong_json`)
	_, err = AppFromJson(data)
	assert.Error(t, err)
	assert.Equal(t, "{\"errors\":[{\"field\":\"\",\"message\":\"Invalid JSON object\"}]}", err.Error())

	// err invalid type
	data = strings.NewReader(`{"name": 21}`)
	obj, err = AppFromJson(data)
	assert.Error(t, err)
	assert.Equal(t, "{\"errors\":[{\"field\":\"name\",\"message\":\"Invalid type. Must be a string\"}]}", err.Error())
	core.Log(err.Error())

	// err forbidden field
	// err not existing field

	// ok
	data = strings.NewReader(`{"name": "oko"}`)
	obj, err = AppFromJson(data)
	assert.Nil(t, err)
	assert.Equal(t, obj.Name, "oko")

	//assert.Equal(t, "{\"errors\":[{\"field\":\"\",\"message\":\"Invalid JSON object\"}]}", err.Error())
}

func TestIntegration_AppRP_Create(t *testing.T) {
	// setup
	tests.GetTestStorage()

	var err error
	var cn = tests.TEST_CUSTOMER
	cc := &core.CustomerContext{CustomerName: cn}

	appRp := NewAppRP(cc)

	// create device on non existing customer
	app := &model.App{}
	app.Name = "xo"
	app.Os = OS_ANDRIOD
	app.CustomerName = cn
	err = appRp.Create(app)
	assert.Equal(t, core.ErrorFrom(core.ErrNotFound, "Customer not found"), err, "#1")

	// create customer
	customer, err := CreateCustomer(cn)
	fmt.Println(customer)
	assert.Equal(t, nil, err)
	assert.Equal(t, customer.CustomerName, cn, "#2")

	// create app and device
	appNew, err := CreateAndroidApp(cc, "", "")
	assert.True(t, err == nil)
	assert.Equal(t, appNew.Os, OS_ANDRIOD)
}
