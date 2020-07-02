package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"

	//获取字符串长度
	fmt.Println(len(s))

	//通过索引获取byte
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	//索引不能越界
	//c := s[len(s)] // panic: index out of range
	//fmt.Println(c)

	//字符串切片
	fmt.Println(s[0:5])
	fmt.Println("goodbye" + s[5:])

	//不能直接修改字符串的值
	//s[0] = 'L' // compile error: cannot assign to s[0]
	//fmt.Println(s)

	//中文字符串长度的区别
	str := "hello 世界"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
