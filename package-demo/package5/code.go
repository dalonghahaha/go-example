package package5

import "fmt"

func init() {
	fmt.Println("package5 init")
	fmt.Println()
}

func Hello() {
	fmt.Println("call package5 Hello")
	fmt.Println()
}
