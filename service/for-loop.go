package service

import "fmt"

func PrintForLoop() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	// for dengan dua statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("For dengan dua Statement Perulangan Ke ", counter)
	}
}
