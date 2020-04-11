package polymorphism

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

func (phone Phone) Start() {
	fmt.Printf("%v phone strat...\n", phone.Name)
}

func (phone Phone) Stop() {
	fmt.Printf("%v phone Stop...\n", phone.Name)
}

func (phone Phone) Call() {
	fmt.Printf("%v phone call...\n", phone.Name)
}

func (camera Camera) Start() {
	fmt.Printf("%v camera strat...\n", camera.Name)
}

func (camera Camera) Stop() {
	fmt.Printf("%v camera Stop...\n", camera.Name)
}

//Camera 和　Phone 都可以实现Start Stop 就是多态的一种体现