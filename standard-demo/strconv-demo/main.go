package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// bool操作
	exampleBool()
	// float操作
	exampleFloat()
	// int操作
	exampleInt()
	// uint操作
	exampleUint()
}

func exampleBool() {
	// bool转string
	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)
	// string转bool
	bl, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", bl, bl)
	// 添加string类型的bool值
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}

func exampleFloat() {
	v := 3.1415926535
	// float32转string
	// fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）
	// prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f
	s32 := strconv.FormatFloat(float64(v), 'f', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
	// float64转string
	s64 := strconv.FormatFloat(v, 'f', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)
	// string转float32
	if s, err := strconv.ParseFloat(s32, 32); err == nil {
		s1 := float32(s)
		fmt.Printf("%T, %v\n", s1, s1)
	}
	// string转float64
	if s, err := strconv.ParseFloat(s64, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// 添加string类型的float32值
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, v, 'f', -1, 32)
	fmt.Println(string(b32))
	// 添加string类型的float64值
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, v, 'f', -1, 64)
	fmt.Println(string(b64))
}

func exampleInt() {
	v := 12
	// int转string
	s := strconv.Itoa(v)
	fmt.Printf("%T, %v\n", s, s)
	// string转int
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i, i)
	// int64转十进制string
	// base 必须在2到36之间
	s = strconv.FormatInt(int64(v), 10)
	fmt.Printf("%T, %v\n", s, s)
	// 十进制string转int64
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i64, i64)
	// 添加string类型的int十进制值
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	// 添加string类型的int十六进制值
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}

func exampleUint() {
	v := uint64(42)
	// uint64转十进制string
	s := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s, s)
	// 十进制string转uint64
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", u, u)
	// 添加string类型的uint十进制值
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))
}
