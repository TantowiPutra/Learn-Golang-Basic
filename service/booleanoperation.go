package service

import "fmt"

func PrintBooleanOperation() {
	var nilaiAkhir = 90
	var absensi = 80

	lulusNilaiAkhir := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80

	if lulusNilaiAkhir && lulusAbsensi {
		fmt.Println("Selamat Anda Lulus Mata Kuliah Ini!")
	}
}
