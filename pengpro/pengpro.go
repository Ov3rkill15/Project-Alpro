package pengpro

import (
	"Project-Alpro/atribut"
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

var Choice int
var Choice2 string

func MainMenu() {
	atribut.ClearScreen()
	Submenu()
}

func Submenu() {
	welcome := figure.NewFigure("PENGPRO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")
	fmt.Println("== Pengenalan Pemrograman ==")
	fmt.Println("1. Apa itu Pemrograman bahasa Go?")
	fmt.Println("2. Pengenalan tipe data(Number, String, Boolean)")
	fmt.Println("3. Variable dan Constant")
	fmt.Println("4. Konversi tipe data")
	fmt.Println("5. Operasi aritmatika dan logika")
	fmt.Println("6. Perulangan(For Loop, While Loop, Repeat Until)")
	fmt.Println("7. Percabangan(If Else, Switch Case)")
	fmt.Print("Pilih materi: ")
	fmt.Scan(&Choice)
	switch Choice {
	case 1:
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		}
	case 2:
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		}
	case 3, 4, 5, 6, 7:
		fmt.Println("Belum ada materi")
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		}
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
