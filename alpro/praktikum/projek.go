package praktikum

import (
	array "Project-Alpro/alpro/praktikum/Array"
	binary "Project-Alpro/alpro/praktikum/Binary_Search"
	fungsi "Project-Alpro/alpro/praktikum/Fungsi"
	prosedur "Project-Alpro/alpro/praktikum/Prosedur"
	rekursif "Project-Alpro/alpro/praktikum/Rekursif"
	maxmin "Project-Alpro/alpro/praktikum/Searching"
	selection "Project-Alpro/alpro/praktikum/Selection_Sort"
	sequential "Project-Alpro/alpro/praktikum/Sequential_Search"
	Tbentukan "Project-Alpro/alpro/praktikum/Tipe_Bentukan"
	insertion "Project-Alpro/alpro/praktikum/insertion_sort"
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama dan mengelola pilihan user.
func MainMenu() {
	for { // Loop utama, akan terus berjalan sampai user memilih '0'
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
10. Insertion Sort
0. Keluar
		`)
		fmt.Print("Masukkan pilihan Anda: ")
		var pilihan string
		fmt.Scan(&pilihan)

		if pilihan == "0" {
			fmt.Println("Terima kasih! Sampai jumpa.")
			break // Keluar dari loop utama MainMenu
		}

		// Panggil submenu dengan pilihan yang didapat
		submenu(pilihan)
	}
}

// tanyaKembali mengelola pilihan user setelah menyelesaikan soal.
func tanyaKembali(masukan string, currentParentChoice string) {
	for { // Loop untuk validasi input, akan terus berjalan sampai input valid
		fmt.Printf("Tetap di %s ini (1), pilih %s lain (2) atau kembali ke menu utama (3)? ", masukan, masukan)
		var choice2 string
		fmt.Scan(&choice2)
		choice2 = strings.TrimSpace(choice2)

		switch choice2 {
		case "1":
			kumpulanTugasPendahuluan(currentParentChoice, "") // Panggil tanpa n2, nanti akan diminta lagi di submenu
			break
		case "2":
			submenu(currentParentChoice) // Kembali ke submenu saat ini
			break
		case "3":
			break
		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan 1, 2, atau 3.")
			// Loop akan mengulang secara otomatis karena tidak ada break di sini
		}
	}
}

// submenu menampilkan sub-menu dan mengelola pilihan user.
func submenu(parentChoice string) {
	for { // Loop agar user bisa pilih soal lain di submenu yang sama
		var choise string // Perbaiki ejaan
		switch parentChoice {
		case "1": // Fungsi
			fmt.Println(`
Soal-soal fungsi
Pilih Soal:
1. Fungsi Celcius, Reamur, Fahrenheit, Kelvin
2. Fungsi Rubah huruf kecil menjadi huruf besar
3. Fungsi Cek bilangan
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal fungsi: ")
			fmt.Scan(&choise)
		case "2": // Prosedur
			fmt.Println(`
Soal-soal Prosedur
Pilih Soal:
1. Prosedur keliling dan luas lingkaran
2. Prosedur hitung menang,draw dan kalah bola
3. Prosedur kalkulator sederhana
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal prosedur: ")
			fmt.Scan(&choise)
		case "3": // Rekursif
			fmt.Println(`
Soal-soal rekursif
Pilih Soal:
1. Rekursif fibonaci
2. Rekursif Faktor Persekutuan Terbesar
3. Rekursif Faktorial
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal rekursif: ")
			fmt.Scan(&choise)
		case "4": // Tipe Bentukan
			fmt.Println(`
Soal-soal tipe bentukan
Pilih Soal:
1. Tipe Bentukan merek, tahun produksi dan kecepatan mobil
2. Tipe Bentukan titik koordinat dan warna
3. Tipe Bentukan gaji pegawai
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal tipe bentukan: ")
			fmt.Scan(&choise)
		case "5": // Array
			fmt.Println(`
Soal-soal array
Pilih Soal:
1. Input dan ouput array sederhana
2. Input dan output array dengan baca dan cetak data
3. Jumlah dan Rata-rata Nilai dengan Array
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal array: ")
			fmt.Scan(&choise)
		case "6": // Max/Min Searching
			fmt.Println(`
Soal-soal maxmin searching
Pilih Soal:
1. Mencari nilai Max
2. Mencari nilai Min
3. Gabungan Keduanya
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal maxmin searching: ")
			fmt.Scan(&choise)
		case "7": // Sequential Search
			fmt.Println(`
Soal-soal Sequential Search
Pilih Soal:
1. Sequential Search dengan return boolean
2. Sequential Search dengan return index
3. Modifikasi Algoritma Sequential Search
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal sequential search: ")
			fmt.Scan(&choise)
		case "8": // Binary Search
			fmt.Println(`
Soal-soal Binary Search
Pilih Soal:
1. Binary Search dengan return boolean
2. Binary Search dengan return index
3. Modifikasi Algoritma Binary Search
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal binary search: ")
			fmt.Scan(&choise)
		case "9": // Selection Sort
			fmt.Println(`
Soal-soal Selection Sort
Pilih Soal:
1. Selection Sort
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal selection sort: ")
			fmt.Scan(&choise)
		case "10": // Insertion Sort
			fmt.Println(`
Soal-soal Insertion Sort
Pilih Soal:
1. Insertion Sort
0. Kembali ke Menu Utama
			`)
			fmt.Print("Masukkan pilihan soal insertion sort: ")
			fmt.Scan(&choise)
		default:
			fmt.Println("Pilihan kategori tidak valid. Silakan coba lagi.")
			return // Keluar dari submenu jika kategori tidak valid
		}

		if choise == "0" {
			break // Keluar dari loop submenu, kembali ke MainMenu
		}

		kumpulanTugasPendahuluan(parentChoice, choise)
		// Setelah kumpulanTugasPendahuluan selesai dan user kembali dari tanyaKembali,
		// loop submenu akan mengulang untuk menampilkan menu soal lagi.
	}
}

// kumpulanTugasPendahuluan menjalankan fungsi/prosedur yang dipilih oleh user.
func kumpulanTugasPendahuluan(n1, n2 string) {
	switch n1 {
	case "1": // Fungsi
		switch n2 {
		case "1":
			fungsi.Soal1fungsi()
			tanyaKembali("fungsi", n1)
		case "2":
			fungsi.Soal2fungsi()
			tanyaKembali("fungsi", n1)
		case "3":
			fungsi.Soal3fungsi()
			tanyaKembali("fungsi", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "2": // Prosedur
		switch n2 {
		case "1":
			prosedur.Soal1prosedur()
			tanyaKembali("prosedur", n1)
		case "2":
			prosedur.Soal2prosedur()
			tanyaKembali("prosedur", n1)
		case "3":
			prosedur.Soal3prosedur()
			tanyaKembali("prosedur", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "3": // Rekursif
		switch n2 {
		case "1":
			rekursif.Soal1rekursif()
			tanyaKembali("rekursif", n1)
		case "2":
			rekursif.Soal2rekursif()
			tanyaKembali("rekursif", n1)
		case "3":
			rekursif.Soal3rekursif()
			tanyaKembali("rekursif", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "4": // Tipe Bentukan
		switch n2 {
		case "1":
			Tbentukan.Soal1bentukan()
			tanyaKembali("tipe bentukan", n1)
		case "2":
			Tbentukan.Soal2bentukan()
			tanyaKembali("tipe bentukan", n1)
		case "3":
			Tbentukan.Soal3bentukan()
			tanyaKembali("tipe bentukan", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "5": // Array
		switch n2 {
		case "1":
			array.Soal1array()
			tanyaKembali("array", n1)
		case "2":
			array.Soal2array()
			tanyaKembali("array", n1)
		case "3":
			array.Soal3array()
			tanyaKembali("array", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "6": // Max/Min Searching
		switch n2 {
		case "1":
			maxmin.Soal1searching()
			tanyaKembali("maxmin", n1)
		case "2":
			maxmin.Soal2searching()
			tanyaKembali("maxmin", n1)
		case "3":
			maxmin.Soal3searching()
			tanyaKembali("maxmin", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "7": // Sequential Search
		switch n2 {
		case "1":
			sequential.Soal1SeqSearch()
			tanyaKembali("seqsearch", n1)
		case "2":
			sequential.Soal2SeqSearch()
			tanyaKembali("seqsearch", n1)
		case "3":
			sequential.Soal3SeqSearch()
			tanyaKembali("seqsearch", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "8": // Binary Search
		switch n2 {
		case "1":
			binary.Soal1BinSearch()
			tanyaKembali("binsearch", n1)
		case "2":
			binary.Soal2BinSearch()
			tanyaKembali("binsearch", n1)
		case "3":
			binary.Soal3BinSearch()
			tanyaKembali("binsearch", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "9": // Selection Sort
		switch n2 {
		case "1":
			selection.SoalSelectionSort()
			tanyaKembali("selection", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	case "10": // Insertion Sort
		switch n2 {
		case "1":
			insertion.SoalInsertionSort()
			tanyaKembali("insertion", n1)
		default:
			fmt.Println("Pilihan soal tidak valid. Silakan coba lagi.")
		}
	default:
		fmt.Println("Pilihan kategori tidak valid.")
	}
}
