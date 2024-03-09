package main

import (
	"fmt"
)

type MenuItem struct {
	idx  int    `int`
	menu string `string`
}

// Basic Output to Console & Code Compilation using "go build" command
func main() {
	var option int

	menuList := []MenuItem{
		{1, "Hello, World!"},
		{2, "Basic Numbers"},
		{3, "Boolean"},
		{4, "Exit"},
	}

	for true {
		printMenu(&menuList)

		fmt.Printf("Select Option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			printHelloWorld()
			fmt.Println()
		case 2:
			printBasicNumbers()
			fmt.Println()
		case 3:
			printBoolean()
			fmt.Println()
		}

		if option == 4 {
			break
		}
	}
}

func printMenu(menuList *[]MenuItem) {
	fmt.Println("Welcome Abord! Please Choose a Menu To Proceed: ")
	fmt.Println("================================================================")

	for _, item := range *menuList {
		fmt.Printf("%d. %s\n", item.idx, item.menu)
	}
}
