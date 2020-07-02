package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin")
    stop := make(chan bool)
    go sayHello(stop)
    <-stop
    fmt.Println("finish!")
}

func sayHello(ch chan bool) {
    for index := 0; index < 10; index++ {
        fmt.Println("hello ", index)
        time.Sleep(1 * time.Second)
    }
    ch <- true
}