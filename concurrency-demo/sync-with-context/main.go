package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin")
    ctx, cancel := context.WithCancel(context.Background())
    go sayHello(cancel)
    <-ctx.Done()
    fmt.Println("finish!")
}

func sayHello(cancel func()) {
    for index := 0; index < 10; index++ {
        fmt.Println("hello ", index)
        time.Sleep(1 * time.Second)
    }
    cancel()
}