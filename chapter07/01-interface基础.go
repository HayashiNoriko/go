package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia, I can call you!")
}

type IPhone struct{}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main1() {
	var phone Phone
	phone = NokiaPhone{}
	phone.call()

	phone = IPhone{}
	phone.call()

}
