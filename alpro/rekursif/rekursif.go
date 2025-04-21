package rekursif

import (
	"fmt"
	"os"
	"os/exec"
)

var pilihan string
var Choice2 string

func MainMenu() {
	clearScreen()
	Submenu()
	switch pilihan {
	case "1":
		clearScreen()
		tampilanfibo()
		fmt.Println("Tetap di Fibonacci(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
		for Choice2 == "y" {
			tampilanfibo()
			fmt.Println("Tetap di Fibonacci(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
			fmt.Scan(&Choice2)
		}
		if Choice2 == "r" {
			clearScreen()
			MainMenu()
		} else {
			clearScreen()
			fmt.Println("Kembali ke menu utama...")
			// Kode untuk kembali ke menu utama
			// Misalnya, panggil fungsi MainMenu() dari package lain atau lakukan tindakan lain sesuai kebutuhan
		}
	case "2":
		var num int
		clearScreen()
		fmt.Println("==================================")
		fmt.Println("Fungsi (x - 3) + (x - 2) + (x - 1)")
		fmt.Println("==================================")
		fmt.Print("Masukkan nilai x: ")
		fmt.Scan(&num)
		fmt.Print("Hasilnya adalah: ", x(num))
		fmt.Println()
	case "3":
		var a int
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
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
	case "4":
		var a int
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
		fmt.Println("==========")
		fmt.Println("Fungsi 2^n")
		fmt.Println("==========")
		fmt.Print("Masukkan nilai n: ")
		fmt.Scan(&a)
		fmt.Println("Hasilnya adalah: ", eksponen(a))
		fmt.Println()
	case "5":
		var a int
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
		fmt.Println("==========")
		fmt.Println("Fungsi Faktorial")
		fmt.Println("==========")
		fmt.Print("Masukkan nilai faktorial: ")
		fmt.Scan(&a)
		fmt.Println("Hasilnya adalah: ", faktorial(a))
		fmt.Println()
	case "6":
		var input string
		clearScreen()
		Submenu1()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)
		switch input {
		case "1":
			var n, r, hasil int
			clearScreen() // Membersihkan layar sebelum menampilkan hasil
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
			clearScreen() // Membersihkan layar sebelum menampilkan hasil
			fmt.Println("==========")
			fmt.Println("Implementasi Rekursif:Permutasi")
			fmt.Println("==========")
			fmt.Print("Masukkan nilai n dan r (nPr): ")
			fmt.Scan(&n, &r)
			permutasi(n, r, &hasil)
			fmt.Println("Hasilnya adalah: ", hasil)
			fmt.Println()
		case "0":
			fmt.Println("Terima kasih telah menggunakan kalkulator ini.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Submenu() {
	fmt.Println(`
=================
Jenis rekursif mana?
=================
1. Fibonacci
2. (x - 3) + (x - 2) + (x - 1)
3. Fungsi nilai ke-n beda tingkat dua
4. Eksponen 2*n
5. Faktorial
6. Implementasi faktorial
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}
func Submenu1() {
	fmt.Println(`
=================
Implementasi Rekursif!
=================
1. kombinasi
2. permutasi
	`)
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
func x(x int) int {
	if x == 1 || x == 2 || x == 3 {
		return 1
	} else {
		return (x - 3) + (x - 2) + (x - 1)
	}
}
func FUNGSI(x int) int {
	if x == 1 {
		return 1
	} else {
		return FUNGSI(x-1) + x
	}
}
func eksponen(x int) int {
	if x == 1 {
		return 2
	} else {
		return eksponen(x-1) * 2
	}
}
func faktorial(a int) int {
	if a == 0 || a == 1 {
		return 1
	} else {
		return a * faktorial(a-1)
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
