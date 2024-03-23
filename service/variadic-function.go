package service

import "strconv"

// Varargs / variadic function hanya bisa diletakan parameternya dipaling akhir (hanya  bisa satu)
// Varargs tidak wajib melempar data ke parameter
func SumAll(name string, numbers ...int) string {
	total := 0
	for _, number := range numbers {
		total += number
	}

	text := "Total " + name + " Is: " + strconv.Itoa(total)

	return text
}
