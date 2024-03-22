package service

import (
	"fmt"
	"strconv"
)

func PrintOperation() {
	a := 10
	b := 10

	c := a + b

	fmt.Println(strconv.Itoa(a) + " + " + strconv.Itoa(b))
	fmt.Println(c)
}
