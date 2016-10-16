package core

import (
    "fmt"
    "time"
    "sync"
)

type P map[string]interface{}

type Event struct {
    Name string
    Time string
    DeviceId string
    AppId string
    Os string
    // Device: createdAt, os, appId
    // Flow: Id, StepId,
    Payload P
}

type Events []Event

type EventService struct {
    events chan Event
    mtx sync.RWMutex
}

func NewEventService() {
    return &EventService{
        events: new(chan Event, 500),
    }
}

// workers ... slice -> pack -> send

func (es *EventService) flushLogs() {
    defer func() {
        close(es.events)
    }()

    ticker := time.NewTicker(time.Duration(600) * time.Second)

    for {
        select {
        case ev := <-r:
            es.mtx.Lock()
            es.events[ev.Source] = time.Now()
            es.mtx.Unlock()
        case <-ticker.C:
            t := time.Now()
            ws.mtx.Lock()
            for log, ttl := range ws.logs {
                if t.Sub(ttl).Seconds() > 600 {
                    delete(ws.logs, log)
                }
            }
            ws.mtx.Unlock()
        }
    }
}

func (es *EventService) Emit(name string, deviceId string) bool {
    // send to queue

    time.Now()

    return true
}

func AddToPack() {
    // events pack: up to 500 events,
    pack:= P{
        "meta": {"test": ""},
    }
    fmt.Println(pack)

}