package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"Project-Alpro/kalkulator"
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
			kalkulator.Menu()
		case "2":
			fmt.Println("Thanks yak")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		// Menunggu pengguna untuk menekan Enter atau mengetik 'x'
		for {
			fmt.Println("Tekan Enter untuk kembali ke menu atau ketik 'x' untuk keluar...")
			input, _ = reader.ReadString('\n')
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
===============
DAFTAR FUNGSI
===============
1. FIBONACCI
2. (x - 3) + (x - 2) + (x - 1)
3. Fungsi nilai ke-n beda tingkat dua
4. Eksponen 2*n
5. Faktorial
6. 
7. keluar
	`)
}
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
