package main

import (
	"fmt"
	"github.com/supple/gorest/core"
	"github.com/supple/gorest/server"
	"github.com/supple/gorest/storage"
	"github.com/supple/gorest/worker"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var app = core.AppServices{}

func initServices(app *core.AppServices) {
	// Init storage instances: 192.168.1.106 mongo-crm
	storage.SetInstance("crm", storage.NewMongoDB("mongo-crm:27017", "crm"))
	storage.SetInstance("events", storage.NewMongoDB("mongo-events:27017", "events"))
	app.Storage = storage.NewMemStorage()

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
	d.Run(app)

	// dispatch worker producers
	go d.Dispatch()
}

func signalCatcher() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return c
}

func main() {
	initServices(&app)
	signalCatcher()
	fmt.Println("v1.0.0")

	r := server.SetupRouter()
	r.Run(":8080")
}
