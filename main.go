package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    _ "log"

    "github.com/gin-gonic/gin"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/storage"
    "os"
    "os/signal"
    "syscall"
    "github.com/supple/gorest/events"
    "github.com/supple/gorest/worker"
)

var app = core.AppServices{}

func jsonResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    encoder := json.NewEncoder(w)
    encoder.Encode(data)
}

func InitCache(app *core.AppServices) {
    app.Storage = storage.NewMemStorage()
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

func init() {

    // Create the job queue.
    maxQueueSize := 50
    maxWorkers := 5

    events.EventJobQueue = make(chan worker.Job, maxQueueSize)

    // Start the dispatcher.
    d := worker.NewDispatcher(events.EventJobQueue, maxWorkers)
    d.Run(&app)

    // dispatch worker producers
    go d.Dispatch()
}

func main() {
    fmt.Println("v1.0.0")
    InitCache(&app)
    r := gin.Default()
    v1 := r.Group("api/v1")
    {
        v1.POST("/events", events.HandleEvents)
    }

    r.Run(":8080")
}
