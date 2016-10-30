package events

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"time"
	"github.com/supple/gorest/worker"
)



type GEvent struct {
	Name string `json:"name"`
	Time string `json:"time"`

	EventId string `json:"eventId,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`

	Data string `json:"data,omitempty"`
	Request string `json:"request"`
}

func GEnqueueWorker(c web.C, w http.ResponseWriter, r *http.Request) {
	hello := map[string]interface{}{}

	hello["reqID"] = middleware.GetReqID(c)
	hello["test"] = "xo"


	job := worker.Job{Payload: hello, Name: "test"}
	EventJobQueue <- job
	jsonResponse(w, hello)
}

func GMyHandlerReport(c web.C, w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Microsecond)

	x := middleware.GetReqID(c)
	fmt.Fprintf(w, "Reqv1 id: %v, %v\n", x, c.URLParams["id"])
}

func GMyHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", c.URLParams["id"])
}

func jsonResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

type GEventHandler struct { }
func (eh * GEventHandler) GServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there %s!", r.URL.Path[1:])
}