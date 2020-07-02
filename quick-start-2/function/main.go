package main

import "fmt"

func main() {
	a := 1
	b := 2
	str := "hello world"
	func1()
	fmt.Println("调用func2结果：", func2(str))
	fmt.Println("调用func3结果：", func3(&str))
	fmt.Println("调用func4结果：", func4(a, b, str))
	fmt.Println("调用func5结果：", func5(a, b))
	c, d := func6(str)
	fmt.Println("调用func6结果：", c, d)
	func7(a, b)
	func7(a, b, 1)
	func7(a, b, 1, 2)
	e := []int{1, 2, 3}
	func7(a, b, e...)
	func8(func1)
	func9(func2)
	s := func(x string) int {
		return 0
	}
	func9(s)
}

//无参数
func func1() {
	fmt.Println("调用func1")
}

//有参数和返回值
func func2(x string) int {
	fmt.Println("调用func2,参数:", x)
	return len(x)
}

//指针作为函数参数
func func3(x *string) int {
	fmt.Println("调用func3,参数:", x)
	return len(*x)
}

//相同类型的参数共同声明
func func4(a, b int, c string) bool {
	fmt.Println("调用func4,参数:", a, b, c)
	return a+b == len(c)
}

//返回值定义名字
func func5(a int, b int) (ret bool) {
	fmt.Println("调用func5,参数:", a, b)
	ret = a == b
	return
}

//多个返回值
func func6(a string) (bool, int) {
	fmt.Println("调用func6,参数:", a)
	return len(a) > 5, len(a)
}

//可变参数
func func7(a, b int, c ...int) {
	fmt.Println("调用func7,参数:", a, b, c)
}

//函数作为参数
func func8(callback func()) {
	callback()
}

func func9(callback func(s string) int) {
	a := ""
	callback(a)
}
