package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("defer main")
		if err := recover(); err != nil {
			fmt.Println("recover main:", err)
			return
		}
	}()
	fmt.Println("begin")
	panicWithRecover(10)
	panicWithUnRecover(100000)
	time.Sleep(5 * time.Second)
	fmt.Println("finished!")
}

func panicWithRecover(max int) {
	defer func() {
		fmt.Println("defer panicWithRecover")
		if err := recover(); err != nil {
			fmt.Println("recover panicWithRecover:", err)
		}
	}()
	index := 0
	for {
		if index > max {
			panic(fmt.Sprintf("index is bigger than %d", max))
		} else {
			index++
		}
	}
}

func panicWithUnRecover(max int) {
	defer func() {
		fmt.Println("defer panicWithUnRecover")
	}()
	index := 0
	for {
		if index > max {
			panic(fmt.Sprintf("index is bigger than %d", max))
		} else {
			index++
		}
	}
}
