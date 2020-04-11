package main

import (
	"fmt"
	"sort"
)

var list map[string]map[string]string

func define() {
	//map[key_type]val_type
	var m map[string]string
	fmt.Printf("%T\n", m)
	m = make(map[string]string, 2)
	m["idx0"] = "val0"
	m["idx1"] = "idx1"
	m["idx2"] = "idx2"
	m["idx3"] = "idx3"
	fmt.Printf("%v\n", m)
}

func define2() {
	list = make(map[string]map[string]string, 3)
	list["no1"] = make(map[string]string)
	list["no1"]["name"] = "张三"
	list["no1"]["age"] = "10"
	list["no2"] = map[string]string{
		"name": "李四",
		"age":  "11",
	}
	list["no0"] = map[string]string{
		"name": "王五",
	}
	list["no3"] = map[string]string{
		"name": "王五3",
	}
	list["no2"] = map[string]string{
		"name": "王五2",
	}
	fmt.Printf("%v\n", list)
	//delete(list, "no3")
	delete(list, "no4") //不处理也不报错
	return
	fmt.Printf("%v\n", list)
	fmt.Printf("%v\n", list["no2"])
	val, findRes := list["no2"]
	fmt.Println(val, findRes) //map[age:11 name:李四] true
	val2, findRes2 := list["no4"]
	fmt.Println(val2, findRes2) //map[] false
	//没有提供全部处理的方法，但是可以用下面这个方式，原先的就会被gc
	//list = make(map[string]map[string]string)
	//fmt.Printf("%v\n", list)

}

func forRange() {
	for idx, val := range list {
		fmt.Println(idx, val)
		//no1 map[age:10 name:张三]
		//no2 map[age:11 name:李四]
		for idx1, val1 := range val {
			fmt.Println(idx1, val1)
		}
	}
}

func sliceSort() {
	var keys []string
	for k, _ := range list {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println(k, list[k])
	}
}

func modify(m map[string]map[string]string) {
	m["no1"]["name"] += "aaa"
}

func run() {
	//define()
	define2()
	fmt.Println("______________________")
	fmt.Println(list)
	//forRange()
	//sliceSort()
	modify(list)
	fmt.Println(list)
}

func main() {
	run()
	//pkg.SliceOfMap()
}
