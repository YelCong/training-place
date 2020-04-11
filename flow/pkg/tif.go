package pkg

import "fmt"

func Run() {
	var a int = 1
	var b int = 2
	if a > b {
		fmt.Println("max:", a)
	} else {
		fmt.Println("max:", b)
	}
}
