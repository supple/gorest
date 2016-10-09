package main

import (
	"testing"
	a "github.com/stretchr/testify/assert"
	s "github.com/supple/gorest/storage"
	r "github.com/supple/gorest/resources"
	lc "github.com/supple/gorest/utils"
	"gopkg.in/mgo.v2"
	"fmt"
)


func TestFlow(t *testing.T) {

	at := r.AccessTo{Resource:"device", Action: "create"}
	//cc := Auth("xoz", at)

	db := s.NewMongo("localhost:27017", "lcache")

	cRp := r.NewCustomerRP()
	akRp := r.NewApiKeyRP()

	var err error
	var c *r.Customer

	enil := func(value interface{}) {
        a.True(t, value == nil)
        if (value != nil) { fmt.Println(value) }
    }
	// set up db
	err = cRp.Install(db)
    enil(err)

	err = akRp.Install(db)
	if (err != nil) { fmt.Println(err) }

	// delete if there is some old ones
	id := "67158007-b5ff-495f-83bf-36867429a731"
	apiKeyStr := "OiBTGDVxmZnZHAITDMjqyQRJ-cElsforb"

	// find non existing
	c, err = cRp.FindOne(db, id)
	a.True(t, err == mgo.ErrNotFound)
	//enil(c)

	// save
	model := &r.Customer{}
	model.Id = id
	model.Name = "marek"
	err = cRp.Create(db, model)
    enil(err)

	// repeat
	model.Id = lc.NewId()
	err = cRp.Create(db, model)
	a.True(t, err != nil)

	// find
	c, err = cRp.FindOne(db, id)
	a.True(t, c.Name == "marek")
	a.True(t, err == nil)

	// create api key
	ak := &r.ApiKey{CustomerName: model.Name, Key: apiKeyStr}
	err = akRp.Create(db, ak)
	if (err != nil) { fmt.Println(err) }
	a.True(t, err == nil)

	fmt.Println("ApiKey: " + ak.Key)

	cc, err := Auth(db, apiKeyStr, at)
    if (err != nil) { fmt.Println(err) }
	a.True(t, cc != nil)

	fmt.Println("ApiKey Id: " + ak.Id)
    akRp.Delete(db, ak.Id)

    // delete customer
    err = cRp.Delete(db, id)
    enil(err)
}