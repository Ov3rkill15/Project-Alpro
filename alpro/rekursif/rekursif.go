package rekursif

import (
	"Project-Alpro/atribut"
	"fmt"
)

var pilihan string

func MainMenu() {
	atribut.ClearScreen()
	Submenu()
	switch pilihan {
	case "1":
		atribut.ClearScreen()
		tampilanfibo()
		ifelse(1) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "2":
		atribut.ClearScreen()
		tampilanx()
		ifelse(2) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "3":
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilanFUNGSI()
		ifelse(3) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "4":
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilaneksponen()
		ifelse(4) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "5":
		atribut.ClearScreen() // Membersihkan layar sebelum menampilkan hasil
		tampilanFaktorial()
		ifelse(5) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	case "6":
		atribut.ClearScreen()
		Submenu1()
		tampilankomper()
		ifelse(6) // Memanggil fungsi ifelse untuk menampilkan menu pilihan
	}
}
func Submenu() {
	fmt.Println(`
=================
Jenis rekursif mana?
=================
1. Fibonacci
2. (x - 3) + (x - 2) + (x - 1)
3. Fungsi nilai ke-n beda tingkat dua
4. Eksponen 2*n
5. Faktorial
6. Implementasi faktorial
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}
func Submenu1() {
	fmt.Println(`
=================
Implementasi Rekursif!
=================
1. kombinasi
2. permutasi
	`)
}
