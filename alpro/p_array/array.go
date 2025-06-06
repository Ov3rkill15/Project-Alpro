package p_array

import (
	soal "Project-Alpro/alpro/praktikum/Array"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu array.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	// isRunning akan menjadi false jika pengguna memilih untuk keluar (0).
	isRunning := true

	for isRunning {
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
=====================================
Selamat Datang di Pembelajaran Array
=====================================
1. Apa itu Array
2. Referensi Soal Array
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarArray sekarang akan mengembalikan true jika ingin tetap di MainMenu,
			// atau false jika pengguna memutuskan untuk keluar dari program sepenuhnya.
			shouldStayInMenu := belajarArray()
			if !shouldStayInMenu {
				isRunning = false // Jika belajarArray mengindikasikan keluar dari program, hentikan loop MainMenu
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// Asumsi: soal.Soal1array() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal1array()
			if !shouldStayInMenu {
				isRunning = false // Jika soal.Soal1array() mengindikasikan keluar dari program, hentikan loop MainMenu
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran array.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarArray mengelola materi pembelajaran array.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program.
func belajarArray() bool {
	halamanSekarang := 1
	totalHalaman := 2 // Hanya ada 2 halaman materi yang relevan
	var Choice string

	// isStudying akan menjadi false jika pengguna memilih 'n' untuk kembali ke menu utama.
	isStudying := true

	for isStudying {
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
		default:
			fmt.Println("Halaman tidak ditemukan.")
			// Jika halaman tidak valid, kita akan menganggap pengguna ingin kembali ke menu utama.
			return true // Kembali ke MainMenu
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		// Kita akan meminta input sampai valid atau pengguna memilih 'n'.
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya (outer loop)
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
			}
		}
	}
	// Setelah loop isStudying berakhir (karena Choice adalah 'n'), kita mengembalikan true
	// yang menandakan bahwa MainMenu harus terus berjalan.
	return true
}
