package main

import (
    "fmt"
    "time"
)


type job struct {
    name string
}

func(j job) doing() {
    fmt.Println("I'am ", j.name, ", now working。。。")
}

var jobQueue chan job

type worker struct {
    name string
    workpool chan chan job
    currentJob chan job
    close chan bool
}

func newWorker(name string, workpool chan chan job) *worker {
    return &worker{
        name: name,
        workpool:workpool,
        currentJob: make(chan job),
        close: make(chan bool),
    }
}

func (w *worker) start() {
    go func() {
       for {
           w.workpool <- w.currentJob
           select {
           case j := <-w.currentJob:
                fmt.Println(w.name,"接收到任务")
               j.doing()
           case <- w.close:
               return
           }
       }
    }()
}

func (w *worker) stop() {
    go func() {
       w.close <- true
    }()
}

type dispatcher struct {
    name string
    workpool chan chan job
    maxWorker int
}

func newDispatcher(max int) *dispatcher {
    return &dispatcher{
        name: "hello",
        workpool: make(chan chan job,max),
        maxWorker:max,
    }
}

func (d *dispatcher) run() {
    for i := 0; i < d.maxWorker; i++ {
        w := newWorker(fmt.Sprintf("worker-%d", i),d.workpool)
        w.start()
    }
    
    go d.dispatch()
}

func (d *dispatcher) dispatch() {
    for {
        select {
        case j := <-jobQueue:
            fmt.Println("调度者,接收到一个工作任务")
            go func( j job) {
                jobchannel := <- d.workpool
                jobchannel <- j
            }(j)
        default:
            
        }
    }
}

func init() {
    maxWorker := 4
    queueMax := 4
    dispatch := newDispatcher(maxWorker)
    jobQueue = make(chan job, queueMax)
    dispatch.run()
}

func main() {
    for i := 0 ; i < 16 ; i++ {
        j := job{
            name: fmt.Sprintf("job--%d",i),
        }
        jobQueue <- j
    }
    time.Sleep(5*time.Second)
    close(jobQueue)
}
