package praktikum

import (
	array "Project-Alpro/alpro/projek_kami/praktikum/Array"
	binary "Project-Alpro/alpro/projek_kami/praktikum/Binary_Search"
	fungsi "Project-Alpro/alpro/projek_kami/praktikum/Fungsi"
	prosedur "Project-Alpro/alpro/projek_kami/praktikum/Prosedur"
	rekursif "Project-Alpro/alpro/projek_kami/praktikum/Rekursif"
	maxmin "Project-Alpro/alpro/projek_kami/praktikum/Searching"
	selection "Project-Alpro/alpro/projek_kami/praktikum/Selection_Sort"
	sequential "Project-Alpro/alpro/projek_kami/praktikum/Sequential_Search"
	Tbentukan "Project-Alpro/alpro/projek_kami/praktikum/Tipe_Bentukan"
	"fmt"
)

var pilihan string
var Choise string
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
	submenu(pilihan)
}

func tanyaKembali(masukan string) {
	fmt.Printf("Tetap di %s ini, pilih %s lain atau kembali ke menu utama? (1/2/3): ", masukan, masukan)
	fmt.Scan(&Choice2)
	if Choice2 == "1" {
		kumpulanTugasPendahuluan(pilihan, Choise)
	} else if Choice2 == "2" {
		submenu(pilihan)
	} else {
		MainMenu()
	}
}

func submenu(subshoice string) {
	switch subshoice {
	case "1":
		fmt.Println(
			`Soal-soal fungsi
Pilih Soal:
1. Fungsi Celcius, Reamur, Fahrenheit, Kelvin
2. Fungsi Rubah huruf kecil menjadi huruf besar
3. Fungsi Cek bilangan
0.keluar
	`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("1", Choise)
	case "2":
		fmt.Println(
			`Soal-soal Prosedur
Pilih Soal:
1. Prosedur keliling dan luas lingkaran
2. Prosedur hitung menang,draw dan kalah bola
3. Prosedur kalkulator sederhana
0.keluar
	`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("2", Choise)
	case "3":
		fmt.Println(
			`Soal-soal rekursif
Pilih Soal:
1. Rekursif fibonaci
2. Rekursif Faktor Persekutuan Terbesar
3. Rekursif Faktorial
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("3", Choise)
	case "4":
		fmt.Println(
			`Soal-soal tipe bentukan
Pilih Soal:
1. Tipe Bentukan merek, tahun produksi dan kecepatan mobil
2. Tipe Bentukan titik koordinat dan warna
3. Tipe Bentukan gaji pegawai
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("4", Choise)
	case "5":
		fmt.Println(
			`Soal-soal array
Pilih Soal:
1. Input dan ouput array sederhana
2. Input dan output array dengan baca dan cetak data
3. Jumlah dan Rata-rata Nilai dengan Array
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("5", Choise)
	case "6":
		fmt.Println(
			`Soal-soal maxmin searching
Pilih Soal:
1. Mencari nilai Max
2. Mencari nilai Min
3. Gabungan Keduanya
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("6", Choise)
	case "7":
		fmt.Println(
			`Soal-soal Sequential Search
Pilih Soal:
1. Sequential Search dengan return boolean
2. Sequential Search dengan return index
3. Modifikasi Algoritma Sequential Search
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("7", Choise)
	case "8":
		fmt.Println(
			`Soal-soal Binary Search
Pilih Soal:
1. Binary Search dengan return boolean
2. Binary Search dengan return index
3. Modifikasi Algoritma Binary Search
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("8", Choise)
	case "9":
		fmt.Println(
			`Soal-soal Selection Sort
Pilih Soal:
1. Selection Sort
0. keluar
			`)
		fmt.Scan(&Choise)
		kumpulanTugasPendahuluan("9", Choise)
	}
}

func kumpulanTugasPendahuluan(n1, n2 string) {
	switch n1 {
	case "1":
		switch n2 {
		case "1":
			fungsi.Soal1fungsi()
			tanyaKembali("fungsi")
		case "2":
			fungsi.Soal2fungsi()
			tanyaKembali("fungsi")
		case "3":
			fungsi.Soal3fungsi()
			tanyaKembali("fungsi")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "2":
		switch n2 {
		case "1":
			prosedur.Soal1prosedur()
			tanyaKembali("prosedur")
		case "2":
			prosedur.Soal2prosedur()
			tanyaKembali("prosedur")
		case "3":
			prosedur.Soal3prosedur()
			tanyaKembali("prosedur")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "3":
		switch n2 {
		case "1":
			rekursif.Soal1rekursif()
			tanyaKembali("rekursif")
		case "2":
			rekursif.Soal2rekursif()
			tanyaKembali("rekursif")
		case "3":
			rekursif.Soal3rekursif()
			tanyaKembali("rekursif")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "4":
		switch n2 {
		case "1":
			Tbentukan.Soal1bentukan()
			tanyaKembali("tipe bentukan")
		case "2":
			Tbentukan.Soal2bentukan()
			tanyaKembali("tipe bentukan")
		case "3":
			Tbentukan.Soal3bentukan()
			tanyaKembali("tipe bentukan")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "5":
		switch n2 {
		case "1":
			array.Soal1array()
			tanyaKembali("array")
		case "2":
			array.Soal2array()
			tanyaKembali("array")
		case "3":
			array.Soal3array()
			tanyaKembali("array")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "6":
		switch n2 {
		case "1":
			maxmin.Soal1searching()
			tanyaKembali("maxmin")
		case "2":
			maxmin.Soal2searching()
			tanyaKembali("maxmin")
		case "3":
			maxmin.Soal3searching()
			tanyaKembali("maxmin")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "7":
		switch n2 {
		case "1":
			sequential.Soal1SeqSearch()
			tanyaKembali("seqsearch")
		case "2":
			sequential.Soal2SeqSearch()
			tanyaKembali("seqsearch")
		case "3":
			sequential.Soal3SeqSearch()
			tanyaKembali("seqsearch")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "8":
		switch n2 {
		case "1":
			binary.Soal1BinSearch()
			tanyaKembali("binsearch")
		case "2":
			binary.Soal2BinSearch()
			tanyaKembali("binsearch")
		case "3":
			binary.Soal3BinSearch()
			tanyaKembali("binsearch")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	case "9":
		switch n2 {
		case "1":
			selection.SoalSelectionSort()
			tanyaKembali("selection")
		case "0":
			MainMenu()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}

}
