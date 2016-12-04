package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "syscall"
    "strconv"
    "os"
    "os/signal"
    "github.com/gin-gonic/gin"
    "github.com/supple/gorest/core"
    "github.com/supple/gorest/storage"
    "github.com/supple/gorest/worker"
    "github.com/supple/gorest/handlers"
    "github.com/supple/gorest/resources"
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

func init() {
    // Init storage instances
    storage.SetInstance("crm", storage.NewMongoDB("192.168.1.106:27017", "crm"))
    //storage.SetInstance("entities", storage.NewMongoDB("192.168.1.106:27017", "entities"))
    storage.SetInstance("events", storage.NewMongoDB("192.168.1.106:27017", "events"))

    // Create the job queue.
    var maxQueueSize = 3
    var maxWorkers = 50
    if len(os.Args) > 2 {
        maxQueueSize, _ = strconv.Atoi(os.Args[1]) // 3
        maxWorkers, _ = strconv.Atoi(os.Args[2])   // 50
    }

    worker.EventJobQueue = make(chan worker.Job, maxQueueSize)

    // Start the dispatcher.
    d := worker.NewDispatcher(worker.EventJobQueue, maxWorkers)
    d.Run(&app)

    // dispatch worker producers
    go d.Dispatch()
}

func signalCatcher() chan os.Signal {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    return c
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.Abort()
            return
        }
        c.Next()
    }
}

func AuthMiddleware(c *gin.Context) {
    //log.Print("[x] Request\n")
    // zbCrVUXQSLseDVruJIBwYgke-cRaddKsc
    ac := resources.AccessTo{Resource: "test", Action:"test"}
    cc, err := resources.Auth(c.Request.Header.Get("API-KEY"), ac)

    if err == nil {
        c.Set("cc", cc)
        c.Next()
        return
    }

    // handle error
    //log.Print()

    switch err.(type) {
    case *core.ApiError:
        ae := err.(*core.ApiError)
        c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
        c.AbortWithError(ae.Status, ae)
        return
    case error:
        c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
        c.AbortWithError(500, err)
        return
    }


}

func ErrorHandler(c *gin.Context) {
    c.Next()
    if len(c.Errors) > 0 {
        c.JSON(-1, c.Errors) // -1 == not override the current error code
    }
}

func main() {
    fmt.Println("v1.0.0")
    InitCache(&app)

    gin.SetMode(gin.ReleaseMode)
    r := gin.New()
    r.Use(gin.Recovery())
    r.Use(gin.Logger())
    r.Use(AuthMiddleware)
    r.Use(ErrorHandler)

    //r.Use(gzip.Gzip(gzip.DefaultCompression))
    //r.Use(CORSMiddleware())
    ca := handlers.CustomerApi{}
    d := handlers.DeviceApi{}

    //r := gin.Default()
    v1 := r.Group("api/v1")
    {
        v1.POST("/events", handlers.HandleEvents)

        v1.GET("/customers", ca.Get)
        v1.POST("/customers", ca.Post)

        v1.GET("/devices/:id", d.Get)
        v1.POST("/devices", d.Post)
    }
    // customers
    // api-keys
    // apps
    // devices

    r.Run(":8080")
}
