package main

import (
	"Project-Alpro/alpro"
	"Project-Alpro/contohsoal"

	// "Project-Alpro/linklanjutan"
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro"
	"bufio"
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	var input string
	var stop bool = false // Variabel untuk menghentikan loop

	for !stop {
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan menu
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)
		bufio.NewReader(os.Stdin).ReadString('\n')

		switch input {
		case "1":
			atribut.ClearScreen()
			fmt.Println("Menuju Pengenalan Pemrograman...")
			atribut.Loading()
			pengpro.MainMenu()
		case "2":
			atribut.ClearScreen()
			fmt.Println("Menuju Algoritma Pemrograman...")
			atribut.Loading()
			alpro.MainMenu()
		case "3":
			atribut.ClearScreen()
			fmt.Println("Menuju Contoh Soal...")
			atribut.Loading()
			contohsoal.MainMenu()
		case "4":
			atribut.ClearScreen()
			fmt.Println("Belum ada materi")
			atribut.Loading()
			fmt.Println("Menuju Menu utama...")
			return
		case "10":
			atribut.ClearScreen()
			fmt.Println("COMING SOON!!!")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading()
		case "11":
			atribut.ClearScreen()
			fmt.Println("COMING SOON!!!")
			fmt.Println("Menuju Menu utama...")
			atribut.Loading()
		case "0":
			atribut.ClearScreen()
			figure.NewFigure("THX THX THX", "basic", true).Blink(5000, 1000, 500)
			return
		default:
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading()
			fmt.Println()
		}

	}

}

func menu() {
	fmt.Println(`
=========================================
BELAJAR BARENG ALWAN & FATHUR!!!!!!
=========================================
Materi yang tersedia:
=========================================
1. Pengenalan Pemrograman
2. Algoritma dan Pemrograman 
3. Contoh Soal
4. Link untuk belajar lebih lanjut


=========================================
10. about us:
11. link referensi:
=========================================

0. keluar
	`)
}
