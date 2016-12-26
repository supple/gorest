package sender

import (
    "github.com/supple/gorest/resources"
    "net/http"
    "time"
)

type GcmMessage struct {
    To string `json:"to,omitempty"`
    Priority string `json:"priority,omitempty"`
    TimeToLive uint `json:"time_to_live,omitempty"`
    Data string `json:"data,omitempty"`
}

type Message struct {
    Title string
    Body string
    Params map[string]interface{}
}

func Send(msg Message, app resources.App) {

}

type Client struct {
    url string
    timeout int32
    cl *http.Client
}

func NewClient(timeout int32) *Client {
    var cl *http.Client = &http.Client{Timeout: time.Second * time.Duration(timeout)}
    return &Client{cl: cl}
}

func (c *Client) Post() {

}