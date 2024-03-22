package service

import "fmt"

type NoKTP string

func PrintTypeDeclaration() {
	var ExampleKTP NoKTP = "01234567890"

	fmt.Println(ExampleKTP)
}
