package main

import (
	"fmt"
	"github.com/training-place/function/pkg"
)

var a = test()

//b := "err"//=>  var a string  a =  "err" 赋值语句不能在函数体外
var b string = "error"

func test() int {
	fmt.Println("variable declare")
	return 1
}

func init() {
	fmt.Println("init ...")
}

//可变参数
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func AddUpper() func(int) int {
	var n = 10
	fmt.Println("var")
	return func(i int) int {
		n = n + i
		return n
	}
}

func tDefer() {
	fmt.Println("tDefer..")
	defer fmt.Println("000000", b)
}

func main() {
	pkg.Run()
	return
	tDefer()
	return
	f1 := AddUpper()
	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))
	return
	fmt.Println("main....")
	sum1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("sum=", sum1)
	var f = sum
	//f type:func(int, ...int) int,sum type:func(int, ...int) int
	fmt.Printf("f type:%T,sum type:%T\n", f, sum)
}
