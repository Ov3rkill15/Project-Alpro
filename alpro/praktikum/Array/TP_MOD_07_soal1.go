package Array

import (
	"fmt"
	"strings"
)

func Soal1array() bool {
	fmt.Println(`
package main

import "fmt"
============================================
func main() {
	var n int
	const NMAX = 10
	fmt.Println("Masukkan jumlah data: (maks 10)")
	fmt.Scan(&n)
	var A [NMAX]string
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
		}
		fmt.Println()
		for i = 0; i < n; i++ {
			fmt.Println(A[i])
			}
		}
============================================
			`)
	fmt.Println("SILAHKAN DICOBA: ")
	var n int
	const NMAX = 10
	fmt.Println("Masukkan jumlah data: (maks 10)")
	fmt.Scan(&n)
	var A [NMAX]string
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println()
	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
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
