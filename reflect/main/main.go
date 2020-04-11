package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 10
	reflectTest01(a)
}

//专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)//rType= int


	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v,type of rVal %T\n", rVal,rVal)//rVal=10,type of rVal reflect.Value

	rVal.SetInt(4444)

	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}
