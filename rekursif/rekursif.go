package rekursif

import (
	"fmt"
	"os"
	"os/exec"
)

func MainMenu() {
	clearScreen()
	Submenu()
	fmt.Print("Masukkan pilihan: ")
	var pilihan string
	fmt.Scan(&pilihan)
	switch pilihan {
	case "1":
		var a int
		clearScreen()
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
			var n, r int
			clearScreen() // Membersihkan layar sebelum menampilkan hasil
			fmt.Println("==========")
			fmt.Println("Implementasi Rekursif:Kombinasi")
			fmt.Println("==========")
			fmt.Print("Masukkan nilai a dan b (aCb): ")
			fmt.Scan(&n, &r)
			fmt.Println("Hasilnya adalah: ", kombinasi(n, r))
			fmt.Println()
		case "2":
			var n, r int
			clearScreen() // Membersihkan layar sebelum menampilkan hasil
			fmt.Println("==========")
			fmt.Println("Implementasi Rekursif:Permutasi")
			fmt.Println("==========")
			fmt.Print("Masukkan nilai a dan b (aCb): ")
			fmt.Scan(&n, &r)
			fmt.Println("Hasilnya adalah: ", permutasi(n, r))
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
	}
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	return result
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
