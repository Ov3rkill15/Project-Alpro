package Rekursif

import (
	"fmt"
	"strings"
)

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func Soal1rekursif() bool {
	fmt.Print(`
============================================
package main

import "fmt"

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	var bil int
	fmt.Scan(&bil)
	fmt.Println(fibo(bil))
}

============================================
		`)
	var bil int
	fmt.Scan(&bil)
	fmt.Println(fibo(bil))
	fmt.Println()
	var pilihan string
	// Variabel untuk mengontrol loop input validasi
	inputValid := false
	for !inputValid {
		fmt.Println("\n------------------------------------")
		fmt.Print("Ingin kembali ke menu utama (n)? ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		if strings.ToLower(pilihan) == "n" {
			inputValid = true // Input valid, keluar dari loop input
			// Mengembalikan true berarti kembali ke MainMenu
			return true // Kembali ke MainMenu
		} else {
			fmt.Println("Pilihan tidak valid. Harap masukkan 'n'.")
			// Loop akan terus berjalan sampai input 'n' diterima
		}
	}
	// Jika loop berakhir karena input valid, berarti 'n' sudah dimasukkan,
	return true
}
