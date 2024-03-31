package service

import "fmt"

// type Customer struct {
// 	name string
// 	age  int
// }

func (customer Customer) SayHello(test string) {
	fmt.Println("Hello, My Name Is", customer.Name, test)
}

func PrintStructMethod() {
	rully := Customer{Name: "Rully"}
	rully.SayHello("Test")
}
