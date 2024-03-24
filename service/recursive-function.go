package service

import "fmt"

func PrintRecursiveNumbers(number int64) {
	if number == 1 {
		fmt.Println(number, "")
	} else {
		fmt.Println(number)
		PrintRecursiveNumbers(number - 1)
	}

}
