package alpro

import (
	"Project-Alpro/alpro/p_array"
	"Project-Alpro/alpro/projek_kami"
	"Project-Alpro/atribut"
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
			fmt.Println("Coming soon")
		case "3":
			atribut.ClearScreen()
			projek_kami.Project_menu()
		case "0":
			atribut.ClearScreen()
			atribut.Loading()
			fmt.Println("Menuju menu utama...")
			return
		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading()
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
1. Pembelajaran
2. Contoh Soal
3. Projek Kami

0. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}
