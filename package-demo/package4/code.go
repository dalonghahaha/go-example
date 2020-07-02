package package4

import (
	"fmt"

	"package-demo/package5"
	"package-demo/package6"
)

func init() {
	fmt.Println("package4 init")
	fmt.Println()
}

func Hello() {
	fmt.Println("call package4 Hello")
	fmt.Println()
	package5.Hello()
	package6.Hello()
}
