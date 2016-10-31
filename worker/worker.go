package worker

import (
    "fmt"
    "github.com/supple/gorest/core"
    _ "github.com/supple/gorest/services"
    "github.com/supple/gorest/services"
)

// Job holds the attributes needed to perform unit of work.
type Job struct {
    Name    string
    Payload map[string]interface{}
    Value   interface{}
}

func NewJob(name string, p map[string]interface{}, value interface{}) (Job) {
    return Job{Value:value, Name:name, Payload:p}
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job, app *core.AppServices) *Worker {
    return &Worker{
        id:         id,
        // create job queue for worker
        jobQueue:   make(chan Job),
        // set global worker poll
        workerPool: workerPool,
        app:        app,
        quitChan:   make(chan bool),
    }
}

type Worker struct {
    id         int
    jobQueue   chan Job
    workerPool chan chan Job
    quitChan   chan bool
    app        *core.AppServices
}

//func setValue(app *core.AppServices, job *Job) (*core.AppError)  {
//	err := app.SaveEvent(job.Payload)
//	if (err != nil) {
//		fmt.Println("Error: "+err.Error())
//		return &core.AppError{Error:err, Message:"Error adding value to queue", Code: 3001}
//	}
//
//	return nil;
//}

func (w *Worker) doWork(app *core.AppServices, job Job) {
    services.SaveEvent(job.Value)

    //setValue(sc, &job)

    // decrement global counter
    /*
    w.app.Cnt.Decr()
    if (w.app.Cnt.Items == 0) {
        fmt.Println("All work is done")
        w.app.Quit <- true
    }
    */

    //fmt.Fprintln("Do work: %s", job.Name)
}

// run worker
func (w *Worker) start(app *core.AppServices) {
    go func() {
        for {
            // Add my jobQueue to the worker pool.
            w.workerPool <- w.jobQueue

            // and wait for job and do it
            select {
            case job := <-w.jobQueue:
            // Dispatcher has added a job to my jobQueue.
            //fmt.Printf("worker%d: started %s, blocking for %f seconds\n", w.id, job.Name, job.Delay.Seconds())
                w.doWork(app, job)

            //fmt.Printf("worker%d: completed %s!\n", w.id, job.Name)
            case <-w.quitChan:
            // We have been asked to stop.
                fmt.Printf("Worker %d stopping\n", w.id)
                return
            }
        }
    }()
}

func (w Worker) stop() {
    go func() {
        w.quitChan <- true
    }()
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
    workerPool := make(chan chan Job, maxWorkers)

    return &Dispatcher{
        jobQueue:   jobQueue,
        maxWorkers: maxWorkers,
        workerPool: workerPool,
    }
}

type Dispatcher struct {
    workerPool chan chan Job
    maxWorkers int
    jobQueue   chan Job
}

func (d *Dispatcher) Run(app *core.AppServices) {
    for i := 0; i < d.maxWorkers; i++ {
        worker := NewWorker(i + 1, d.workerPool, app)
        worker.start(app)
        fmt.Sprintf("Worker %d start\n", worker.id)
    }
}

func (d *Dispatcher) Dispatch() {
    for {
        select {
        case job := <-d.jobQueue:
            go func() {
                //fmt.Printf("fetching workerJobQueue for: %s\n", job.Name)
                // get available worker
                workerJobQueue := <-d.workerPool
                //fmt.Printf("adding %s to workerJobQueue\n", job.Name)

                // send job to this worker
                workerJobQueue <- job
            }()
        }
    }
}