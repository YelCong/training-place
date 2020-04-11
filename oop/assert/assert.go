package assert

/**
类型断言
*/
import (
	"fmt"
	"github.com/training-place/oop/polymorphism"
)

func Run2() {
	var x interface{}
	var a float32 = 1.1
	x = a
	y, ok := x.(float32)
	if ok {
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}
	fmt.Printf("%T,%v\n", y, y)
}

func Run1() {
	var usbArr [3]polymorphism.Usb
	usbArr = [3]polymorphism.Usb{
		polymorphism.Phone{"p1"},
		polymorphism.Phone{"p2"},
		polymorphism.Camera{"c1"},
	}
	fmt.Println(usbArr)
	for i := 0; i < len(usbArr); i++ {
		usbArr[i].Start()
		if terminal, ok := usbArr[i].(polymorphism.Phone); ok {
			terminal.Call()
		}
		usbArr[i].Stop()
	}
}

func Run() {
	typeJudge()
}

//传入的参数个数不确定,类型也不确定
func _typeJudge(items ...interface{}) {
	for _, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("%T,%v\n", x, x)
		case int32:
			fmt.Printf("%T,%v\n", x, x)
		case int64:
			fmt.Printf("%T,%v\n", x, x)
		case *polymorphism.Phone:
			fmt.Printf("%T,%v\n", x, x)
		default:
			fmt.Printf("%v type unknown\n", x)
		}
	}
}

func typeJudge() {
	var a int32 = 1
	var b int64 = 2
	var c float32 = 1
	var d string = "望海路2-2"
	var p1 *polymorphism.Phone = &polymorphism.Phone{"p1"}
	//t := []interface{}{1, 2, 3}
	_typeJudge(a, b, c, d, p1)
}

func  a(){
}
