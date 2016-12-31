package tests

import (
    "github.com/supple/gorest/storage"
)

var TEST_CUSTOMER = "ctest"

func GetTestStorage() {
    db := storage.NewMongoDB("192.168.1.106:27017", "unittest")
    storage.SetInstance("crm", db)
    storage.DropDatabase(db)
}

