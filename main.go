package main

import (
	"errors"
	"fmt"
	"idx/service"
	"strconv"
)

type Blacklist func(string) bool

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
		{19, "Variadic Function (Varargs)"},
		{20, "Function as Value"},
		{21, "Function as Parameter"},
		{22, "Anonymous Function"},
		{23, "Recursive Function"},
		{24, "Closure"},
		{25, "Defer, Panic, and Recover"},
		{26, "Struct"},
		{27, "Struct Function"},
		{28, "Interface"},
		{29, "Nil"},
		{30, "Test Code"},
		{31, "Error Interface"},
		{32, "Pointer"},
		{33, "Exit"},
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

			fmt.Println()
		case 19:
			response := service.SumAll("Apple", 1, 2, 3, 4, 5, 6, 7)
			fmt.Println(response)
			fmt.Println()

			// Or We Could Pass Slice To A Variadic Function using the notation (...) to indicate that, the slice should be treated
			// Individially instead as a group (slice)

			numberSlice := []int{1, 3, 5}
			fmt.Println(service.SumAll("Melons", numberSlice...))
			fmt.Println()
		case 20:
			service.PrintFunctionAsValue()
			fmt.Println()
		case 21:
			service.SayHelloWithFilter("Anjing", service.SpamFilter, service.UpperFilter)
			service.SayHelloWithFilter("test nama", service.SpamFilter, service.UpperFilter)
		case 22:
			blacklist := func(name string) bool {
				return name == "anjing"
			}

			service.RegisterUser("anjing", blacklist)
			service.RegisterUser("test nama", blacklist)
			fmt.Println()
		case 23:
			service.PrintRecursiveNumbers(120)
			fmt.Println()
		case 24:
			service.PrintClosure()
			fmt.Println()
		case 25:
			// 1. Defer
			// Simulasi Error, pembagian angka dengan, karena error maka masuk state panic
			// service.RunApplication(0)

			// 2. Panic
			// Simulasi Error Panic
			// service.RunApp(true)

			// 3. Recover
			// Simulasi Recover dari Error menggunakan fungsi panic()
			service.RunApp2(true)

			fmt.Println()
		case 26:
			// Struct
			service.PrintStruct()
			fmt.Println()
		case 27:
			// Struct Method
			service.PrintStructMethod()
			fmt.Println()
		case 28:
			// Interface
			person := service.Person{Name: "Tantowi"}
			service.SayHello2(person)
			fmt.Println(service.PrintSample(person))
			fmt.Println()
		case 29:
			mapNil := service.NewMap("")
			mapVal := service.NewMap("Tantowi")

			fmt.Println(mapNil)

			if mapNil == nil {
				fmt.Println("Data Kosong")
			}

			fmt.Println(mapVal)
			fmt.Println()
		case 30:
			number := 1
			numberString := "1"
			parsedVal, err := strconv.Atoi(numberString)

			parsedStringVal := strconv.Itoa(parsedVal)

			if err == nil {
				if number == parsedVal {
					fmt.Println("Sama")
				} else {
					fmt.Println(number, numberString)
				}
			} else {
				fmt.Println("Error", err)
			}

			fmt.Println(parsedStringVal)
			fmt.Println()
		case 31:
			res, err := service.Division(1, 0)
			fmt.Println(res, err)

			// Contoh Error
			exampleError := errors.New("ups error")
			fmt.Println(exampleError)
			fmt.Println()
		case 32:
			service.PrintPointer()
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
