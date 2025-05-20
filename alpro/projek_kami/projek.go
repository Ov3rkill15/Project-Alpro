package projek_kami

import (
	kalkulator "Project-Alpro/alpro/projek_kami/kalkulator"
	praktikum "Project-Alpro/alpro/projek_kami/praktikum"
	rekursif "Project-Alpro/alpro/projek_kami/rekursif"

	"Project-Alpro/atribut"
	"fmt"
)

var pilihan string

func Project_menu() {
	atribut.ClearScreen()
	fmt.Println(`
1. Kalkulator
2. Rekursif
3. Kumpulan Tugas Pendahuluan praktikum
0. Keluar

Masukkan pilihan: `)

	fmt.Scan(&pilihan)
	if pilihan == "1" {
		atribut.ClearScreen()
		kalkulator.MainMenu()
	} else if pilihan == "2" {
		atribut.ClearScreen()
		rekursif.MainMenu()
	} else if pilihan == "3" {
		atribut.ClearScreen()
		praktikum.MainMenu()
	}
}
