package main

import "fmt"

func main() {
	//普通的循环用法
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//使用break和continue
	for i := 0; i < 10; i++ {
		if i < 3 {
			continue
		} else if i > 8 {
			break
		} else {
			fmt.Println(i)
		}
	}

	//类似while的用法
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	//for-range遍历

	a := [3]int{6, 7, 8}
	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}

	b := map[string]string{
		"field1": "jack",
		"field2": "vencent",
		"field3": "mike",
	}
	for k, v := range b {
		fmt.Println(k, v)
	}
	for k, _ := range b {
		fmt.Println(k)
	}
	for _, v := range b {
		fmt.Println(v)
	}
}
