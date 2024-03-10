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
		{4, "Convert Data Type"},
		{5, "Type Declaration"},
		{6, "Exit"},
	}

	for {
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
		case 4:
			printConvertDataType()
			fmt.Println()
		case 5:
			printTypeDeclaration()
			fmt.Println()
		}

		if option == len(menuList) {
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
