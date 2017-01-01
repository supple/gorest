package tests

import (
    "github.com/supple/gorest/storage"
)

var TEST_CUSTOMER = "ctest"

func GetTestStorage() {
    db := storage.NewMongoDB("127.0.0.1:27017", "unittest")
    storage.SetInstance("crm", db)
    storage.DropDatabase(db)
}

