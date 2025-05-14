package array

import (
	"Project-Alpro/atribut"
	"fmt"
)

var pilihan string
var pilihan2 string

func Submenu() {
	fmt.Println(`
==================
Selamat Datang di Pembelajaran Array
==================
1. apa itu array
2. soal array
3. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}

func MainMenu() {
	atribut.ClearScreen()
	Submenu()
	switch pilihan {
	case "1":
		atribut.ClearScreen()
		belajarArray()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}

	case "2":
		atribut.ClearScreen()
		soalArray()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "3":
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func belajarArray() {
	fmt.Print(`
	
	deklarasi array
	var names [3]string
	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy" // tidak bisa lebih dari 3 data jika di deklarasi hanya ada 3
	//atau dengan
	var num = [4]int{
		90, 80, 95, 90,
	}
	var kosong = [10]int{}

	//cara akses data dari array
	//1. panggil array
	fmt.Println(names) //[Eko Kurniawan Khannedy]
	fmt.Println(num)   // [90 80 95 90]
	fmt.Println()

	//2. pemanggilan data indeks 0 berupa data pertama
	fmt.Println(names[0]) //eko
	fmt.Println(num[0])   //90
	fmt.Println()

	//3. pemanggilan jumlah panjang array, bukan panjang data
	fmt.Println(len(names))  //3
	fmt.Println(len(num))    // 4
	fmt.Println(len(kosong)) //10bb
	fmt.Println()

	//4. mengubah value dari index tertentu
	names[0] = "Alwan"    // [Eko Kurniawan Khannedy] -> [Alwan Kurniawan Khannedy]
	num[1] = 100          // [90 100 95 90] -> [90 100 95 90]
	fmt.Println(names[0]) //Alwan
	fmt.Println(num[0])   // 100
	fmt.Println()
`)
}

func soalArray() {
	fmt.Println("COMING SOON!")
}
