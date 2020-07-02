package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    myMap = map[int]interface{}{}
    lock sync.Mutex
)

func SetMap(c int) {
    lock.Lock()
    myMap[c] = c * c
    lock.Unlock()
}
func main() {
    for index := 1; index <= 10; index++ {
        go add(index)
    }
    time.Sleep(1 * time.Second)
    fmt.Println(myMap)
}

func add(val int) {
    SetMap(val)
}