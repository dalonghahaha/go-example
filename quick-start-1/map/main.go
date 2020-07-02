package main

import "fmt"

func main() {
	// 声明空map
	var m map[int]string
	fmt.Println(m)
	//m[1] = `aa`  //编译会报错，不能向未初始化的map中直接赋值

	// 声明并初始化，初始化使用{} 或 make 函数（创建类型并分配空间）
	var m1 = map[string]int{}
	var m2 = make(map[string]int)
	var m3 = map[string]string{"a": "aaa", "b": "bbb"}
	m1["a"] = 11
	m2["b"] = 22
	m3["c"] = "ccc"
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	//直接取值
	fmt.Println(m3["c"])

	// 查找map中的元素
	v, ok := m3["b"]
	if ok {
		fmt.Println(ok)
		fmt.Println("m3中b的值为：", v)
	}
	if v, ok := m3["b"]; ok {
		fmt.Println("m3中b的值为：", v)
	}

	// 删除map中的值
	delete(m3, "a")
	fmt.Println(m3)

	//多维map
	m4 := make(map[string][5]int)
	m4["a"] = [5]int{1, 2, 3, 4, 5}
	m4["b"] = [5]int{11, 22, 33}
	fmt.Println(m4)
}
