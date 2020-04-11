package pkg

import "fmt"

func RunFor() {
	for i := 0; i <= 3; i++ {
		fmt.Print(i)
	}

	var j int = 1
	for ; j <= 3; j++ {
		fmt.Print(j)
	}

	var str string = "abc fg"
	for index, val := range str {
		fmt.Printf("%d,%c\n", index, val)
	}
	utf8()
	while()
	tBreak()
}

func utf8() {
	fmt.Println("____________utf8___________")
	var str string = "golang 中国"
	str2 := []rune(str)
	for index, val := range str {
		fmt.Printf("%d,%c,", index, val)
	}
	fmt.Println()

	for i := 0; i < len(str); i++ {
		//普通的遍历是按照字节来的。这里会乱码
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}
}

func while() {
	fmt.Println("___________while______________")
	var judge bool = false
	for {
		fmt.Println("in")
		if !judge {
			break
		}
	}
}

func tBreak() {
	fmt.Println("________break___________")

	for {
		fmt.Println("label1")
	label1:
		for {
			for {
				break label1
			}
		}
		break
	}
}
