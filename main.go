package main

import (
	"fmt"
	"idx/service"
)

type MenuItem struct {
	idx  int
	menu string
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
		{7, "Print Operator Comparison"},
		{8, "Print Boolean Operation"},
		{9, "Print Array Contents"},
		{10, "Slice Data Type"},
		{11, "Map Data Type"},
		{12, "If Expression"},
		{13, "Switch Expression"},
		{14, "For Loop"},
		{15, "Function Parameter"},
		{16, "Function Return Value"},
		{17, "Function Return Multiple Value"},
		{18, "Function Named Return Value"},
		{19, "Exit"},
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
		case 7:
			service.PrintOperatorComparison()
			fmt.Println()
		case 8:
			service.PrintBooleanOperation()
			fmt.Println()
		case 9:
			service.PrintArray()
			fmt.Println()
		case 10:
			service.PrintSlice()
			fmt.Println()
		case 11:
			service.PrintMap()
			fmt.Println()
		case 12:
			service.PrintIf()
			fmt.Println()
		case 13:
			service.PrintSwitch()
			fmt.Println()
		case 14:
			service.PrintForLoop()
			fmt.Println()
		case 15:
			service.SayHelloTo("Tantowi", "123")
			fmt.Println()
		case 16:
			fmt.Println(service.PrintFunctionReturnValue("Tantowi"))
			fmt.Println()
		case 17:
			firstString, firstSlice, firstMap := service.PrintFunctionReturnMultipleValue()
			fmt.Println(firstString)

			for _, item := range firstSlice {
				fmt.Println(item)
			}

			for _, item := range firstMap {
				fmt.Println(item)
			}
		case 18:
			// Function Named Return Value digunakan untuk meningkatkan readability function
			firstString, firstSlice, firstMap := service.PrintFunctionNamedReturnValue()
			fmt.Println(firstString)

			for _, item := range firstSlice {
				fmt.Println(item)
			}

			for _, item := range firstMap {
				fmt.Println(item)
			}
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
