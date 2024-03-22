package service

import "fmt"

func PrintMap() {
	person := map[string]string{
		"name":    "Tantowi",
		"address": "Ciledug",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
}