package admin

import (
	"Project-Alpro/login"
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func MainAdmin() {
	var input string
	user := figure.NewFigure("ADMIN", "doom", true).String()
	welcome := figure.NewFigure("WELCOME", "doom", true).String()
	fmt.Print("\033[31m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[31m") // Reset warna ke default
	fmt.Print("\033[31m")
	fmt.Print(user)
	fmt.Print("\033[0m")
	fmt.Println(`
=========================================
			ADMIN PERMISSIONS
=========================================
1. Riwayat Pengguna
2. Cetak Pengguna
3. Logout

0. keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&input)
}

func Riwayat() {
	var data riwayat
	var jumlahPengguna int
	// Hanya masukkan sesuai kapasitas array
	jumlahPengguna = 0
	for i := 0; i < 100 && i < len(data); i++ {
		users[i] = data[i]
		jumlahPengguna++
	}
}

func CetakPengguna() {
	users := login.GetUsers()
	jumlah := login.GetJumlahPengguna()

	fmt.Println("\n--- Daftar Pengguna ---")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s\n", i+1, users[i].Username)
	}
}
