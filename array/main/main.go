package main

import (
	"fmt"
	"math/rand"
	"time"
)

func tDefine() {
	var hens [2]float32
	hens[0] = 1
	hens[1] = 2
	//hens[2] = 3//invalid array index 2 (out of bounds for 2-element array)
	for i := 0; i < len(hens); i++ {
		fmt.Println(hens[i])
	}
}

func tDefine2() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])
}

func tDefine3() {
	var arr = [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr)
}

func tDefine4() {
	var arr = [...]string{"aaa", "bbb", "ccc"}
	//arr[4] = "eee"//invalid array index 4 (out of bounds for 3-element array)
	fmt.Printf("%T\n", arr)    //[3]string
	fmt.Printf("%T\n", arr[0]) //string
}

func tDefine5() {
	//var arr [...]string //use of [...] array outside of array literal
	//fmt.Printf("%T", arr)
}

func tDefine6() {
	var arr [3]string
	arr = [3]string{"aa", "bb", "cc"}
	fmt.Printf("%T\n", arr)
}

func tDefine7() {
	var arr = [...]string{1: "aaa", 2: "bbb"}
	fmt.Printf("%T\n", arr) //[3]string
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d,%v,%T\n", i, arr[i], arr[i]) //0,,string
	}
}

func traverse() {
	var arr [3]string = [3]string{"abc", "bcd", "def"}
	for index, val := range arr {
		fmt.Println(index, val)
	}
}

//默认是传值引用
func test(arr [3]int) {
	for index, val := range arr {
		fmt.Println(index, val)
		val--
		arr[index] = val
	}
}

//传地引用
func quote(pointArr *[3]int) {
	for index, val := range *pointArr {
		fmt.Println(index, val)
		val--
		(*pointArr)[index] = val
		fmt.Println(index, val)
	}
}

func test1() {
	var arr [3]int = [3]int{1, 2, 3}
	test(arr)
	for index, val := range arr {
		fmt.Println(index, val)
	}
	quote(&arr)
}

func test2() {
	var alphabet = [26]byte{}
	for index, _ := range alphabet {
		alphabet[index] = 'A' + byte(index) //注意这里index 要转化为byte类型
	}
	for i := 0; i < len(alphabet); i++ {
		fmt.Printf("%d:%c,", i, alphabet[i])
	}
	fmt.Println()
}

func randNum() {
	rand.Seed(time.Now().UnixNano()) //获取随机数，不加随机种子，编译完获取都是重复的
	fmt.Println(rand.Int31n(100))
	fmt.Println(rand.Int31n(100))
}

func define2D() {
	var arr = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Printf("%T,%v\n", arr, arr)
}

//数组的长度及数据类型是固定的，无法改变的
//可变的是切片
//使用数组的步骤 1.生命数组并开辟空间
//2.给数组的各个值赋值（默认值）
//3.使用数组
//4.数组默认是值传递
func main() {
	//tDefine()
	//tDefine2()
	//tDefine3()
	//tDefine4()
	//tDefine5()
	//tDefine6()
	//tDefine7()
	//traverse()
	//test1()
	//test2()
	//randNum()
	define2D()
}
