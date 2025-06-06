package Searching

import (
	"fmt"
	"strings"
)

func Soal2searching() bool {
	fmt.Println(`
============================================
package main

import "fmt"

func main() {
	const NMAX int = 10
	var A [NMAX]int
	var i, n, MIN int
	fmt.Scan(&n)
	if n > NMAX {
		n = NMAX
	}
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MIN = 0
	for i = 1; i < n; i++ {
		if A[i] < A[MIN] {
			MIN = i
		}
	}
	fmt.Println(A[MIN], MIN)
}

============================================
		`)
	const NMAX int = 10
	var A [NMAX]int
	var i, n, MIN int
	fmt.Scan(&n)
	if n > NMAX {
		n = NMAX
	}
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	MIN = 0
	for i = 1; i < n; i++ {
		if A[i] < A[MIN] {
			MIN = i
		}
	}
	fmt.Println(A[MIN], MIN)
	fmt.Println()
	var pilihan string
	// Variabel untuk mengontrol loop input validasi
	inputValid := false
	for !inputValid {
		fmt.Println("\n------------------------------------")
		fmt.Print("Ingin kembali ke menu utama (n)? ") // Asumsi setelah soal selesai, hanya ada opsi kembali
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
	// jadi kita mengembalikan true untuk kembali ke MainMenu.
	return true
}
