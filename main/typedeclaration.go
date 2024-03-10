package main

import "fmt"

type NoKTP string

func printTypeDeclaration() {
	var ExampleKTP NoKTP = "01234567890"

	fmt.Println(ExampleKTP)
}
