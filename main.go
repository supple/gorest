package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
    "log"
    "github.com/supple/gorest/resources"
    "os"
    "os/signal"
    "syscall"
)


type AppService struct {
    Storage *MemStorage
}

var app = AppService{}

func jsonResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    encoder := json.NewEncoder(w)
    encoder.Encode(data)
}

func InitCache(app *AppService) {
    app.Storage = NewMemStorage()
}
//
//func CampaignGet(c web.C, w http.ResponseWriter, r *http.Request) {
//    var id = c.URLParams["id"]
//    obj := app.Storage.Get(id)
//    jsonResponse(w, obj)
//}
//
//func CampaignList(c web.C, w http.ResponseWriter, r *http.Request) {
//    names := app.Storage.GetByCriteria()
//    jsonResponse(w, names)
//}
//
//func CampaignAdd(c web.C, w http.ResponseWriter, r *http.Request) {
//    obj := resources.Device{}
//    decoder := json.NewDecoder(r.Body)
//    if err := decoder.Decode(&obj); err != nil {
//        log.Print(err.Error())
//        http.Error(w, http.StatusText(422), 422)
//        return
//    }
//    app.Storage.Set(obj.Id, &obj)
//    jsonResponse(w, obj)
//}
//
//// https://github.com/quintans/goSQL#update
//func CampaignUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    defer r.Body.Close()
//    tmp := make(map[string]interface{})
//    if err := decoder.Decode(&tmp); err != nil {
//        http.Error(w, http.StatusText(422), 422)
//        return
//    }
//
//    var id = c.URLParams["id"]
//    obj := app.Storage.Get(id).(*resources.Device)
//    resources.UpdateModel(obj, tmp)
//
//    app.Storage.Set(id, obj)
//    jsonResponse(w, obj)
//}

func signalCatcher() chan os.Signal {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    return c
}

func main() {
    fmt.Println("v1.0.0")
    InitCache(&app)

    goji.DefaultMux.Abandon(middleware.Logger)

    //goji.Get("/v1/devices/", CampaignList)
    //goji.Get("/v1/devices/:id", CampaignGet)
    //goji.Post("/v1/devices/", CampaignAdd)
	//goji.Patch("/v1/devices/:id", CampaignUpdate)

    goji.Serve()
}

