package main

import (
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/utils"
    s "github.com/supple/gorest/storage"
    r "github.com/supple/gorest/resources"
    "fmt"
)

func init() {
    s.SetInstance("crm", s.NewMongoDB("192.168.1.106:27017", "test"))

    //
    var err error

    // init useful variables
    var id = "67158007-b5ff-495f-83bf-36867429a731"
    var apiKeyStr = "OiBTGDVxmZnZHAITDMjqyQRJ-cElsforb"
    var customerName = "customer_test"

    var db = s.GetInstance("crm")

    var cc *core.CustomerContext = &core.CustomerContext{CustomerName: customerName}

    cRp := r.NewCustomerRP(cc)
    akRp := r.NewApiKeyRP(cc)

    // clean
    s.DropCollection(db, cRp.CollectionName())
    s.DropCollection(db, akRp.CollectionName())

    // set up db
    err = cRp.Install(db)
    err = akRp.Install(db)

    // save customer
    model := &r.Customer{}
    model.Id = id
    model.CustomerName = customerName
    err = cRp.Create(model)

    // repeat with new id
    model.Id = utils.NewId()
    err = cRp.Create(model)

    // find by first id
    c, err := cRp.FindOne(id)

    // create api key
    ak := &r.ApiKey{ApiKey: apiKeyStr}
    ak.CustomerName = c.CustomerName
    err = akRp.Create(ak)
    if (err != nil) {
        fmt.Println(err)
    }
}
