package pkg

import (
	"errors"
	"fmt"
)

func tErr() {
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("endding", res)
}

func tCatch() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("catch error", err)
		}
	}()
	tErr()
}

func customErr() (err error) {
	return errors.New("custom err")
}

//java -> try catch finally
//golang -> defer recover panic
func Run() {
	//tErr()
	//tCatch()
	err := customErr()
	if err != nil {
		panic(err) //执行到这里下面的东西就不会再执行了
	}
	fmt.Println("main")
}
