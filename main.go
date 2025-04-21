package main

import (
	"Project-Alpro/alpro"
	// "Project-Alpro/contohsoal"
	// "Project-Alpro/linklanjutan"
	"Project-Alpro/pengpro"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input

	for {
		clearScreen() // Membersihkan layar sebelum menampilkan menu
		menu()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			clearScreen()
			pengpro.MainMenu()
		case "2":
			clearScreen()
			alpro.MainMenu()
		case "3":
			clearScreen()
			// contohsoal.MainMenu()
		case "4":
			clearScreen()
			// linklanjutan.MainMenu()
		case "0":
			clearScreen()
			fmt.Println("Thanks yak")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		// Menunggu pengguna untuk menekan Enter atau mengetik 'x'
		for {
			fmt.Println("Tekan Enter untuk kembali ke menu atau ketik 'x' untuk keluar...")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "" {
				break // Kembali ke menu
			} else if input == "x" {
				fmt.Println("Thanks yak")
				return // Keluar dari program
			} else {
				fmt.Println("Input tidak valid. Silakan coba lagi.")
			}
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
