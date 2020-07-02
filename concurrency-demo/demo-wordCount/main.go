package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/afero"
)

var fileSystem = afero.NewOsFs()

func main() {
	fmt.Println("begin")
	dir := "/Users/dengjialong/git/"
	paths, err := readDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("文件数:", len(paths))
	t1 := time.Now()
	total := 0
	for _, path := range paths {
		lineCount(path, &total)
	}
	t2 := time.Now()
	fmt.Println(total)
	fmt.Println("正常计算总耗时:", t2.Sub(t1).Seconds())
	t3 := time.Now()
	total2 := 0
	stop := make(chan bool)
	go async(paths, &total2, stop)
	<-stop
	t4 := time.Now()
	fmt.Println(total2)
	fmt.Println("并行计算总耗时:", t4.Sub(t3).Seconds())
	t5 := time.Now()
	total3 := 0
	stop2 := make(chan bool)
	go asyncWithLock(paths, &total3, stop2)
	<-stop2
	t6 := time.Now()
	fmt.Println(total3)
	fmt.Println("并行计算总耗时:", t6.Sub(t5).Seconds())
	fmt.Println("end!")
}

func readDir(path string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Mode().IsRegular() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return files, nil
}

func async(paths []string, count *int, ch chan bool) {
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go asyncLineCount(path, count, &wg)
	}
	wg.Wait()
	ch <- true
}

func asyncWithLock(paths []string, count *int, ch chan bool) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, path := range paths {
		wg.Add(1)
		go asyncLineCountWithLock(path, count, &wg, &lock)
	}
	wg.Wait()
	ch <- true
}

func lineCount(path string, count *int) {
	file, err := os.Open(path)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
	}
	stat, _ := file.Stat()
	*count += int(stat.Size())
}

func asyncLineCount(path string, count *int, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
	}
	stat, _ := file.Stat()
	*count += int(stat.Size())
	wg.Done()
}

func asyncLineCountWithLock(path string, count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	file, err := os.Open(path)
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
	}
	stat, _ := file.Stat()
	lock.Lock()
	*count += int(stat.Size())
	lock.Unlock()
	wg.Done()
}
