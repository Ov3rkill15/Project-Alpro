package main

import (
	"Project-Alpro/alpro"
	// "Project-Alpro/login"

	// "Project-Alpro/linklanjutan"
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro"
	"fmt"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	var input string
	var stop bool = false // Variabel untuk menghentikan loop
	menuKonfirmasi()

	for !stop {
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan menu
		// login.Login()         // Memanggil fungsi login dari package login
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			atribut.ClearScreen()
			fmt.Println("Menuju Pengenalan Pemrograman...")
			atribut.Loading(100)
			pengpro.MainMenu()
		case "2":
			atribut.ClearScreen()
			fmt.Println("Menuju Algoritma Pemrograman...")
			atribut.Loading(100)
			alpro.MainMenu()
		case "10":
			atribut.ClearScreen()
			fmt.Println("COMING SOON!!!")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading(100)
		case "0":
			stop = true
			atribut.ClearScreen()
			figure.NewFigure("THX THX THX", "basic", true).Blink(5000, 1000, 500)

		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(100)
			fmt.Println()
		}

	}

}

func menuKonfirmasi() {
	var konfirmasi string
	var konfirmasi2 string
	fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Apakah kamu pernah belajar pemrograman go?
=========================================
1. ya
2. tidak
=========================================`)
	fmt.Scan(&konfirmasi)
	switch konfirmasi {
	case "ya":
		fmt.Println("Kalau begitu kami sarankan untuk kamu langsung menuju pilihan 2")
		fmt.Println("1. Setuju")
		fmt.Println("2. Dari awal aja!")
		fmt.Scan(&konfirmasi2)
		if konfirmasi2 == "1" {
			fmt.Println("Menuju Algoritma Pemrograman!")
			atribut.Loading(100)
			atribut.ClearScreen()
			alpro.MainMenu()
		} else {
			fmt.Println("Oke, mari belajar sama-sama!")
			atribut.Loading(100)
			atribut.ClearScreen()
			menu()
		}
	case "tidak":
		fmt.Println("Kalau begitu mari belajar sama-sama!")
		atribut.Loading(100)
		atribut.ClearScreen()
		menu()
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}

}

func menu() {
	letsgo := figure.NewFigure("LET'S GO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(letsgo)     // Cetak teks ASCII
	fmt.Print("\033[0m")  // Reset warna ke default
	fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Materi yang tersedia:
=========================================
1. Pengenalan Pemrograman
2. Algoritma dan Pemrograman 


=========================================
10. about us:
=========================================

0. keluar
	`)
}
