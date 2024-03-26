package service

import "fmt"

// Defer
/*
	Function yang dapat dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	Penggunaannya biasanya dipair dengan recover, karena dapat digunakan untuk menangkap error ketika
	sebuah function telah selesai dieksekusi
*/

func Logging() {
	fmt.Println("Berhasil membuat log")
}

func RunApplication(value int) {
	defer Logging()
	fmt.Println("Run Application", 1)
	result := 10 / value
	fmt.Println("Result:", result)
}

// Panic
/*
	Panic: Digunakan untuk memberhentikan program
	Digunakan ketika terjadi error saat program berjalan
	saat panic function dipanggil, program akan berhenti, namun defer tetap dieksekusi
*/

func EndApp() {
	message := recover()
	if message != nil && message != "" {
		fmt.Println("Masuk Kondisi")
		fmt.Println(message)
	}
	fmt.Println("Program Closed")
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi Berhasil Berjalan")
}

/*
	Recover: Function yang dapat digunakan untuk menangkap panic
	Dengan fungsi recover, panic akan berhenti sehingga program tetap dapat berjalan
*/

func RunApp2(error bool) {
	defer EndApp()
	if error {
		panic("Ditemukan Error pada Aplikasi")
	}
}
