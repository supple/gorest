package events

import (
	"github.com/supple/gorest/storage"
	"log"
	"gopkg.in/mgo.v2"
)
var coll *mgo.Collection

func init() {
	coll = storage.GetInstance("events").Coll("events")
}

func SaveEvent(data interface{}) (error) {
	err := coll.Insert(data)
	if (err != nil) {
		log.Println("[ERROR] Error: " + err.Error());
	}
    //fmt.Println("Saved", data.(*Event).Id)

	return err
}
