package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
    "log"
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
    cmp := Device{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&cmp); err != nil {
        log.Print(err.Error())
        http.Error(w, http.StatusText(422), 422)
        return
    }
    app.ca.Set(&cmp)
    jsonResponse(w, cmp)
}

// https://github.com/quintans/goSQL#update
func CampaignUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()
    tmp := make(map[string]interface{})
    if err := decoder.Decode(&tmp); err != nil {
        http.Error(w, http.StatusText(422), 422)
        return
    }

    var id = c.URLParams["id"]
    cmp := app.ca.Get(id)
    cmp.Update(tmp)

    app.ca.Set(cmp)
    jsonResponse(w, cmp)
}

func main() {
    fmt.Println("v1.0.0")
    InitCache(&app)

    goji.DefaultMux.Abandon(middleware.Logger)

    goji.Get("/v1/devices/", CampaignList)
    goji.Get("/v1/devices/:id", CampaignGet)
    goji.Post("/v1/devices/", CampaignAdd)
	goji.Patch("/v1/devices/:id", CampaignUpdate)

    goji.Serve()
}

