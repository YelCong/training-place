package pkg

import "fmt"

var virus []map[string]string

func define() {
	virus = make([]map[string]string, 2)
	virus[0] = map[string]string{
		"name": "H1N1",
		"desc": "H1N1是一种RNA病毒，属于正黏液病毒科",
	}
	virus[1] = make(map[string]string)
	virus[1]["name"] = "H7N9"
	virus[1]["desc"] = "H7N9型禽流感是一种新型禽流感"
}

func tAppend() {
	var newVirus = map[string]string{
		"name": "health",
		"desc": "有毒",
	}
	virus = append(virus, newVirus)
}

func SliceOfMap() {
	define()
	fmt.Println(virus)
	tAppend()
	fmt.Println(virus)
}
