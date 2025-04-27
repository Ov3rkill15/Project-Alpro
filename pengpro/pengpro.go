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
	fmt.Println("1. Introduction (Pendahuluan)Dasar-dasar Coding")
	fmt.Println("2. konsep Menara Hanoi")
	fmt.Println("3. Algoritma Sorting")
	fmt.Println("4. Algoritma Searching")
	fmt.Println("5. Algoritma djikstra and minimum spanning tree(kruskal's algorithm)")
	fmt.Println("6. Algoritma dynamic programmingg")
	fmt.Println("7. algoritma TSP and Neural neightbor")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih materi: ")
	fmt.Scan(&Choice)
	switch Choice {
	case 1:
		atribut.Openbrowser("https://1drv.ms/p/c/87a899df7dfa3990/EXebJrgPSrBDlN3TeSN9CJ0BwdKVzYaNfIFZSTBWiMu3Qw?e=jom0HL")
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case 2:
		atribut.Openbrowser("https://1drv.ms/p/c/87a899df7dfa3990/EfoM86_zr51Kpt9sXLBLOO0B7_kz35yYlbNUEQ27uExU9A?e=9dzJKB")
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case 3, 4, 5, 6, 7:
		fmt.Println("Belum ada materi")
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		} else {
			return
		}
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
