package main

import "fmt"

func main() {
	// 声明和使用切片切片
	var sl = []int{2, 3, 5, 7, 11}
	fmt.Println(sl)
	fmt.Println(sl[1])

	//追加数据
	sl = append(sl, 1)
	fmt.Println(sl)
	sl = append(sl, 1, 2, 3)
	fmt.Println(sl)
	// = append(sl, "a") //编译会报错，不能追加不同类型的数据
	var sla = []int{4, 5, 6}
	sl = append(sl, sla...)
	fmt.Println(sl)

	//make初始化：make([]类型, 长度，容量)
	var sl2 = make([]int, 5)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))
	fmt.Println(sl2[1])
	var sl3 = make([]int, 5, 10)
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	//数组的切片引用,冒号:左边为起始位（包含起始位数据），右边为结束位（不包含结束位数据）；不填则默认为头或尾
	var arr = [5]int{1, 2, 3, 4, 5}
	var sl4 = arr[0:2]
	var sl5 = arr[3:]
	var sl6 = arr[:5]
	fmt.Println(sl4)
	fmt.Println(sl5)
	fmt.Println(sl6)
}
