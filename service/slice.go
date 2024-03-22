package service

import "fmt"

func PrintSlice() {

	// [...] -> Deklarasi Array yang tidak diketahui sizenya
	// []	 -> Deklarasi Slice
	var months = []string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)

	// Len = (end - start)
	fmt.Println(len(slice1))

	// Cap = Max Capacity - Start
	fmt.Println(cap(slice1))

	// Mengubah original slice, maka slice yang mereferensikan ke slice tersebut juga berubah (pass by reference)
	// Begitu juga sebaliknya, jika yang diubah slice yang referensi, maka slice asli juga akan berubah
	months[5] = "Diubah"
	fmt.Println(months)
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)
	fmt.Println(slice1)

	fmt.Println()
	fmt.Println()

	// 	Append untuk menambahkan data ke slice, tapi jadi pass by value (original slice tidak berubah)
	//	Dalam append, kita perlu menambahkan ... setelah slice yang ingin kita tambahkan ke slice utama. Itu karena append menerima argumen variadic, yang berarti dapat menerima banyak argumen. Dengan menambahkan ..., kita memecah slice kedua menjadi elemen-elemen individu yang disisipkan ke slice pertama.

	//  Variadic Function: Dapat menerima argmumen tanpa batas
	slice2 := append(months, []string{
		"Senin Lagi",
		"Selasa Lagi",
	}...)

	fmt.Println(slice1)
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
