package main

import (
    "time"
    "fmt"
    "runtime"
)

func timeout(f func()) {
    done := make(chan struct{})
    
    judge := true
    go func() {
       for {
           select {
           case <-done:
               runtime.Goexit()
           default:
               if judge{
                   fmt.Println("enter")
                   judge = false
                   f()
                   close(done)
               }
           }
       }
    }()
    
    loop:
    for {
        select {
        case <- time.After(time.Second*3):
            fmt.Println("timeout")
            break loop
            close(done)
        case <-done:
            break loop
        }
    }
}

func main() {
    timeout(msleep1)
    timeout(msleep2)
}


func msleep1() {
    time.Sleep(time.Second * 4)
    fmt.Println("sleep 5")
}

func msleep2() {
    time.Sleep(time.Second * 2)
    fmt.Println("sleep 2")
}
