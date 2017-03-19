package main

import (
    "fmt"
    "time"
)

//延时执行的工作
type cronJob struct {
    name  string  //工作的名称
    delay int64   //延时的时间
}

func (c *cronJob) cjobPrint() {
    fmt.Println(c.name, ": 我在工作中...., 我的延时时间为:",c.delay)
}

//放入环形队列工作池节点中的set<Task>
type task struct {
    cronJob
    circleNum int64   //每一个工作需要运行几圈才触发
}

//每一个环形队列工作池的Set
type set []map[string]*task

//环形队列
type circle struct {
    tasks set
    currentIndex int64   //当前环形队列运行的位置
    maxNumber int64    //环形队列的大小
    tick time.Duration  //环形队列每运行一个节点，需要的时间
    stop chan bool
}

func newCircle(maxNumber int64, tick time.Duration) *circle {
    return &circle{
        tasks:     make([]map[string]*task, maxNumber, maxNumber),
        currentIndex: 0,
        maxNumber: maxNumber,
        tick:      tick,
        stop: make(chan bool),
    }
}

func (c *circle) push(j cronJob) {
    //计算工作应该放置在环形队列中的位置
    index := j.delay % c.maxNumber + c.currentIndex
    //计算工作应该在环形队列中运行几圈触发
    num := j.delay / c.maxNumber
    
    //判断需要放置的位置的set是否初始化了
    if c.tasks[index] == nil {
        c.tasks[index] = make(map[string]*task)
    }
    
    //把工作任务放置到环形队列中对应节点的set中，如果已经存在，则不在放置
    if _, ok := c.tasks[index][j.name]; !ok {
        c.tasks[index][j.name] = &task{j,num}
    }
}

func (c *circle) run() {
    tick := time.Tick(c.tick)
    go func() {
       for {
           select {
           case <-tick:
               //判断是否已经超出环形队列的寻址范围，超出的话，重新开始
               if c.currentIndex > c.maxNumber - 1 {
                   c.currentIndex = c.currentIndex % c.maxNumber
               }
                
               //判断当前的环形队列运行的节点位置中的set<task>中的每一个任务圈数是否为0
               //为0，则执行，不为0，则对圈数-1
               if c.tasks[c.currentIndex] != nil { //如果当前的set<task>已经被初始化了，代表里面存在任务
                   for k, v := range c.tasks[c.currentIndex] {
                       if v.circleNum == 0 {
                           v.cjobPrint()
                           delete(c.tasks[c.currentIndex], k)
                       }else {
                           c.tasks[c.currentIndex][k].circleNum--
                       }
                   }
               }
               
               //当前的节点下标移一位
               c.currentIndex++
           case <-c.stop:
               return
           }
       }
    }()
}

func (c *circle) exit () {
    go func() {
       c.stop <- true
    }()
}

func main() {
    c := newCircle(7, time.Second*1)
    c.run()
    
    job1 := cronJob{
        name : "job1",
        delay: 1,
    }
    c.push(job1)
    
    job2 := cronJob{
        name : "job2",
        delay: 3,
    }
    c.push(job2)
    
    job3 := cronJob{
        name : "job3",
        delay: 5,
    }
    c.push(job3)
    
    job4 := cronJob{
        name : "job4",
        delay: 7,
    }
    c.push(job4)
    
    job5 := cronJob{
        name : "job5",
        delay: 9,
    }
    c.push(job5)
    
    time.Sleep(time.Second * 15)
}