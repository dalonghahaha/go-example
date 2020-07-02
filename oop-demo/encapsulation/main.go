package main

import (
	"fmt"

	"encapsulation/mans"
)

func main() {
	supperMan := mans.Supperman{
		Name: "Clark Kent",
	}
	fmt.Printf("%+v\n", supperMan.Name)
	//fmt.Printf("%+v\n", supperMan.age)
	supperMan.Hello()
	//supperMan.imformation()
	supperMan.CallSpiderman()
}
