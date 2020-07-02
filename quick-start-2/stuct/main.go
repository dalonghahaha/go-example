package main

import "fmt"

func main() {
	//实例化结构体
	var person1 person
	fmt.Println(person1)

	//访问成员
	person1.Name = "jack"
	person1.age = 10

	//调用方法
	fmt.Println(person1.Hello())
	fmt.Println(person1.printAge())

	//字面量实例化
	person2 := person{
		Name: "mike",
		age:  20,
	}
	person3 := person{
		Name: "mike",
		age:  20,
	}

	//结构体比较
	fmt.Println(person2 == person3)

	//结构体指针
	person4 := new(person) //person4是一个指针
	fmt.Println(person4)
	person4.Name = "Bill gates"
	fmt.Println(person4.Hello())

	//实例化包含结构体的结构体
	var shop1 shop
	fmt.Println(shop1)
	shop1.address = "Bei Jing"
	shop1.manager.Name = "xi da da"
	shop1.manager.age = 60
	fmt.Println(shop1.Greeting())

	//字面量实例化包含结构体的结构体
	shop2 := shop{
		address: "new york",
		manager: person{
			Name: "jobs",
			age:  50,
		},
	}
	fmt.Println(shop2.Greeting())

	//实例化结构体带匿名成员
	var employee1 employee
	employee1.Name = "jack" //访问匿名成员属性
	employee1.age = 10
	employee1.company = "Apple"
	fmt.Println(employee1.Hello()) //调用匿名成员方法
	fmt.Println(employee1.Work())
}

//定义结构体
type person struct {
	Name string
	age  int
}

func (p *person) Hello() string {
	return "my name is " + p.Name
}

func (p *person) printAge() string {
	return fmt.Sprintf("my age is %d", p.age)
}

//定义包含结构体的结构体
type shop struct {
	address string
	manager person
}

func (s *shop) Greeting() string {
	return fmt.Sprintf("my shop is in %s, my manager name is %s age is %d", s.address, s.manager.Name, s.manager.age)
}

//定义带匿名成员的结构体，模拟"继承"
type employee struct {
	person
	company string
}

func (e *employee) Work() string {
	return "my name is " + e.Name + ",I Work in " + e.company
}
