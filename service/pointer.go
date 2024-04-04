package service

import "fmt"

type Address struct {
	City, Province, Country string
}

func PrintPointer() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// Dengan Pointer, kita menyimpan alamat memori dari deklarasi variable address1, 
	// dengand demikian, perubahan pada variable address2 akan berpengaruh pada address1
	var address2 *Address = &address1
	fmt.Println(*address2)
}
