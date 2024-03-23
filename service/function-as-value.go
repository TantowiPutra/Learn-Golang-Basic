package service

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func PrintFunctionAsValue() {
	// Menyimpan Function Sebagai Value
	goodbye := getGoodBye
	fmt.Println(goodbye("Sir"))
}
