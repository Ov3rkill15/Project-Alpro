package p_tipeBentukan

import (
	soal "Project-Alpro/alpro/praktikum/Tipe_Bentukan"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu manages the main navigation for the Struct and Type Alias module.
func MainMenu() {
	var pilihan string
	var pilihan2 string // Local variable for pilihan2

	fmt.Println(`
============================================
Selamat Datang di Pembelajaran Tipe Bentukan
============================================
1. Apa itu Tipe Bentukan (Struct) dan Alias
2. Soal referensi
0. Keluar
    `)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Scanln() // Clear newline
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarTipeBentukkan() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			fmt.Scan(&pilihan2)
			fmt.Scanln() // Clear newline
			if strings.ToLower(pilihan2) == "y" {
				MainMenu() // Rekursi
			} else {
				return
			}
		} else { // Jika false, artinya user ingin keluar dari modul
			return
		}
	case "2":
		atribut.ClearScreen()
		if soal.Soal1bentukan() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			berhenti2 := false
			for !berhenti2 {
				fmt.Scan(&pilihan2)
				fmt.Scanln()
				if strings.ToLower(pilihan2) == "y" {
					berhenti2 = true
					MainMenu()
					return
				} else if strings.ToLower(pilihan2) == "n" {
					berhenti2 = true
					return
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
					fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
				}
			}
		}
	case "0":
		fmt.Println("Terima kasih telah menggunakan program pembelajaran tipe bentukan.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '3'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Wait for Enter
		MainMenu()   // Recursive call for invalid input
	}
}

