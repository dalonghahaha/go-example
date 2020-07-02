package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    fmt.Println("begin")
    var wg sync.WaitGroup
    for index := 0; index < 10; index++ {
        wg.Add(1)
        go sayHello(index, &wg)
    }
    wg.Wait()
    fmt.Println("finish!")
}

func sayHello(index int, wg *sync.WaitGroup) {
    sleep := time.Duration(index) * time.Second
    time.Sleep(sleep)
    fmt.Println("hello ", index)
    wg.Done()
}