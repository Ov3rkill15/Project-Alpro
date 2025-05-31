package p_array 

import (
	"Project-Alpro/atribut" 
	"fmt"
	"strings"
)


func MainMenu() {
	var pilihan string
	var pilihan2 string 
	fmt.Println(`
==================================
Selamat Datang di Pembelajaran Array
==================================
1. Apa itu Array
2. Soal Array
3. Keluar
    `)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Scanln() // Clear newline
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarArray() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			berhenti2 := false
			for !berhenti2 {
				fmt.Scan(&pilihan2)
				fmt.Scanln()
				if strings.ToLower(pilihan2) == "y" {
					MainMenu()
				} else if strings.ToLower(pilihan2) == "n" {
					return
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
					fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
				}
			}
		} else {
			return
		}

	case "2":
		atribut.ClearScreen()
		soalArray()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		
		berhenti2 := false
		for !berhenti2 {
			fmt.Scan(&pilihan2)
			fmt.Scanln() 
			if strings.ToLower(pilihan2) == "y" {
				MainMenu() 
				return     
			} else if strings.ToLower(pilihan2) == "n" {
				return 
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
			}
		}
	case "3":
		fmt.Println("Terima kasih telah menggunakan program pembelajaran array.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '3'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
		MainMenu()
	}
}

func belajarArray() bool {
	halamanSekarang := 1
	totalHalaman := 2
	var berhenti bool = true 
	var Choice string

	for berhenti { 
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Array di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan Anda ingin menyimpan daftar nilai ujian siswa, daftar nama buah, atau daftar suhu harian. Daripada membuat variabel terpisah untuk setiap nilai (misalnya, `nilai1`, `nilai2`, `nilai3`), akan jauh lebih efisien jika kita bisa mengelompokkan semuanya dalam satu 'wadah'.")
			fmt.Println("\nItulah fungsi dari **Array** dalam pemrograman.")
			fmt.Println("\nSebuah **Array** adalah koleksi berurutan dari elemen-elemen yang memiliki **tipe data yang sama** dan **ukuran (panjang) yang tetap**.")
			fmt.Println("\n**Konsep Penting Array:**")
			fmt.Println("- **Ukuran Tetap:** Setelah array dibuat, ukurannya tidak bisa diubah. Anda harus menentukan berapa banyak elemen yang akan disimpan di dalamnya saat deklarasi.")
			fmt.Println("- **Tipe Data Seragam:** Semua elemen dalam array harus memiliki tipe data yang sama (misalnya, semua `int`, semua `string`, dll.).")
			fmt.Println("- **Akses Berbasis Indeks:** Setiap elemen dalam array memiliki posisi atau 'indeks' numerik. Indeks dimulai dari 0 untuk elemen pertama, 1 untuk elemen kedua, dan seterusnya.")
			fmt.Println("\n**Contoh Deklarasi Array di Go:**")
			fmt.Println("```go")
			fmt.Println("var angka [5]int        // Deklarasi array 'angka' dengan 5 elemen int")
			fmt.Println("var namaHari [7]string // Deklarasi array 'namaHari' dengan 7 elemen string")
			fmt.Println("```")
			fmt.Println("\nDi sini, `[5]int` berarti \"array dengan 5 integer\".")
		case 2:
			fmt.Println("### Menggunakan Array: Inisialisasi dan Akses Elemen")
			fmt.Println("Setelah array dideklarasikan, Anda bisa memberinya nilai (inisialisasi) dan mengakses elemen-elemennya menggunakan indeks.")
			fmt.Println("\n**Inisialisasi Array:**")
			fmt.Println("```go")
			fmt.Println("// Inisialisasi satu per satu")
			fmt.Println("var angka [3]int")
			fmt.Println("angka[0] = 10")
			fmt.Println("angka[1] = 20")
			fmt.Println("angka[2] = 30")
			fmt.Println("\n// Inisialisasi langsung saat deklarasi")
			fmt.Println("huruf := [4]rune{'a', 'b', 'c', 'd'}")
			fmt.Println("\n// Inisialisasi tanpa menentukan ukuran, Go akan menghitungnya otomatis")
			fmt.Println("buah := [...]string{\"Apel\", \"Jeruk\", \"Mangga\"} // Ukuran otomatis 3")
			fmt.Println("```")
			fmt.Println("\n**Mengakses Elemen Array:**")
			fmt.Println("Anda mengakses elemen array menggunakan `namaArray[indeks]`. Ingat, indeks dimulai dari 0!")
			fmt.Println("```go")
			fmt.Println("fmt.Println(angka[0])  // Output: 10")
			fmt.Println("fmt.Println(huruf[2])  // Output: c")
			fmt.Println("fmt.Println(buah[1])   // Output: Jeruk")
			fmt.Println("```")
			fmt.Println("\nAnda juga bisa mengetahui panjang (ukuran) array menggunakan `len()`:")
			fmt.Println("```go")
			fmt.Println("fmt.Println(len(angka)) // Output: 3")
			fmt.Println("```")
			fmt.Println("\n---")
			fmt.Println("### Array vs. Slice (Sekilas)")
			fmt.Println("Di Go, array memiliki ukuran tetap, yang membuatnya kurang fleksibel untuk banyak kasus penggunaan. Oleh karena itu, Anda akan sering menemukan **Slice** yang lebih banyak digunakan. Slice adalah abstraksi di atas array yang memungkinkan ukuran dinamis. Namun, pemahaman tentang array sangat penting sebagai fondasi untuk memahami slice.")
			fmt.Println("\nPelajari array dengan baik, dan Anda akan memiliki dasar yang kuat untuk struktur data yang lebih canggih di Go!")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln()

				if strings.ToLower(Choice) == "y" {
					halamanSekarang++
					berhenti2 = true
				} else if strings.ToLower(Choice) == "n" {
					atribut.ClearScreen()
					berhenti = false
					berhenti2 = true 
					MainMenu() 
					return false 
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else { 
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Materi selesai! Mau lanjut ke soal (s), atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() 

				if strings.ToLower(Choice) == "s" {
					soalArray()
					berhenti = false 
					berhenti2 = true 
					return true      
				} else if strings.ToLower(Choice) == "n" {
					berhenti = false 
					berhenti2 = true 
					MainMenu()
					atribut.ClearScreen()
					return false 
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 's' atau 'n'.")
				}
			}
		}
	}
	return true
}

func soalArray() {
	atribut.ClearScreen()
	fmt.Println("=================")
	fmt.Println(" 	SOAL ARRAY")
	fmt.Println("=================")
	fmt.Println("COMING SOON!")
	fmt.Println("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln() // Wait for Enter
}
