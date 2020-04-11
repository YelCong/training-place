package main

import (
	"fmt"
	"github.com/training-place/slice/pkg"
)

//切片是数组的一个引用，因此是引用类型  => 可动态变化的数组
func define() {
	var slice1 []int
	fmt.Printf("%T\n", slice1)
	var slice2 []int = make([]int, 5, 10)
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice2, cap(slice2), len(slice2))
	var slice3 = []int{1, 2, 3, 4}
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice3, cap(slice3), len(slice3)) //type:[]int,cap:4,len:4
}

func arrToSlice() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice = arr[1:4]
	fmt.Printf("%v", slice)                                              //[2 3 4]
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice, cap(slice), len(slice)) //type:[]int,cap:9,len:3
	//这里可以发现cap和截取的开始位置有关系
	var slice2 = arr[3:4]
	fmt.Printf("%v", slice2)                                                //[4]
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice2, cap(slice2), len(slice2)) //type:[]int,cap:7,len:1
}

func sliceToSlice() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var slice = arr[2:8]
	fmt.Printf("%v", slice) //[3 4 5 6 7 8]
	var slice2 = slice[3:4]
	fmt.Printf("%v", slice2) //[6]
}

//append 底层的机制就是新建个数组,更换引用的地址。原数组会被垃圾回收
func tAppend() {
	var slice = []int{1, 2, 3}
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice, cap(slice), len(slice))
	var slice2 = append(slice, 4, 5, 6)
	slice2[3] = 666
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice, slice, cap(slice), len(slice))
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice2, slice2, cap(slice2), len(slice2))
	//注意这里的写法  ...
	var slice3 = append(slice, slice...)
	fmt.Printf("type:%T,cap:%d,len:%d\n", slice3, cap(slice3), len(slice3))
}

func tCopy() {
	//从内存上理解引用传递
	var slice = []int{1, 2, 3}
	var slice2 = slice[:2]
	slice2[1] = 666
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice, slice, cap(slice), len(slice))
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice2, slice2, cap(slice2), len(slice2))
	//[1 666 3],type:[]int,cap:3,len:3
	//[1 666],type:[]int,cap:3,len:2

	//因为没有分配空间。。。
	var slice3 []int
	copy(slice3, slice)
	//[],type:[]int,cap:0,len:0
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice3, slice3, cap(slice3), len(slice3))

	//slice slice4的数据空间是独立的，不会互相影响
	var slice4 = make([]int, 5)
	copy(slice4, slice)
	fmt.Printf("%v,type:%T,cap:%d,len:%d\n", slice4, slice4, cap(slice4), len(slice4))
}

func main() {
	//define()
	//arrToSlice()
	//sliceToSlice()
	//tAppend()
	//tCopy
	//pkg.Split()
	pkg.Change()
}
