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

	rp := r.CustomerRP{}
	akRp := r.ApiKeyRP{}

	var err error
	var c *r.Customer

	err = rp.Install(db)
	if (err != nil) { fmt.Println(err) }

	err = akRp.Install(db)
	if (err != nil) { fmt.Println(err) }

	// delete if there is some old ones
	id := "67158007-b5ff-495f-83bf-36867429a731"
	apiKeyStr := "pOlqDsToiIdZAoCmNexYNubE-sozgdHsN"

	err = rp.Delete(db, id)
	a.True(t, err == nil)

	// find non existing
	c, err = rp.FindOne(db, id)
	a.True(t, err == mgo.ErrNotFound)
	a.True(t, c == nil)

	// save
	model := &r.Customer{}
	model.Id = id
	model.Name = "marek"
	err = rp.Create(db, model)
	a.True(t, err == nil)

	// repeat
	model.Id = lc.NewId()
	err = rp.Create(db, model)
	a.True(t, err != nil)

	// find
	c, err = rp.FindOne(db, id)
	a.True(t, c.Name == "marek")
	a.True(t, err == nil)

	// create api key
	ak := &r.ApiKey{CustomerName: model.Name}
	err = akRp.Create(db, ak)
	if (err != nil) { fmt.Println(err) }
	a.True(t, err == nil)

	fmt.Println("ApiKey: " + ak.Key)


	cc, err := Auth(db, apiKeyStr, at)
	a.True(t, cc != nil)

	fmt.Println(ak.Id)
}