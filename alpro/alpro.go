package alpro

import (
	"Project-Alpro/alpro/p_array"
	"Project-Alpro/alpro/p_function"

	"Project-Alpro/alpro/p_maxmin"
	"Project-Alpro/alpro/p_procedure"
	"Project-Alpro/alpro/p_rekursif"
	"Project-Alpro/alpro/p_searching"

	// "Project-Alpro/alpro/p_sorting"
	p_tipeBentukan "Project-Alpro/alpro/p_tipebentukan"
	"Project-Alpro/alpro/projek_kami"
	"Project-Alpro/atribut"
	"Project-Alpro/contohsoal"
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

var pilihan string

func MainMenu() {
	for {
		atribut.ClearScreen()
		Submenu()
		switch pilihan {
		case "1":
			atribut.ClearScreen()
			p_array.MainMenu()
		case "2":
			atribut.ClearScreen()
			p_function.MainMenu()
		case "3":
			atribut.ClearScreen()
			p_procedure.MainMenu()
		case "4":
			atribut.ClearScreen()
			p_rekursif.MainMenu()
		case "5":
			atribut.ClearScreen()
			p_tipeBentukan.MainMenu()
		case "6":
			atribut.ClearScreen()
			p_maxmin.MainMenu()
		case "7":
			atribut.ClearScreen()
			p_searching.MainMenu()
		case "8":
			atribut.ClearScreen()
			// p_sorting.MainMenu()
		case "9":
			atribut.ClearScreen()
			projek_kami.Project_menu()
		case "10":
			atribut.ClearScreen()
			contohsoal.MainMenu()
		case "0":
			atribut.ClearScreen()
			return
		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(100)
		}
	}
}
func Submenu() {
	welcome := figure.NewFigure("ALPRO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")
	fmt.Println(`
====================================
Selamat Datang di Algoritma dan Pemrograman
====================================
1. Pembelajaran Array
2. Pembelajaran Function
3. Pembelajaran Procedure
4. Pembelajaran Rekursif
5. Pembelajaran tipeBentukan
6. Pembelajaran maxmin
7. Pembelajaran Searching
8. Pembelajaran Sorting
9. Projek kami
10. Contoh Soal

0. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}
