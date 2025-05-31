package main

import (
	"Project-Alpro/alpro"
	"Project-Alpro/login"

	// "Project-Alpro/linklanjutan"
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	var input string
	var stop bool = false // Variabel untuk menghentikan loop
	var sign, helo string
	login.Mainlogin(&sign, &helo)
	if sign != "signin" {
		menuKonfirmasi("tidak")
	} else {
		menuKonfirmasi("ya")
	}

	for !stop {
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan menu
		user := figure.NewFigure(helo, "doom", true).String()
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		fmt.Print("\033[32m")
		fmt.Print(user)
		fmt.Print("\033[0m")
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

func menuKonfirmasi(n string) {
	// var konfirmasi string
	var konfirmasi2 string
	switch n {
	case "ya":
		fmt.Println("Kami sarankan untuk kamu langsung menuju pilihan 2")
		fmt.Println("1. Setuju")
		fmt.Println("2. Dari awal aja!")
		fmt.Scan(&konfirmasi2)
		cek := strings.ToLower(konfirmasi2)
		switch cek {
		case "1", "setuju":
			fmt.Println("Menuju Algoritma Pemrograman!")
			atribut.Loading(100)
			atribut.ClearScreen()
			alpro.MainMenu()
		case "2", "dari awal aja", "awal", "dari awal":
			fmt.Println("Oke, mari belajar sama-sama!")
			atribut.Loading(100)
			atribut.ClearScreen()
			menu()
		}
	case "tidak":
		fmt.Println(`
			=========================================
			BELAJAR BARENG ALWAN & FATHUR!!!!!!
			=========================================
			Apakah kamu pernah belajar pemrograman go?
			=========================================
			1. ya
			2. tidak
			=========================================`)
		fmt.Scan(&konfirmasi2)
		cek := strings.ToLower(konfirmasi2)
		switch cek {
		case "1", "setuju":
			fmt.Println("Menuju Algoritma Pemrograman!")
			atribut.Loading(100)
			atribut.ClearScreen()
			alpro.MainMenu()
		case "2", "dari awal aja", "awal", "dari awal":
			fmt.Println("Oke, mari belajar sama-sama!")
			atribut.Loading(100)
			atribut.ClearScreen()
			menu()
		}
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
