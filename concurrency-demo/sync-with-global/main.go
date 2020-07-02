package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin")
    finish := false
    go sayHello(&finish)
    for !finish {
        time.Sleep(1 * time.Second)
    }
    fmt.Println("finish!")
}

func sayHello(finish *bool) {
    for index := 0; index < 10; index++ {
        fmt.Println("hello ", index)
        time.Sleep(1 * time.Second)
    }
    *finish = true
}