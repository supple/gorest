package main

import (
    "testing"
    "github.com/supple/gorest/core"
    a "github.com/stretchr/testify/assert"
    s "github.com/supple/gorest/storage"
    r "github.com/supple/gorest/resources"
    lc "github.com/supple/gorest/utils"
    "gopkg.in/mgo.v2"
    "fmt"
)

func TestFlow(t *testing.T) {
    //
    var err error
    var c *r.Customer

    // init useful variables
    var id = "67158007-b5ff-495f-83bf-36867429a731"
    var apiKeyStr = "OiBTGDVxmZnZHAITDMjqyQRJ-cElsforb"
    var customerName = "customer_test"
    var at = r.AccessTo{Resource:"device", Action: "create"}
    var db = s.GetInstance("entities")

    var cc *core.CustomerContext = &core.CustomerContext{CustomerName: customerName}

    cRp := r.NewCustomerRP(cc)
    akRp := r.NewApiKeyRP(cc)

    // clean
    s.DropCollection(db, cRp.CollectionName())
    s.DropCollection(db, akRp.CollectionName())

    // helper function, error should be nil if not print it
    enil := func(value interface{}) {
        a.True(t, value == nil)
        if (value != nil) {
            fmt.Println(value)
        }
    }
    // set up db
    err = cRp.Install(db)
    enil(err)

    err = akRp.Install(db)
    enil(err)

    // find non existing customer
    c, err = cRp.FindOne(db, id)
    a.True(t, err == mgo.ErrNotFound)

    // save customer
    model := &r.Customer{}
    model.Id = id
    model.CustomerName = customerName
    err = cRp.Create(db, model)
    enil(err)

    // repeat with new id
    model.Id = lc.NewId()
    err = cRp.Create(db, model)
    a.True(t, err != nil)

    // find by first id
    c, err = cRp.FindOne(db, id)
    a.True(t, c.CustomerName == customerName)
    a.True(t, err == nil)

    // create api key
    ak := &r.ApiKey{Key: apiKeyStr}
    ak.CustomerName = model.CustomerName
    err = akRp.Create(db, ak)
    if (err != nil) {
        fmt.Println(err)
    }
    a.True(t, err == nil)

    // authenticate
    ccAuth, err := Auth(db, apiKeyStr, at)
    if (err != nil) {
        fmt.Println(err)
    }
    a.True(t, ccAuth != nil)

    // delete api key
    err = akRp.Delete(db, ak.Id)
    enil(err)

    // delete customer
    err = cRp.Delete(db, id)
    enil(err)
}