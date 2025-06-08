package Binsearch

import (
	atribut "Project-Alpro/alpro/praktikum/atributPraktikum"
	"fmt"
	"strings"
)

func BinSearch1(tab atribut.Array, n, x int) bool {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && tab.Info[mid] != x {
		if x < tab.Info[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > -1 && tab.Info[mid] == x
}

func Soal1BinSearch() bool {
	fmt.Println(`
============================================
package main

import "fmt"

type Array struct {
	Info [NMAX]int
	N    int
}

func BacaData1(A *Array) {
	fmt.Scan(&A.N)
	if A.N > NMAX {
		A.N = NMAX
	}
	for i := 0; i < A.N; i++ {
		fmt.Scan(&A.Info[i])
	}
}

func CetakData1(A Array) {
	if A.N == 0 {
		fmt.Println("Array kosong")
	}
	for i := 0; i < A.N; i++ {
		fmt.Print(A.Info[i], " ")
	}
	fmt.Println()
}

func BinSearch1(tab Array, n, x int) bool {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2
	for left <= right && tab.Info[mid] != x {
		if x < tab.Info[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > -1 && tab.Info[mid] == x
}

func main() {

	var data Array
	var x int
	fmt.Scan(&x)
	if data.N > NMAX {
		data.N = NMAX
	}
	BacaData1(&data)
	CetakData1(data)
	BinSearch1(data, data.N, x)

}
============================================`)

	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	BinSearch1(data, data.N, x)
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
