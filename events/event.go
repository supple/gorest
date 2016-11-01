package events

import (
    _ "github.com/rs/xid"
)

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