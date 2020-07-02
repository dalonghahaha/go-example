package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin")
    for index := 0; index < 10; index++ {
        go sayHello(index)
    }
    time.Sleep(1 * time.Second)
    fmt.Println("end!")
}

func sayHello(index int) {
    fmt.Println("hello ", index)
}