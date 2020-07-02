package mans

import "fmt"

type Spiderman struct {
	id   int
	age  int
	Name string
}

func (h *Spiderman) Hello() {
	fmt.Println()
	fmt.Println("Hello,I am a Spiderman, My name is ", h.Name)
}

func (h *Spiderman) imformation() {
	fmt.Println()
	fmt.Println("My id is ", h.id)
	fmt.Println("My age is ", h.age)
}
