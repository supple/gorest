package main

import (
	"gopkg.in/mgo.v2/bson"
	s "github.com/supple/gorest/storage"
	r "github.com/supple/gorest/resources"
)


// User represents the structure of our resource
type Device struct {
	Id       string `json:"id" bson:"_id"`
	AppId    string        `json:"name" bson:"name" `
	AppToken string        `json:"appToken" bson:"appToken"`
}

type App struct {
	Id     string `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	GcmKey string        `json:"gcmKey" bson:"gcmKey"`
}



//// CreateUser creates a new user resource
//func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Stub an user to be populated from the body
//	u := models.User{}
//
//	// Populate the user data
//	json.NewDecoder(r.Body).Decode(&u)
//
//	// Add an Id
//	u.Id = bson.NewObjectId()
//
//	// Write the user to mongo
//	uc.session.DB("go_rest_tutorial").C("users").Insert(u)
//
//	// Marshal provided interface into JSON structure
//	uj, _ := json.Marshal(u)
//
//	// Write content-type, statuscode, payload
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(201)
//	fmt.Fprintf(w, "%s", uj)
//}


func Auth(db *s.Mongo, apiKey string, accessTo r.AccessTo) (*r.CustomerContext, error) {
	var cc *r.CustomerContext
	akRp := r.ApiKeyRP{}
	ak, err := akRp.FindOneBy(db, bson.M{"key": apiKey})
	// @todo: hasAccess(accessTo)
	if err == nil {
		cc = &r.CustomerContext{}
		cc.ApiKey = ak.Key
		cc.CustomerName = ak.CustomerName
	}

	return cc, err
}
