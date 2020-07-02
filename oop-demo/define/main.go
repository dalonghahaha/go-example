package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func (h *Human) Hello() {
	fmt.Println()
	fmt.Println("Hello,I am a Human, My name is ", h.Name)
}

func (h *Human) Imformation() {
	fmt.Println()
	fmt.Println("My age is ", h.Age)
	fmt.Println("My height is ", h.Height)
	fmt.Println("My weight is ", h.Weight)
}

func main() {
	//使用var实例化
	var human1 Human
	fmt.Println()
	fmt.Printf("%+v\n", human1)

	//使用字面量实例化
	human2 := Human{}
	fmt.Println()
	fmt.Printf("%+v\n", human2)
	human3 := Human{
		Name: "张三",
		Age:20,
		Height: 170,
		Weight: 120,
	}
	fmt.Println()
	fmt.Printf("%+v\n", human3)

	//使用new实例化，拿到的指针
	human4 := new(Human)
	fmt.Println()
	fmt.Printf("%+v\n", human4)

	human := Human{
		Name:   "李四",
		Height: 180,
		Weight: 150,
	}
	//使用属性
	fmt.Println()
	fmt.Println(human.Name)
	fmt.Println(human.Height)
	fmt.Println(human.Weight)
	//调用方法
	human.Hello()
	human.Imformation()
}
