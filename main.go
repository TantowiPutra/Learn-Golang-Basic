package main

import (
	"fmt"
	"idx/service"
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
		{6, "Print Addition"},
		{7, "Exit"},
	}

	for {
		printMenu(&menuList)

		fmt.Printf("Select Option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			// ! Hello World
			service.PrintHelloWorld()
			fmt.Println()
		case 2:
			// ! Basic Numbers
			service.PrintBasicNumbers()
			fmt.Println()
		case 3:
			// ! Boolean
			service.PrintBoolean()
			fmt.Println()
		case 4:
			// ! Convert Data Type
			service.PrintConvertDataType()
			fmt.Println()
		case 5:
			// ! Type Declaration
			service.PrintTypeDeclaration()
			fmt.Println()
		case 6:
			service.PrintOperation()
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
