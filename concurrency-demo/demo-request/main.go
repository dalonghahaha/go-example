package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("begin")
	t1 := time.Now()
	total := 0
	normal(&total)
	t2 := time.Now()
	fmt.Println(total)
	fmt.Println("正常计算总耗时:", t2.Sub(t1).Seconds())
	t3 := time.Now()
	total2 := 0
	stop := make(chan bool)
	go async(&total2, stop)
	<-stop
	t4 := time.Now()
	fmt.Println(total2)
	fmt.Println("并行计算总耗时:", t4.Sub(t3).Seconds())
	t5 := time.Now()
	total3 := 0
	stop2 := make(chan bool)
	go asyncWithLock(&total3, stop2)
	<-stop2
	t6 := time.Now()
	fmt.Println(total3)
	fmt.Println("并行计算总耗时:", t6.Sub(t5).Seconds())
	fmt.Println("end!")
}

func normal(count *int) {
	for index := 0; index < 100; index++ {
		request(count, strconv.Itoa(index))
	}
}

func async(count *int, ch chan bool) {
	var wg sync.WaitGroup
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go syncRequest(count, strconv.Itoa(index), &wg)
	}
	wg.Wait()
	ch <- true
}

func asyncWithLock(count *int, ch chan bool) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go syncRequestWithLock(count, strconv.Itoa(index), &wg, &lock)
	}
	wg.Wait()
	ch <- true
}

func request(count *int, id string) {
	response, err := http.Get("http://localhost:10086/employee/?id=" + id)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		*count += len(body)
	}
}

func syncRequest(count *int, id string, wg *sync.WaitGroup) {
	response, err := http.Get("http://localhost:10086/employee/?id=" + id)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		*count += len(body)
	}
	wg.Done()
}

func syncRequestWithLock(count *int, id string, wg *sync.WaitGroup, lock *sync.Mutex) {
	response, err := http.Get("http://localhost:10086/employee/?id=" + id)
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err == nil {
		body, _ := ioutil.ReadAll(response.Body)
		lock.Lock()
		*count += len(body)
		lock.Unlock()
	}
	wg.Done()
}