// belajarTipeBentukkan manages the display of learning material pages for Structs and Type Aliases.
// It returns true if the user wants to return to the MainMenu, false if they want to exit this module.
func belajarTipeBentukkan() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true
	var Choice string // Local variable for Choice

	for berhenti { // This loop now primarily controls page progression
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Tipe Bentukan (Struct) dan Alias di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, kita sering perlu mengelompokkan data yang berkaitan menjadi satu unit logis. Bayangkan Anda ingin menyimpan informasi tentang seorang siswa: nama, umur, dan nilai. Anda bisa menggunakan variabel terpisah, tapi itu bisa jadi berantakan. Di sinilah **Struct (Struktur)** dan **Type Alias (Alias Tipe)** berperan.")
			fmt.Println("\n---")
			fmt.Println("## Struct (Struktur)")
			fmt.Println("Sebuah **struct** adalah kumpulan dari nol atau lebih bidang (field) yang memiliki tipe data berbeda. Mirip dengan 'objek' dalam bahasa pemrograman lain, tetapi lebih fokus pada data. Struct memungkinkan kita untuk membuat tipe data *kustom* Anda sendiri.")
			fmt.Println("\n**Mengapa Menggunakan Struct?**")
			fmt.Println("- **Organisasi Data:** Mengelompokkan data yang saling terkait dalam satu unit yang rapi.")
			fmt.Println("- **Reusabilitas:** Membuat 'cetak biru' untuk objek dengan properti yang sama.")
			fmt.Println("- **Kejelasan Kode:** Meningkatkan keterbacaan kode karena data terstruktur dengan baik.")
			fmt.Println("\n**Contoh Deklarasi Struct:**")
			fmt.Println("```go")
			fmt.Println("type Person struct {")
			fmt.Println("    Name string")
			fmt.Println("    Age  int")
			fmt.Println("    City string")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDi sini, kita mendefinisikan sebuah struct bernama `Person` dengan tiga bidang: `Name` (string), `Age` (int), dan `City` (string).")
		case 2:
			fmt.Println("## Menggunakan Struct")
			fmt.Println("Setelah struct didefinisikan, Anda bisa membuat *instance* (contoh) dari struct tersebut dan mengakses bidang-bidangnya.")
			fmt.Println("\n**Membuat Instance Struct:**")
			fmt.Println("```go")
			fmt.Println("// Deklarasi dan inisialisasi struct")
			fmt.Println("var p1 Person") // Deklarasi variabel bertipe Person, nilai default kosong
			fmt.Println("p1.Name = \"Alice\"")
			fmt.Println("p1.Age = 30")
			fmt.Println("p1.City = \"New York\"")
			fmt.Println("\n// Inisialisasi dengan nilai literal")
			fmt.Println("p2 := Person{\"Bob\", 25, \"London\"}") // Urutan harus sesuai deklarasi bidang
			fmt.Println("\n// Inisialisasi dengan nama bidang (lebih disarankan untuk kejelasan)")
			fmt.Println("p3 := Person{Name: \"Charlie\", Age: 35, City: \"Paris\"}")
			fmt.Println("```")
			fmt.Println("\n**Mengakses Bidang Struct:**")
			fmt.Println("Anda dapat mengakses bidang-bidang struct menggunakan operator titik (`.`):")
			fmt.Println("```go")
			fmt.Println("fmt.Println(p1.Name) // Output: Alice")
			fmt.Println("fmt.Println(p2.Age)  // Output: 25")
			fmt.Println("```")
			fmt.Println("\n---")
			fmt.Println("## Type Alias (Alias Tipe)")
			fmt.Println("**Type alias** (alias tipe) adalah cara untuk memberi nama lain pada tipe data yang sudah ada. Ini tidak membuat tipe baru, melainkan hanya memberikan nama alternatif. Ini berguna untuk meningkatkan keterbacaan kode atau mempersingkat nama tipe yang panjang.")
			fmt.Println("\n**Contoh Deklarasi Type Alias:**")
			fmt.Println("```go")
			fmt.Println("type KTP string") // KTP adalah alias untuk tipe string")
			fmt.Println("type Usia int")   // Usia adalah alias untuk tipe int")
			fmt.Println("\nfunc main() {")
			fmt.Println("    var id KTP = \"1234567890\"")
			fmt.Println("    var umur Usia = 28")
			fmt.Println("    fmt.Println(\"Nomor KTP:\", id)")
			fmt.Println("    fmt.Println(\"Umur:\", umur)")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDi sini, `KTP` dan `Usia` hanyalah nama lain untuk `string` dan `int`. Mereka *kompatibel* satu sama lain. Anda bisa menggunakan `id` sebagai `string` dan `umur` sebagai `int` tanpa konversi eksplisit.")
		case 3:
			fmt.Println("## Perbedaan Utama: Struct vs. Type Alias")
			fmt.Println("- **Struct:** Membuat **tipe data baru** yang unik. Variabel dengan tipe struct yang berbeda (meskipun memiliki bidang yang sama) tidak dapat langsung ditugaskan satu sama lain tanpa konversi eksplisit.")
			fmt.Println("- **Type Alias:** Hanya memberikan **nama alternatif** pada tipe yang sudah ada. Tipe alias sepenuhnya kompatibel dengan tipe aslinya.")
			fmt.Println("\n**Kapan Menggunakan yang Mana?**")
			fmt.Println("- Gunakan **Struct** ketika Anda perlu mengelompokkan beberapa nilai (bisa berbeda tipe) menjadi satu unit logis dan ingin membuat tipe data *baru* yang berbeda secara fundamental.")
			fmt.Println("  Contoh: `Person`, `Product`, `Order`.")
			fmt.Println("- Gunakan **Type Alias** ketika Anda ingin meningkatkan keterbacaan kode dengan memberikan nama yang lebih deskriptif pada tipe yang sudah ada, tanpa membuat tipe baru.")
			fmt.Println("  Contoh: `UserID` untuk `int`, `EmailAddress` untuk `string`.")
			fmt.Println("\n---")
			fmt.Println("Dengan memahami struct dan type alias, Anda dapat membuat kode Go yang lebih terstruktur, mudah dibaca, dan mudah dikelola, terutama saat menangani data yang kompleks.")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			// Loop untuk input valid y/n
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() // Clear newline

				if strings.ToLower(Choice) == "y" {
					halamanSekarang++
					berhenti2 = true // Input valid, keluar dari loop
				} else if strings.ToLower(Choice) == "n" {
					atribut.ClearScreen()
					berhenti = false // Stop the main 'for berhenti' loop
					berhenti2 = true // Input valid, exit loop
					MainMenu()       // Recursive call to MainMenu
					return false     // Return false to exit belajarTipeBentukkan
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else { // On the last page
			// Loop untuk input valid s/n
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() // Clear newline
				if strings.ToLower(Choice) == "n" {
					berhenti = false // Stop the main 'for berhenti' loop
					berhenti2 = true // Input valid, exit loop
					MainMenu()       // Recursive call to MainMenu
					atribut.ClearScreen()
					return false // Return false to exit belajarTipeBentukkan
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'n'.")
				}
			}
		}
	}
	return true
}
