package service

import "fmt"

func PrintClosure() {
	counter := 0

	// 3 Nested Functions
	// Func 1
	increment := func() {
		counter++
		fmt.Println("Increment:", counter)

		name := "Tantowi"

		// Func 2
		printName := func() {
			fmt.Println(name)
			fmt.Println(counter)

			// Func 3
			printName2 := func() {
				fmt.Println(name)
				fmt.Println(counter)
			}

			printName2()
		}

		printName()
	}

	increment()
	increment()
	fmt.Println("Final Counter:", counter)
}
