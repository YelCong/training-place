package pkg

import "fmt"

func Split() {
	var str string = "hello world"
	fmt.Println(str)
	var part = str[:len(str)-1]
	fmt.Println(part)
}

func Change() {
	var str string = "hello 北京"
	var arr1 = []rune(str)
	fmt.Println(arr1)
	arr1[0] = 'H'
	arr1[len(arr1)-1] = '境' //""表示的是字符串
	fmt.Println(arr1)
	str = string(arr1)
	fmt.Println(str)
}
