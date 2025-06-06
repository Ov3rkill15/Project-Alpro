package SequentialSearch

import (
	atribut "Project-Alpro/alpro/praktikum/atributPraktikum"
	"fmt"
	"strings"
)

func SeqSearch(T atribut.Array, N, X int) bool {
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for !ketemu && k < N {
		ketemu = T.Info[k] == X
		k = k + 1
	}
	return ketemu
}

func Soal1SeqSearch() bool {
	fmt.Println(`
============================================
package main

import (
	atribut "Project-Alpro/alpro/praktikum/atributPraktikum"
	"fmt"
)

func SeqSearch(T atribut.Array, N, X int) bool {
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for !ketemu && k < N {
		ketemu = T.Info[k] == X
		k = k + 1
	}
	return ketemu
}

func main() {
	var data atribut.Array
	var x int
	fmt.Scan(&x)
	if data.N > atribut.NMAX {
		data.N = atribut.NMAX
	}
	atribut.BacaData1(&data)
	atribut.CetakData1(data)
	SeqSearch(data, data.N, x)

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
	SeqSearch(data, data.N, x)
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
