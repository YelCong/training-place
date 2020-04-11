package main

import (
	"fmt"
	"github.com/training-place/operator/pkg"
)

func main() {
	division()
	compare()
	input()
	return
	pkg.T1()
	pkg.Ot1()
	//pkg.ot1()
	fmt.Println("self code ")
}

func division() {
	var num float32 = 4.0
	var res = 10 / num
	fmt.Printf("10/4 ...type = %T ,val = %v", res, res)
}

func compare() {
	var i int = 1
	var j float32 = 1.0
	var k int = 1
	var l float32 = 1.0
	if i == k {
		fmt.Println("\nif == ")
	} else if j == l {
		fmt.Println("else if ")
	} else {
		fmt.Println("else")
	}
}

func input() {
	var name string
	var age byte
	var sal int
	var i int
	fmt.Scanf("%d %d",&sal, &i)
	fmt.Println(sal,i)
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Println(name, age)
}
