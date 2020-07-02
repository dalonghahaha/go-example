package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("begin")
    ch := make(chan int, 1)
    go readFromChannel(ch)
    go writeToChannel(ch)
    time.Sleep(1 * time.Second)
    fmt.Println("end!")
}

func writeToChannel(ch chan int) {
    for i := 1; i < 10; i++ {
        fmt.Println("write ", i)
        ch <- i
    }
}

func readFromChannel(ch chan int) {
    for i := 1; i < 10; i++ {
        v := <-ch
        fmt.Println("read ", v)
    }
}