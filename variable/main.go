package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/**
数据类型

基础数据类型
int8,int16... byte rune
uint8,uint16...
float32,float64
string
数组 & 结构体 struct

引用数据类型
指针  pointer
管道 chan
函数 function
切片 slice
接口 interface
map
*/
//var PublicNum int = 10
//var privateNum int = 11

func main() {
	var num1 int32  = 111
	fmt.Println(unsafe.Sizeof(num1))
	helloWorld()
	declare()
	declareFloat()
	declareChar()
	declareBool()
	declareString()
	typeChange()
	toString()
	declarePointer()
}

/**
变量= 变量名+变量值+数据类型，声明后数据类型是不能变的
*/
func declare() {
	var num1 int
	var num2 = 10.02
	num3 := 10.03
	var (
		num4 = 10.04
		num5 = 10.05
	)
	age, name := 30, "tom"
	var age2, name2 = 30, "jakerMa"
	fmt.Println(num1, num2, num3, num4, num5)
	fmt.Println(age, name)
	fmt.Println(age2, name2)
}

/**
浮点数= 符号位+指数位+尾数位  尾数位可能丢失，造成精度损失
int 32位机器 int32
uint 32位机器 默认4个字节
float 不受机器影响默认64位
*/
func declareFloat() {
	var num1 float64
	var num2 float32
	fmt.Printf("n1:%f,的数据类型%T 占用字节数量%d\n", num1, num1, unsafe.Sizeof(num1))
	fmt.Printf("n2:%f,的数据类型%T 占用字节数量%d\n", num2, num2, unsafe.Sizeof(num2))
}

/**
GO中的字符串是由单个字节连接起来的
统一的UTF-8编码 英文字母是1个字节,汉字是3个字节
*/
func declareChar() {
	var c1 byte = 'a'
	var c2 byte = 'b'
	fmt.Println("c1=", c1)
	fmt.Println("c1=", c2)
	fmt.Printf("c1=%c,c2=%c\nc1的ascii码值%d\n", c1, c2, c1)
}

func declareBool() {
	var b1 bool = true
	fmt.Println("b=", b1)
}

/**
GO中的字符串是不可变的
``反引号原样输出
*/
func declareString() {
	var str string = "hello Amoy"
	//str[0】 = "H"//error
	var str2 string = `"输出双引号"`
	str3 := "h1 " + //需要讲+号留在上一行的尾部
		"h2"
	fmt.Println(str, str2)
	fmt.Println(str3)
}

/**
类型转换
无显示转化(低精度到高精度)，不能自动转化
转化结果按溢出处理  int64(100000)  二进制  变成了 96
*/
func typeChange() {
	//低精度->高精度
	var i int = 100
	var n1 float32 = float32(i)
	fmt.Printf("i=%v,%T n1=%v,%T\n", i, i, n1, n1)
	//高精度->低精度
	var i2 int64 = 100000; //011 000 011 010 100 000
	var n2 int8 = int8(i2);
	fmt.Printf("i2=%v,%T n2=%v,%T\n", i2, i2, n2, n2)
	var num1 int = 99
	str := fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q \n", str, str)
}

func toString() {
	var num1 int = 999;
	var str = fmt.Sprintf("%d", num1)
	fmt.Printf("类型：%T 值:%q\n", str, str)
	str2 := strconv.FormatFloat(float64(num1), 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str2, str2)
}

func declarePointer() {
	var num int = 100
	var ptr *int = &num
	fmt.Printf("ptr = %T ptr = %v\n", ptr, ptr)
	fmt.Println("ptr所指向内存的值", *ptr)
	//var ptr2 *float32 = &num  这个地方是错的,指针所指向地址的数据类型为float 类型不匹配
}

func helloWorld() {
	fmt.Println("hello world magic golang","\t aaaa")
}
