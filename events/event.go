package events

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/rs/xid"
    "github.com/supple/gorest/worker"
)

var EventJobQueue chan worker.Job
type Payload map[string]interface{}

type Event struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
    Time     string `json:"time, omitempty"` // joda time
    AppId    string `json:"appId"`
    DeviceId string `json:"deviceId"`
    Payload  Payload `json:"payload, omitempty"` //Payload json.RawMessage `json:"payload, omitempty"`

    Request string `json:"request"`
}

func HandleEvents(c *gin.Context) {
    var e Event

    c.Bind(&e)
    e.Id = xid.New().String()
    e.Time = time.Now().Format("2006-01-02T15:04:05.999Z")

    // -> add to queue
    // check flow (context: {flowId, deviceId}) if no flow start timeout ->
    // on beacon_entered, mac: "abc", time.hour: <15, 17>   -> setFlowStage() -> setFlowStartLock(timeout=10h)
    // -> make flow action (route_info params:{DirectionsRequest} ) -> (If DurationInTraffic > Y) -> next task
    // -> make flow action (send_ntf params:{route_msg {placeholders}})
    //
    job := worker.Job{Name: "test", Value: &e}
    EventJobQueue <- job
    print(&e)

    //content := map[string]interface{}{"status": "ok"}
    c.JSON(201, e)
}

//type Events []Event
//
//type EventService struct {
//    events chan Event
//    mtx sync.RWMutex
//}
//
//func NewEventService() *EventService {
//    return &EventService{
//        events: make(chan Event, 500),
//    }
//}

// workers ... slice -> pack -> send
//
//func (es *EventService) flushLogs() {
//    defer func() {
//        close(es.events)
//    }()
//
//    ticker := time.NewTicker(time.Duration(600) * time.Second)
//
//    for {
//        select {
//        case ev := <-r:
//            es.mtx.Lock()
//            es.events[ev.Source] = time.Now()
//            es.mtx.Unlock()
//        case <-ticker.C:
//            t := time.Now()
//            ws.mtx.Lock()
//            for log, ttl := range ws.logs {
//                if t.Sub(ttl).Seconds() > 600 {
//                    delete(ws.logs, log)
//                }
//            }
//            ws.mtx.Unlock()
//        }
//    }
//}
//
//func (es *EventService) Emit(name string, deviceId string, params P) bool {
//    // send to queue
//
//    time.Now()
//
//    return true
//}
//
//func AddToPack() {
//    // events pack: up to 500 events,
//    pack:= P{
//        "meta": {"test": ""},
//    }
//    fmt.Println(pack)
//
//}