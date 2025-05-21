package praktikum

import (
	array "Project-Alpro/alpro/projek_kami/praktikum/Array"
	fungsi "Project-Alpro/alpro/projek_kami/praktikum/Fungsi"
	"Project-Alpro/atribut"

	//prosedur "Project-Alpro/alpro/projek_kami/praktikum/Prosedur"
	// rekursif "Project-Alpro/alpro/projek_kami/praktikum/Rekursif"
	// Tbentukan "Project-Alpro/alpro/projek_kami/praktikum/Tipe_Bentukan"
	// maxmin "Project-Alpro/alpro/projek_kami/praktikum/Max_Min_Searching"
	// sequential "Project-Alpro/alpro/projek_kami/praktikum/Sequential_Search"
	// binary "Project-Alpro/alpro/projek_kami/praktikum/Binary_Search"
	// selection "Project-Alpro/alpro/projek_kami/praktikum/Selection_Sort"

	"fmt"
)

var pilihan string
var Choice2 string

func MainMenu() {
	fmt.Println(`
Berikut Adalah Pilihan Tugas Pendahuluan:
1. Fungsi
2. Prosedur
3. Rekursif
4. Tipe Bentukan
5. Array
6. Max/Min Searching
7. Sequential Search
8. Binary Search
9. Selection Sort
	`)
	fmt.Scan(&pilihan)
	atribut.ClearScreen()
	atribut.Loading(100)
	switch pilihan {
	case "1":
		var choise string
		fmt.Println("Soal - Soal Fungsi")
		fmt.Println(`Pilih Soal:
	1. Soal 1 (Celcius, Reamur, Fahrenheit, Kelvin)
	2. Soal 2 (Rubah huruf kecil menjadi huruf besar)
	3. Soal 3 (Cek bilangan)
	`)
		fmt.Scan(&choise)
		switch choise {
		case "1":
			fungsi.Soal1fungsi()
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			fmt.Scan(&Choice2)
			if Choice2 == "y" {
				MainMenu()
			}
		case "2":
			fungsi.Soal2fungsi()
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			fmt.Scan(&Choice2)
			if Choice2 == "y" {
				MainMenu()
			}
		case "3":
			fungsi.Soal3fungsi()
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			fmt.Scan(&Choice2)
			if Choice2 == "y" {
				MainMenu()
			}
		}
	case "2":
		fmt.Println("Prosedur")
	case "3":
		fmt.Println("Rekursif")
	case "4":
		fmt.Println("Tipe Bentukan")
	case "5":
		fmt.Println("Soal - Soal Array")
		fmt.Println(`Mau pilih soal yang mana?
	1. Soal inisialisasi dan cetak array sederhana
	2. 
	`)
		array.Soal1array()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&Choice2)
		if Choice2 == "y" {
			MainMenu()
		}
	case "6":
		fmt.Println("Max/Min Searching")
	case "7":
		fmt.Println("Sequential Search")
	case "8":
		fmt.Println("Binary Search")
	case "9":
		fmt.Println("Selection Sort")
	}
}
