package services

import (
	"github.com/supple/gorest/storage"
	"log"
)

func SaveEvent(data interface{}) (error) {
	err := storage.GetInstance("events").Coll("events").Insert(data)
	if (err != nil) {
		log.Println("Error: " + err.Error());
	}

	return err
}
//
//func SaveEvent(data map[string]interface{}) (error) {
//	cl, err := storage.GetRedisCluster().GetConn(storage.CLUSTER_TEST);
//	if (err == nil) {
//		jsonString, _ := json.Marshal(data)
//		cl.Set("test", jsonString, 0);
//
//		for i := 0; i < 200; i++ {
//			k := strconv.Itoa(i);
//			s, _ := data["reqID"].(string);
//			_, err := cl.SAdd("testX", s+k+"oo")
//			if (err == nil) {
//				//log.Println("ELEMENTS: "+strconv.FormatInt(ret, 10));
//			} else {
//				log.Println("Error: " + err.Error());
//			}
//
//		}
//	} else {
//		println("ERROR: "+err.Error())
//	}
//
//	return err
//}