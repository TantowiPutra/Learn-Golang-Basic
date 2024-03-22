package service

import "fmt"

func PrintArray() {
	names := [3]string{
		"Tantowi",
		"Putra",
		"Agung",
	}

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}
