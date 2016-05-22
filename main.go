package main

import (
    "fmt"
    "time"
    "expvar"
    "io"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
    mtm "github.com/supple/mtest/middleware"
    "github.com/paulbellamy/ratecounter"
)

var app = AppService{}

func jsonResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    encoder := json.NewEncoder(w)
    encoder.Encode(data)
}

func InitCache(app *AppService) {
    names := []string{"Alpha", "Beta", "Sigma"}
    app.ca = NewCacheArr(names)
}

func CampaignGet(c web.C, w http.ResponseWriter, r *http.Request) {
    var id = c.URLParams["id"]
    campaign := app.ca.Get(id)
    jsonResponse(w, campaign)
}

func CampaignList(c web.C, w http.ResponseWriter, r *http.Request) {
    names := app.ca.GetNames()
    jsonResponse(w, names)
}

func CampaignAdd(c web.C, w http.ResponseWriter, r *http.Request) {
    cmp := Campaign{}
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    if err := decoder.Decode(&cmp); err != nil {
        http.Error(w, http.StatusText(422), 422)
        return
    }
    cmp.Id = RandSeq(5)
    app.ca.Set(&cmp)
    jsonResponse(w, cmp)
}

func CreateUser(c web.C, w http.ResponseWriter, r *http.Request) {
    cmp := Campaign{}
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    if err := decoder.Decode(&cmp); err != nil {
        http.Error(w, http.StatusText(422), 422)
        return
    }
    cmp.Id = RandSeq(5)
    app.ca.Set(&cmp)
    jsonResponse(w, cmp)
}

func metrics(c web.C, w http.ResponseWriter, r *http.Request) {
    mtm.M.Write(w)
    web.GetMatch(c)

    expiration := time.Now().Add(365 * 24 * time.Hour)
    cookie := http.Cookie{Name: "csrftoken",Value:"abcd",Expires:expiration}
    http.SetCookie(w, &cookie)

    //fmt.Println(proto.MarshalTextString(metric))
}


var (
    counter *ratecounter.RateCounter
    hitsperminute = expvar.NewInt("hits_per_minute")
)

func increment(w http.ResponseWriter, r *http.Request) {
    counter.Incr(1)
    hitsperminute.Set(counter.Rate())
    io.WriteString(w, strconv.FormatInt(counter.Rate(), 10))
}

func main() {
    fmt.Println("v1.0.0")
    counter = ratecounter.NewRateCounter(1 * time.Minute)

    InitCache(&app)

    goji.DefaultMux.Abandon(middleware.Logger)
    goji.Use(mtm.Metrics)

    goji.Get("/v1/campaigns/", CampaignList)
    goji.Get("/v1/campaigns/:id", CampaignGet)
    goji.Post("/v1/campaigns/", CampaignAdd)

    goji.Post("/v1/user/", CreateUser)

    goji.Get("/v1/metrics", metrics)
    goji.Get("/v1/increment", increment)

    goji.Serve()
}

