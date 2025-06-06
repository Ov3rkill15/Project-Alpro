package p_rekursif

import (
	soal "Project-Alpro/alpro/praktikum/Rekursif" // Asumsi ini adalah package soal untuk Rekursif
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Rekursif.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
======================================
Selamat Datang di Pembelajaran Rekursif
======================================
1. Apa itu Rekursif
2. Referensi Soal
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarRekursif akan mengelola navigasi internalnya sendiri
			// dan akan kembali ke MainMenu saat selesai atau pengguna memilih 'n' dari dalamnya.
			shouldStayInMenu := belajarRekursif()
			if !shouldStayInMenu {
				// Jika belajarRekursif mengindikasikan keluar dari program, hentikan loop MainMenu.
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// Asumsi: soal.Soal1rekursif() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal1rekursif() // Panggil fungsi soal dari package 'soal'
			if !shouldStayInMenu {
				// Jika soal.Soal1rekursif() mengindikasikan keluar dari program, hentikan loop MainMenu.
				fmt.Println("Terima kasih telah menggunakan program pembelajaran rekursif.")
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran rekursif.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarRekursif mengelola materi pembelajaran rekursif.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program (misalnya dari halaman terakhir).
func belajarRekursif() bool {
	halamanSekarang := 1
	totalHalaman := 3 // Ada 3 halaman materi
	var Choice string

	// isStudying akan menjadi false jika pengguna memilih 'n' untuk kembali ke menu utama.
	isStudying := true

	for isStudying { // Loop tak terbatas sampai pengguna memutuskan untuk keluar dari materi ini
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Rekursif di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan kamu sedang melihat ke cermin yang diletakkan di depan cermin lain. Kamu akan melihat pantulan cermin di dalam cermin, di dalam cermin, dan seterusnya, seolah-olah tidak ada habisnya. Inilah esensi dari **rekursi** dalam pemrograman.")
			fmt.Println("\nDalam pemrograman, **rekursi** adalah teknik di mana sebuah fungsi memanggil dirinya sendiri secara berulang-ulang sampai kondisi dasar (base case) terpenuhi. Setiap kali fungsi memanggil dirinya sendiri, ia mengerjakan sebagian kecil dari masalah yang lebih besar, secara bertahap mendekati solusi.")
			fmt.Println("\n**Mengapa Rekursi Digunakan?**")
			fmt.Println("- **Elegansi dan Keringkasan:** Untuk masalah-masalah tertentu (seperti traversi pohon, faktorial, atau deret Fibonacci), solusi rekursif bisa jauh lebih ringkas dan mudah dibaca daripada solusi iteratif.")
			fmt.Println("- **Pemetaan Masalah Alami:** Beberapa masalah secara alami didefinisikan dalam istilah rekursif (misalnya, definisi faktorial n! adalah n * (n-1)!).")
			fmt.Println("- **Struktur Data Rekursif:** Rekursi sangat cocok untuk memproses struktur data rekursif seperti pohon biner atau *linked list*.")
		case 2:
			fmt.Println("Setiap fungsi rekursif harus memiliki dua bagian penting:")
			fmt.Println("\n1.  **Base Case (Kasus Dasar):** Ini adalah kondisi di mana fungsi tidak lagi memanggil dirinya sendiri. Ini adalah titik berhenti rekursi, mencegah terjadinya *loop* tak terbatas (Stack Overflow). Tanpa *base case*, rekursi akan terus memanggil dirinya sendiri sampai memori habis.")
			fmt.Println("\n2.  **Recursive Case (Kasus Rekursif):** Ini adalah bagian di mana fungsi memanggil dirinya sendiri dengan input yang lebih kecil atau masalah yang lebih sederhana. Setiap pemanggilan rekursif harus mendekati *base case*.")
			fmt.Println("\n**Contoh Sederhana: Menghitung Faktorial**")
			fmt.Println("Faktorial dari sebuah bilangan $n$ (ditulis $n!$) adalah hasil perkalian semua bilangan bulat positif kurang dari atau sama dengan $n$. Contoh: $5! = 5 \\times 4 \\times 3 \\times 2 \\times 1 = 120$.")
			fmt.Println("\n**Definisi Rekursif Faktorial:**")
			fmt.Println("- **Base Case:** $0! = 1$ atau $1! = 1$")
			fmt.Println("- **Recursive Case:** $n! = n \\times (n-1)!$ untuk $n > 1$")
			fmt.Println("\n```go")
			fmt.Println("func faktorial(n int) int {")
			fmt.Println("    if n == 0 { // Base Case")
			fmt.Println("        return 1")
			fmt.Println("    }")
			fmt.Println("    return n * faktorial(n-1) // Recursive Case")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggilnya:")
			fmt.Println("// hasil := faktorial(5) // hasil akan menjadi 120")
			fmt.Println("```")
		case 3:
			fmt.Println("Meskipun rekursi terlihat elegan, ia memiliki beberapa pertimbangan:")
			fmt.Println("\n- **Stack Overflow:** Setiap kali fungsi memanggil dirinya sendiri, sebuah *frame* baru ditambahkan ke *call stack*. Jika rekursi terlalu dalam atau tidak memiliki *base case* yang benar, *stack* bisa meluap (stack overflow), menyebabkan program *crash*.")
			fmt.Println("- **Efisiensi Memori dan Waktu:** Terkadang, solusi rekursif bisa lebih boros memori (karena *stack frame*) dan waktu (overhead pemanggilan fungsi) dibandingkan solusi iteratif. Namun, untuk beberapa masalah, rekursi adalah cara paling intuitif atau bahkan paling efisien.")
			fmt.Println("- **Membaca dan Debugging:** Untuk pemula, memahami alur eksekusi rekursif bisa lebih menantang dibandingkan iterasi. Debugging juga memerlukan pemahaman tentang *call stack*.")
			fmt.Println("\nDalam Go, meskipun rekursi didukung penuh, pengembang sering kali mempertimbangkan tradeoff antara kejelasan kode rekursif dan efisiensi solusi iteratif, terutama untuk masalah yang bisa diselesaikan dengan mudah menggunakan loop.")
			fmt.Println("\n---")

			fmt.Println("### Contoh Penggunaan Rekursif Sederhana")
			fmt.Println("\nMari kita lihat fungsi rekursif untuk menghitung deret Fibonacci:")
			fmt.Println("```go")
			fmt.Println("package main")
			fmt.Println("\nimport \"fmt\"")
			fmt.Println("\n// fibonacci adalah fungsi rekursif untuk menghitung angka ke-n dalam deret Fibonacci.")
			fmt.Println("// Deret Fibonacci: 0, 1, 1, 2, 3, 5, 8, ... (setiap angka adalah jumlah dua angka sebelumnya)")
			fmt.Println("func fibonacci(n int) int {")
			fmt.Println("    if n <= 1 { // Base Case: untuk n=0 atau n=1")
			fmt.Println("        return n")
			fmt.Println("    }")
			fmt.Println("    return fibonacci(n-1) + fibonacci(n-2) // Recursive Case")
			fmt.Println("}")
			fmt.Println("\nfunc main() {")
			fmt.Println("    // Memanggil fungsi fibonacci")
			fmt.Printf("    fmt.Printf(\"Fibonacci ke-7 adalah: %d\\n\", fibonacci(7)) // Output: 13")
			fmt.Printf("    fmt.Printf(\"Fibonacci ke-10 adalah: %d\\n\", fibonacci(10)) // Output: 55")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDalam contoh di atas:")
			fmt.Println("* Fungsi `fibonacci` memanggil dirinya sendiri dua kali dalam *recursive case*.")
			fmt.Println("* *Base case* memastikan rekursi berhenti saat `n` adalah 0 atau 1.")
			fmt.Println("Ini menunjukkan bagaimana rekursi dapat memecah masalah menjadi sub-masalah yang lebih kecil, yang kemudian diselesaikan oleh pemanggilan fungsi yang sama.")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			// Jika halaman tidak valid, kita akan menganggap pengguna ingin kembali ke menu utama.
			return true // Kembali ke MainMenu
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
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
