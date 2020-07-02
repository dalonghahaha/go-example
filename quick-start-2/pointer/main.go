package main

import "fmt"

func main() {
	//声明实际变量
	var a int = 20
	fmt.Println("输出变量的值:", a)
	fmt.Println("输出变量的地址:", &a)
	fmt.Println()

	//声明指针变量
	var p *int
	fmt.Println("输出指针变量的值:", p)
	fmt.Println("输出指针变量的地址:", &p)
	fmt.Println()

	//指针变量赋值
	p = &a
	fmt.Println("输出指针变量的值:", p)
	fmt.Println("输出指针变量的地址:", &p)
	fmt.Println("输出指针变量指向的变量的值", *p)
	fmt.Println("输出指针变量指向的变量的地址", &*p)
	fmt.Println()
}
