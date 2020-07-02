package main

import "fmt"

func main() {
	a := 3

	//单条件语句
	if a > 0 {
		fmt.Println(a)
	}

	if b := 1; a > b {
		//可以直接使用result变量
		fmt.Println(b)
		fmt.Println(a)
	}

	//if-else结构定义
	if a > 0 {
		fmt.Println("a大于0")
	} else {
		fmt.Println("a小于0")
	}

	//多分支if-else结构定义
	if a == 0 {
		fmt.Println("a等于0")
	} else if a == 1 {
		fmt.Println("a等于1")
	} else {
		fmt.Println("a等于其他")
	}

	//switch语句
	switch a {
	case 0:
		fmt.Println("a等于0")
	case 1:
		fmt.Println("a等于1")
	default:
		fmt.Println("a等于其他")
	}

	switch {
	case a == 0:
		fmt.Println("a等于0")
	case a == 1:
		fmt.Println("a等于1")
	default:
		fmt.Println("a等于其他")
	}

	switch b := 1; {
	case a < b:
		fmt.Println(b)
		fmt.Println("a大于b")
	case a < b:
		fmt.Println(b)
		fmt.Println("a小于于b")
	default:
		fmt.Println(b)
		fmt.Println("等于b")
	}
}
