package tests

import (
    "github.com/supple/gorest/storage"
    "github.com/supple/gorest/resources"
    "github.com/supple/gorest/core"
    "fmt"
)

var TEST_CUSTOMER = "ctest"

func GetStorage() {
    db := storage.NewMongoDB("192.168.1.106:27017", "unittest")
    storage.SetInstance("crm", db)
    storage.DropDatabase(db)
}

func CreateTestCustomer()  {
    resources.CreateCustomer(TEST_CUSTOMER)
}

func CreateTestApiKey() *resources.ApiKey {
    var cc = &core.CustomerContext{ApiKey: "", CustomerName: TEST_CUSTOMER}
    apiKey, err := resources.CreateApiKey(cc)
    if err != nil {
        fmt.Println(err.Error())
    }

    return apiKey
}
