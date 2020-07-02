package mans

import "fmt"

type Supperman struct {
	Name   string
	age    int
	height int
	weight int
}

func (h *Supperman) Hello() {
	fmt.Println()
	fmt.Println("Hello,I am a Supperman, My name is ", h.Name)
}

func (h *Supperman) imformation() {
	fmt.Println()
	fmt.Println("My age is ", h.age)
	fmt.Println("My height is ", h.height)
	fmt.Println("My weight is ", h.weight)
}

func (h *Supperman) CallSpiderman() {
	spiderman := Spiderman{
		id:   9527,
		age:  14,
		Name: "Peter Park",
	}
	fmt.Println()
	fmt.Println(spiderman.id)
	fmt.Println(spiderman.Name)
	spiderman.Hello()
	spiderman.imformation()
}
