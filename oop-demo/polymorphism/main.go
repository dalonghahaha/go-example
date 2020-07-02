package main

import "fmt"

func main() {
	man := Man{
		Name:  "tony",
		Hobby: "football",
		Age:   20,
	}
	interview(man)
	woman := Woman{
		Name:   "lily",
		Age:    18,
		Height: 165,
		Weight: 100,
	}
	interview(woman)
	superMan := SuperMan{}
	interview(superMan)
}

func interview(p Person) {
	p.Hello()
	p.Introduce()
}

type Person interface {
	Hello()
	Introduce()
}

type Man struct {
	Name  string
	Age   int
	Hobby string
}

func (m Man) Hello() {
	fmt.Println()
	fmt.Println(fmt.Sprintf("Hello, I am a Man. my name is %s", m.Name))
}

func (m Man) Introduce() {
	fmt.Println()
	fmt.Println("This is My Information:")
	fmt.Println(fmt.Sprintf("I am %d years old.", m.Age))
	fmt.Println(fmt.Sprintf("I like %s.", m.Hobby))
	fmt.Println("I like woman.")
}

type Woman struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func (w Woman) Hello() {
	fmt.Println()
	fmt.Println(fmt.Sprintf("Hello, I am a Woman. my name is %s", w.Name))
}

func (w Woman) Introduce() {
	fmt.Println()
	fmt.Println("This is My Information:")
	fmt.Println(fmt.Sprintf("I am %d years old.", w.Age))
	fmt.Println(fmt.Sprintf("My height is %d.", w.Height))
	fmt.Println(fmt.Sprintf("My weight is %d.", w.Weight))
	fmt.Println("I like Man.")
}

type SuperMan struct {
}

func (s SuperMan) Hello() {
	fmt.Println()
	fmt.Println("Hello, I am a SuperMan.")
}

func (s SuperMan) Introduce() {
	fmt.Println()
	fmt.Println("I can fly.")
}
