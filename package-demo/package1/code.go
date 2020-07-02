package package1

import "fmt"

const ENV = "mac"

var Version = "1.0"

func init() {
	fmt.Println("package1 init")
	fmt.Println()
}

func Hello() {
	fmt.Println("call package1 Hello")
	fmt.Println()
}
