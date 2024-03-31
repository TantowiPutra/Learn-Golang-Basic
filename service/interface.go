package service

import "fmt"

type HasName interface {
	GetName() string
}

type Example interface {
	name() string
}

type HasFunction interface {
	Sample1() string
	Sample2() string
	Sample3() string
}

type Person struct {
	Name string
}

func SayHello2(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func PrintSample(hasFunction HasFunction) HasFunction {
	fmt.Println(hasFunction.Sample1(), hasFunction.Sample2(), hasFunction.Sample3())

	return hasFunction
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) Sample1() string {
	return "Sample 1"
}

func (person Person) Sample2() string {
	return "Sample 2"
}

func (person Person) Sample3() string {
	return "Sample 3"
}
