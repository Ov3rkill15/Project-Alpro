package rekursif

import (
	"fmt"
	"os"
	"os/exec"
	"time"
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
		ifelse(1) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "2":
		clearScreen()
		tampilanx()
		ifelse(2) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "3":
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilanFUNGSI()
		ifelse(3) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "4":
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilaneksponen()
		ifelse(4) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "5":
		clearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilanFaktorial()
		ifelse(5) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "6":
		clearScreen()
		Submenu1()
		tampilankomper()
		ifelse(6) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
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

func ifelse(n int) {
	fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
	fmt.Scan(&Choice2)
	for Choice2 == "y" && n == 1 {
		tampilanfibo()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 2 {
		tampilanx()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 3 {
		tampilanFUNGSI()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 4 {
		tampilaneksponen()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 5 {
		tampilanFaktorial()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}
	for Choice2 == "y" && n == 6 {
		tampilankomper()
		fmt.Println("Tetap di fungsi ini(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
		fmt.Scan(&Choice2)
	}

	if Choice2 == "r" {
		clearScreen()
		MainMenu()
	} else {
		clearScreen()
		fmt.Println("Kembali ke menu utama...")
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
		}
		return
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
	clearScreen()
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
	clearScreen()
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
	clearScreen() // Membersihkan layar sebelum menampilkan hasil
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
	clearScreen() // Membersihkan layar sebelum menampilkan hasil
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
