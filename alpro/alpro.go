package alpro

import (
	"Project-Alpro/alpro/array"
	"Project-Alpro/alpro/kalkulator"
	"Project-Alpro/alpro/rekursif"
	"fmt"
	"os"
	"os/exec"
)

var pilihan string

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MainMenu() {
	clearScreen()
	Submenu()
	switch pilihan {
	case "1":
		clearScreen()
		array.MainMenu()
	case "2":
		clearScreen()
		kalkulator.MainMenu()
	case "3":
		clearScreen()
		rekursif.MainMenu()
	case "4":
		clearScreen()
		fmt.Println("Terima kasih telah menggunakan program ini.")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func Submenu() {
	fmt.Println(`
====================================
Selamat Datang di Algoritma dan Pemrograman
====================================
1. Pembelajaran Array
2. Kalkulator
3. Rekursif
4. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}
