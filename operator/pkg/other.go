package pkg

import (
	"fmt"
)

var PublicNum int = 10
var privateNum int = 11

func ot1() {
	fmt.Println("import error, private function")
}


func Ot1(){
	fmt.Println("import success ,public function other->Ot1")
}