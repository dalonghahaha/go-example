package main

import (
	"fmt"

	"github.com/dalonghahaha/avenger/tools/uuid"

	"package-demo/package1"
	"package-demo/package1/sub"
	. "package-demo/package2"
	_ "package-demo/package3"
	p4 "package-demo/package4"
	test "package-demo/package7"

	"julive/tools/random"
)

func init() {
	fmt.Println("main init")
	fmt.Println()
}

func main() {
	fmt.Println("begin !")
	fmt.Println()
	//使用常量
	fmt.Println(package1.ENV)
	fmt.Println()
	//使用全局变量
	fmt.Println(package1.Version)
	fmt.Println()
	//使用函数
	package1.Hello()
	sub.Hello()
	//不带限定名的调用函数
	Func()
	//使用别名
	p4.Hello()
	test.Hello()
	//使用外部包
	fmt.Println(uuid.GenerateV1())
	fmt.Println()
	//使用私有外部包
	fmt.Println(random.CharsPlus(8))
	fmt.Println()
	fmt.Println("end !")
}
