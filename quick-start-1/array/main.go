package main

import "fmt"

func main() {
	// 声明后赋值：var <数组名称> [<数组长度>]<数组元素>
	var arr [2]int
	fmt.Println(arr[1]) //取值
	arr[0] = 1          //赋值
	//arr[2] = 3 // 编译会报错，数组越界

	// 声明并赋值：var <数组名称> = [<数组长度>]<数组元素>{元素1,元素2,...}
	var intArr = [2]int{1, 2}
	var strArr = [3]string{`aa`, `bb`, `cc`}
	fmt.Println(intArr)
	fmt.Println(strArr)

	// 声明时不设定大小，赋值后语言本身会计算数组大小：var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{元素1,元素2,...}
	var arr1 = [...]int{1, 2}
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)

	//多维数组
	var multi_arr1 = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(multi_arr1[1][2])
	// fmt.Println(multi_arr[3][0])// 编译报错，数组越界

	//多维数组部分初始化，没有初始化的值为0
	var multi_arr2 = [3][4]int{{1, 2, 3}, {5, 6, 7, 8}, {9, 10}}
	fmt.Println(multi_arr2[2][3])
}
