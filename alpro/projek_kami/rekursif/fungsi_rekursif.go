package rekursif

import (
	"Project-Alpro/atribut"
	"fmt"
)

var Choice2 string

func ifelse(n int) {
	fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
	fmt.Scan(&Choice2)
	for Choice2 == "y" && n == 1 {
		atribut.ClearScreen()
		tampilanfibo()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 2 {
		atribut.ClearScreen()
		tampilanx()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 3 {
		atribut.ClearScreen()
		tampilanFUNGSI()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 4 {
		atribut.ClearScreen()
		tampilaneksponen()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 5 {
		atribut.ClearScreen()
		tampilanFaktorial()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 6 {
		atribut.ClearScreen()
		tampilankomper()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	if Choice2 == "r" {
		atribut.ClearScreen()
		MainMenu()
	} else {
		atribut.ClearScreen()
		fmt.Println("Kembali ke menu utama...")
		atribut.Loading()
	}
}
func tampilanfibo() {
	var a int
	fmt.Println("===============")
	fmt.Println("Fungsi FIBONACI")
	fmt.Println("===============")
	fmt.Print("Masukkan nilai untuk Fibonacci: ")
	fmt.Scan(&a)
	fmt.Print(a, " angka Fibonacci pertama adalah: ")
	for i := 1; i <= a; i++ {
		fmt.Print(fibo(i), " ")
	}
	fmt.Println()
	fmt.Println()
}
func fibo(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func tampilanx() {
	var num int
	atribut.ClearScreen()
	fmt.Println("==================================")
	fmt.Println("Fungsi (x - 3) + (x - 2) + (x - 1)")
	fmt.Println("==================================")
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&num)
	fmt.Print("Hasilnya adalah: ", x(num))
	fmt.Println()
	fmt.Println()

}
func x(x int) int {
	if x == 1 || x == 2 || x == 3 {
		return 1
	} else {
		return (x - 3) + (x - 2) + (x - 1)
	}
}
func tampilanFUNGSI() {
	var a int
	atribut.ClearScreen()
	fmt.Println("========================")
	fmt.Println("Fungsi n bilangan pertama")
	fmt.Println("========================")
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&a)
	fmt.Print(a, " angka pertama adalah: ")
	for i := 1; i <= a; i++ {
		fmt.Print(FUNGSI(i), " ")
	}
	fmt.Println()
	fmt.Println()
}
func FUNGSI(x int) int {
	if x == 1 {
		return 1
	} else {
		return FUNGSI(x-1) + x
	}
}
func tampilaneksponen() {
	var a int
	atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
	fmt.Println("==========")
	fmt.Println("Fungsi 2^n")
	fmt.Println("==========")
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&a)
	fmt.Println("Hasilnya adalah: ", eksponen(a))
	fmt.Println()
}
func eksponen(x int) int {
	if x == 1 {
		return 2
	} else {
		return eksponen(x-1) * 2
	}
}
func tampilanFaktorial() {
	var a int
	atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
	fmt.Println("==========")
	fmt.Println("Fungsi Faktorial")
	fmt.Println("==========")
	fmt.Print("Masukkan nilai faktorial: ")
	fmt.Scan(&a)
	fmt.Println("Hasilnya adalah: ", faktorial(a))
	fmt.Println()
}
func faktorial(a int) int {
	if a == 0 || a == 1 {
		return 1
	} else {
		return a * faktorial(a-1)
	}
}
func tampilankomper() {
	var input string
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&input)
	switch input {
	case "1":
		var n, r, hasil int
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
		fmt.Println("==========")
		fmt.Println("Implementasi Rekursif:Kombinasi")
		fmt.Println("==========")
		fmt.Print("Masukkan nilai n dan r (nCr): ")
		fmt.Scan(&n, &r)
		kombinasi(n, r, &hasil)
		fmt.Println("Hasilnya adalah: ", hasil)
		fmt.Println()
	case "2":
		var n, r, hasil int
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
		fmt.Println("==========")
		fmt.Println("Implementasi Rekursif:Permutasi")
		fmt.Println("==========")
		fmt.Print("Masukkan nilai n dan r (nPr): ")
		fmt.Scan(&n, &r)
		permutasi(n, r, &hasil)
		fmt.Println("Hasilnya adalah: ", hasil)
		fmt.Println()
	case "0":
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}
func kombinasi(n, r int, hasil *int) {
	if r > n {
		fmt.Println("Undefined, Nilai r tidak boleh lebih besar dari n.")
	} else {
		*hasil = faktorial(n) / (faktorial(r) * faktorial(n-r))
	}
}
func permutasi(n, r int, hasil *int) {
	if r > n {
		fmt.Println("Undefined, Nilai r tidak boleh lebih besar dari n.")
	} else {
		*hasil = faktorial(n) / faktorial(n-r)
	}
}
