package sender

import (
	"fmt"
	"github.com/supple/gorest/model"
	"net/http"
	"time"
)

type GcmMessage struct {
	To         string `json:"to,omitempty"`
	Priority   string `json:"priority,omitempty"`
	TimeToLive uint   `json:"time_to_live,omitempty"`
	Data       string `json:"data,omitempty"`
}

type Message struct {
	Title  string
	Body   string
	Params map[string]interface{}
}

func Send(msg Message, app model.App) {
	fmt.Println("ok")
}

type Client struct {
	url     string
	timeout int32
	cl      *http.Client
}

func NewClient(timeout int32) *Client {
	var cl *http.Client = &http.Client{Timeout: time.Second * time.Duration(timeout)}
	return &Client{cl: cl}
}

func (c *Client) Post() {

}
