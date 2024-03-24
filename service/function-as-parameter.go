package service

import (
	"fmt"
	"strings"
)

func SayHelloWithFilter(name string, filter func(string) string, upper func(string) string) {
	fmt.Println("Hello, ", upper(filter(name)))
}

func SpamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func UpperFilter(name string) string {
	return strings.ToUpper(name)
}
