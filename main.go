package main

import (
	"Project-Alpro/alpro"
	"Project-Alpro/contohsoal"

	// "Project-Alpro/linklanjutan"
	"Project-Alpro/pengpro"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/common-nighthawk/go-figure" // Import go-figure
)

func main() {
	var input string
	var stop bool = false // Variabel untuk menghentikan loop

	for !stop {
		clearScreen() // Membersihkan layar sebelum menampilkan menu
		welcome := figure.NewFigure("WELCOME", "doom", true).String()
		fmt.Print("\033[32m") // Set warna hijau
		fmt.Print(welcome)    // Cetak teks ASCII
		fmt.Print("\033[0m")  // Reset warna ke default
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			clearScreen()
			fmt.Println("Menuju Pengenalan Pemrograman...")
			loading()
			pengpro.MainMenu()
		case "2":
			clearScreen()
			fmt.Println("Menuju Algoritma Pemrograman...")
			loading()
			alpro.MainMenu()
		case "3":
			clearScreen()
			fmt.Println("Menuju Contoh Soal...")
			loading()
			contohsoal.MainMenu()
		case "4":
			clearScreen()
			fmt.Println("Menuju Kumpulan Link...")
			loading()
			// linklanjutan.MainMenu()
		case "0":
			clearScreen()
			figure.NewFigure("THX THX THX", "basic", true).Blink(5000, 1000, 500)
			return
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			loading()
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

0. keluar
	`)
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func loading() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
}
