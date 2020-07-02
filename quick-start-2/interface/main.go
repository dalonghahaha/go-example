package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//使用接口作为函数参数
	man1 := Man{
		Name: "jack",
		Age:  30,
	}
	SayandWork(man1)
	woman1 := Woman{
		Name:   "jane",
		Age:    25,
		Height: 170,
		Weight: 55,
	}
	SayandWork(woman1)

	//"泛型"map
	map1 := map[string]interface{}{
		"int":    1,
		"string": "hello",
		"bool":   true,
		"man":    man1,
		"women":  woman1,
	}
	fmt.Println(map1)

	//"泛型"参数函数
	genericityCall(1)
	genericityCall("hello")
	genericityCall(true)
	genericityCall(man1)
	genericityCall(woman1)

	//类型断言
	str := "I say hehe to you"
	fmt.Println(checkUnSafe(str))
	fmt.Println(checkSafe(str))
	//fmt.Println(checkUnSafe(man1)) //此处运行时会报错
	fmt.Println(checkSafe(man1))

	//类型开关
	typeSwitch(1)
	typeSwitch("hello world")
	typeSwitch(false)
	typeSwitch(man1)
	typeSwitch(woman1)
	typeSwitch(map1)
}

func SayandWork(person Person) {
	fmt.Println(person.Say("hello"))
	fmt.Println(person.Walk("hi"))
}

func genericityCall(s interface{}) {
	fmt.Println(s)
}

func checkUnSafe(s interface{}) bool {
	_s := s.(string)
	return strings.Contains(_s, "hehe")
}

func checkSafe(s interface{}) bool {
	_s, ok := s.(string)
	if ok {
		return strings.Contains(_s, "hehe")
	} else {
		return false
	}
}

func typeSwitch(s interface{}) {
	switch s.(type) {
	case int, uint:
		fmt.Println("这是一个数字")
	case string:
		fmt.Println("这是一个字符串")
	case bool:
		fmt.Println("这是一个布尔值")
	case Man:
		fmt.Println("这是一个Man")
	case Woman:
		fmt.Println("这是一个Woman")
	default:
		fmt.Println("这是一个特殊类型")
	}
}

// 定义一个 Person 接口
type Person interface {
	Say(s string) string
	Walk(s string) string
}

// 定义一个 Man 结构体
type Man struct {
	Name string
	Age  int
}

// Man 实现 Say 方法
func (m Man) Say(s string) string {
	return s + ", my name is " + m.Name
}

// Man 实现 Walk 方法，strconv.Itoa() 数字转字符串
func (m Man) Walk(s string) string {
	return "Age: " + strconv.Itoa(m.Age) + " and " + s
}

// 定义一个 Woman 结构体
type Woman struct {
	Name   string
	Age    int
	Height int
	Weight int
}

// Woman 实现 Say 方法
func (w Woman) Say(s string) string {
	return s + ", my name is " + w.Name + "。 my heigt is " + strconv.Itoa(w.Height) + " and my weight is " + strconv.Itoa(w.Weight)
}

// Woman 实现 Walk 方法，strconv.Itoa() 数字转字符串
func (w Woman) Walk(s string) string {
	return "Age: " + strconv.Itoa(w.Age) + " and " + s
}
