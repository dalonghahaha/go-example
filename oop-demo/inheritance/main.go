package main

import "fmt"

import "strconv"

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
	fmt.Println("My age is ", strconv.Itoa(h.Age))
	fmt.Println("My height is ", h.Height)
	fmt.Println("My weight is ", h.Weight)
}

type Man struct {
	Human
	Sex string
	Job string
}

func (h *Man) Work() {
	fmt.Println()
	fmt.Println("my Job is ", h.Job)
}

func (h *Man) Imformation() {
	fmt.Println()
	fmt.Println("My age is ", h.Age)
	fmt.Println("My height is ", h.Height)
	fmt.Println("My weight is ", h.Weight)
	fmt.Println("My sex is ", h.Sex)
	fmt.Println("My job is ", h.Job)
}

type SuperMan struct {
	Man
	Power int
}

func (s *SuperMan) Fly() {
	fmt.Println(s.Name, " is flying......")
}

type ET struct {
	Planet string
	Color  string
}

func (e *ET) BornImformation() {
	fmt.Println()
	fmt.Println("My planet is ", e.Planet)
	fmt.Println("My color is ", e.Color)
}

type Titan struct {
	ET
	SuperMan
	Age string
}

func main() {
	man := Man{
		//继承"类"无法使用"父类"的字面量
		//Name: "张三",
		Sex: "男",
	}
	fmt.Println()
	fmt.Printf("%+v\n", man)
	//使用"父类"属性
	man.Name = "比尔·盖茨"
	man.Height = 180
	man.Weight = 160
	//使用继承"类"属性
	man.Job = "码农"
	//调用"父类"方法
	man.Hello()
	//调继承"类"方法
	man.Work()
	//方法重写
	man.Imformation()

	//多级继承
	superMan := SuperMan{}
	superMan.Name = "克拉克·肯特"
	superMan.Job = "记者"
	superMan.Power = 10000
	superMan.Hello()
	superMan.Work()
	superMan.Fly()

	//多重继承
	thanos := Titan{}
	thanos.Name = "灭霸"
	thanos.Planet = "泰坦"
	thanos.Color = "紫"
	thanos.Age = "100岁"
	thanos.SuperMan.Age = 300
	thanos.Man.Age = 200
	thanos.Human.Age = 100
	//注意看一下这个输出
	fmt.Printf("%+v\n", thanos)
	thanos.Hello()
	thanos.Human.Hello()
	thanos.BornImformation()
	thanos.Fly()
	//注意看一下以下三个Imformation输出的结果
	thanos.Imformation()
	thanos.Human.Imformation()
	thanos.Man.Imformation()
}
