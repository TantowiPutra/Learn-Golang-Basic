package service

import "fmt"

// Struct merupakan template yang digunakan untuk menggabungkan >= 0 tipe data untuk membuat tipe data baru

type Customer struct {
	Name, Address string
	Age           int
}

func PrintStruct() {
	customer := Customer{
		"Tantowi", "Pondok Lestari", 20,
	}

	fmt.Println(customer.Name, customer.Address, customer.Age)
}
