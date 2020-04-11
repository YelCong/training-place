package pkg

import "fmt"

func DemoSwitch() {
	var i int
	fmt.Println("i=")
	fmt.Scanln(&i)
	switch i {
	case 1:
		fmt.Println(i)
	case 2, 3, 4, 5, 6, 7, 8:
		fmt.Println(i)
		fallthrough
	case 10:
		fmt.Println(i)
	default:
		fmt.Println("else")
	}
	fmt.Println("end...")
}
