package service

import (
	"fmt"
	"strconv"
)

func PrintConvertDataType() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	var nilaifloat32 float64 = 3.163123
	fmt.Println(strconv.FormatFloat(nilaifloat32, 'f', -1, 32))

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "Tantowi"
	var e byte = name[0]

	convString := string(e)
	fmt.Println(convString)
}
