package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/events"
	"github.com/supple/gorest/worker"
	"time"
)

func HandleEvents(c *gin.Context) {
	//log.Print("[GIN-debug] Request")
	var e events.Event

	c.Bind(&e)
	e.Id = core.RandString(12) // xid.New().String()
	e.Time = time.Now().Format("2006-01-02T15:04:05.999Z")

	// customer context must exists
	cc := c.MustGet("cc").(*core.CustomerContext)
	e.CustomerName = cc.CustomerName
	e.AppId = cc.AppId

	// -> add to queue
	// check flow (context: {flowId, deviceId}) if no flow start timeout ->
	// on beacon_entered, mac: "abc", time.hour: <15, 17>   -> setFlowStage() -> setFlowStartLock(timeout=10h)
	// -> make flow action (route_info params:{DirectionsRequest} ) -> (If DurationInTraffic > Y) -> next task
	// -> make flow action (send_ntf params:{route_msg {placeholders}})

	//
	job := worker.Job{Name: "test", Value: &e}
	worker.EventJobQueue <- job

	//content := map[string]interface{}{"status": "ok"}
	c.JSON(201, e)
}
