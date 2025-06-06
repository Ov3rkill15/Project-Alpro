package p_searching

import (
	binary "Project-Alpro/alpro/praktikum/Binary_Search"
	sequential "Project-Alpro/alpro/praktikum/Sequential_Search"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Searching.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
======================================
Selamat Datang di Pembelajaran Searching
======================================
1. Apa itu Sequential Search
2. Apa itu Binary Search
3. Soal Sequential Search
4. Soal Binary Search
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarSequentialSearch akan mengelola navigasi internalnya sendiri
			// dan akan mengembalikan false jika ingin keluar dari program, true jika kembali ke MainMenu.
			shouldStayInMenu := belajarSequentialSearch()
			if !shouldStayInMenu {
				isRunning = false // Hentikan loop MainMenu jika belajarSequentialSearch mengindikasikan keluar
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// belajarBinarySearch akan mengelola navigasi internalnya sendiri
			// dan akan mengembalikan false jika ingin keluar dari program, true jika kembali ke MainMenu.
			shouldStayInMenu := belajarBinarySearch()
			if !shouldStayInMenu {
				isRunning = false // Hentikan loop MainMenu jika belajarBinarySearch mengindikasikan keluar
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "3":
			// Asumsi: sequential.Soal1SeqSearch() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := sequential.Soal1SeqSearch()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran searching.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "4":
			// Asumsi: binary.Soal1BinSearch() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := binary.Soal1BinSearch()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran searching.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran searching.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', '3', '4', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarSequentialSearch mengelola materi pembelajaran Sequential Search.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program.
func belajarSequentialSearch() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var Choice string
	isStudying := true // Kontrol loop utama materi

	for isStudying {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Sequential Search di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan kamu mencari sebuah buku di rak yang tidak berurutan sama sekali. Cara paling dasar adalah kamu mulai dari buku pertama, lalu periksa satu per satu buku berikutnya, hingga kamu menemukan buku yang kamu cari atau sampai semua buku di rak habis.")
			fmt.Println("\nDalam pemrograman, **Sequential Search (Pencarian Sekuensial)**, atau sering disebut **Linear Search**, adalah algoritma pencarian paling sederhana. Ini bekerja dengan memeriksa setiap elemen dalam daftar (array atau slice) secara berurutan, dari awal hingga akhir, sampai elemen yang cocok ditemukan atau seluruh daftar sudah diperiksa.")
			fmt.Println("\n**Kapan Sequential Search Digunakan?**")
			fmt.Println("- **Daftar Kecil:** Sangat cocok untuk daftar dengan jumlah elemen yang sedikit.")
			fmt.Println("- **Daftar Tidak Terurut:** Ini adalah satu-satunya pilihan dasar jika daftar tidak diurutkan, karena algoritma pencarian lain yang lebih cepat (seperti Binary Search) memerlukan daftar yang terurut.")
			fmt.Println("- **Kesederhanaan Implementasi:** Algoritma ini sangat mudah dipahami dan diimplementasikan.")
		case 2:
			fmt.Println("Cara kerja Sequential Search sangat lugas:")
			fmt.Println("1.  Mulai dari elemen pertama daftar.")
			fmt.Println("2.  Bandingkan elemen saat ini dengan nilai yang dicari (target).")
			fmt.Println("3.  Jika cocok, pencarian berhasil dan indeks elemen dikembalikan.")
			fmt.Println("4.  Jika tidak cocok, pindah ke elemen berikutnya dan ulangi langkah 2.")
			fmt.Println("5.  Jika seluruh daftar sudah diperiksa dan target tidak ditemukan, pencarian gagal.")
			fmt.Println("\n**Contoh Implementasi Sequential Search di Go:**")
			fmt.Println("```go")
			fmt.Println("func sequentialSearch(data []int, target int) int {")
			fmt.Println("    for i, element := range data {")
			fmt.Println("        if element == target {")
			fmt.Println("           return i // Element found at index i")
			fmt.Println("        }")
			fmt.Println("    }")
			fmt.Println("    return -1 // Element not found")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggilnya:")
			fmt.Println("// numbers := []int{10, 5, 20, 8, 15}")
			fmt.Println("// index := sequentialSearch(numbers, 8) // index akan menjadi 3")
			fmt.Println("// index = sequentialSearch(numbers, 7) // index akan menjadi -1")
			fmt.Println("```")
		case 3:
			fmt.Println("Meskipun Sequential Search sederhana, ia memiliki kelemahan utama pada daftar yang besar:")
			fmt.Println("\n- **Efisiensi Waktu (Time Complexity):**")
			fmt.Println("    - **Best Case:** O(1) - Jika elemen yang dicari adalah yang pertama.")
			fmt.Println("    - **Worst Case & Average Case:** O(n) - Jika elemen berada di akhir daftar, atau tidak ditemukan sama sekali, algoritma harus memeriksa setiap elemen (n adalah jumlah elemen). Ini berarti waktu yang dibutuhkan berbanding lurus dengan ukuran daftar.")
			fmt.Println("\n- **Tidak Efisien untuk Data Besar:** Untuk daftar yang memiliki ribuan atau jutaan elemen, Sequential Search akan sangat lambat karena harus memeriksa setiap elemen satu per satu. Ini membuatnya tidak praktis untuk database besar atau aplikasi yang membutuhkan kinerja tinggi.")
			fmt.Println("\nOleh karena itu, Sequential Search jarang digunakan untuk pencarian di data besar, kecuali jika data tersebut tidak dapat diurutkan atau jumlahnya sangat kecil.")
			fmt.Println("\n---")
			fmt.Println("### Latihan Sederhana Sequential Search")
			fmt.Println("Pikirkan bagaimana kamu akan memodifikasi fungsi `sequentialSearch` di atas agar juga mengembalikan `true` jika elemen ditemukan, dan `false` jika tidak ditemukan.")
			fmt.Println("```go")
			fmt.Println("func cariElemen(data []string, target string) bool {")
			fmt.Println("    // Tulis kodemu di sini")
			fmt.Println("    // Hint: Gunakan loop dan perbandingan string")
			fmt.Println("    return false // Ganti ini dengan logika pencarianmu")
			fmt.Println("}")
			fmt.Println("```")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			return true // Kembali ke MainMenu jika ada kesalahan halaman
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
				return true        // Kembali ke MainMenu
			} else if strings.ToLower(Choice) == "s" && halamanSekarang == totalHalaman {
				// Handle redirection to sequential search quiz directly from here if 's' is chosen
				// This assumes sequential.Soal1SeqSearch() will return control to MainMenu or exit.
				atribut.ClearScreen()
				shouldStayInMenu := sequential.Soal1SeqSearch()
				if !shouldStayInMenu {
					return false // Exit program
				}
				isStudying = false // Stop studying loop to prevent re-displaying this section
				inputValid = true  // Exit input loop
				return true        // Return to MainMenu after quiz
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y', 'n', atau 's' (jika di halaman terakhir).")
			}
		}
	}
	return true // Kembali ke MainMenu setelah selesai belajar
}

// belajarBinarySearch mengelola materi pembelajaran Binary Search.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program.
func belajarBinarySearch() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var Choice string
	isStudying := true // Kontrol loop utama materi

	for isStudying {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Binary Search di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan kamu mencari kata di kamus. Kamu tidak akan mencari dari halaman pertama sampai terakhir. Sebaliknya, kamu akan membuka bagian tengah kamus, lihat apakah kata yang kamu cari ada di bagian kiri atau kanan, lalu ulangi proses tersebut di bagian yang lebih kecil. Ini jauh lebih cepat!")
			fmt.Println("\nDalam pemrograman, **Binary Search (Pencarian Biner)** adalah algoritma pencarian yang sangat efisien untuk menemukan elemen dalam daftar (array atau slice) **yang sudah terurut**. Ini bekerja dengan membagi daftar menjadi dua bagian secara berulang.")
			fmt.Println("\n**Kapan Binary Search Digunakan?**")
			fmt.Println("- **Daftar Terurut:** Ini adalah prasyarat mutlak. Binary Search TIDAK akan berfungsi dengan benar pada daftar yang tidak terurut.")
			fmt.Println("- **Daftar Besar:** Sangat efisien untuk mencari elemen di daftar yang sangat besar, seperti di database atau struktur data besar.")
			fmt.Println("- **Kinerja Tinggi:** Memberikan kinerja pencarian yang jauh lebih cepat daripada Sequential Search untuk data terurut yang besar.")
		case 2:
			fmt.Println("Cara kerja Binary Search melibatkan pembagian berulang:")
			fmt.Println("1.  Tentukan titik tengah daftar.")
			fmt.Println("2.  Bandingkan nilai target dengan elemen di titik tengah.")
			fmt.Println("3.  - Jika cocok, elemen ditemukan.")
			fmt.Println("    - Jika target lebih kecil dari elemen tengah, fokus pencarian di paruh kiri daftar.")
			fmt.Println("    - Jika target lebih besar dari elemen tengah, fokus pencarian di paruh kanan daftar.")
			fmt.Println("4.  Ulangi langkah 1-3 pada paruh yang dipilih, terus membagi daftar menjadi dua hingga elemen ditemukan atau daftar menjadi kosong.")
			fmt.Println("\n**Contoh Implementasi Binary Search di Go:**")
			fmt.Println("```go")
			fmt.Println("func binarySearch(data []int, target int) int {")
			fmt.Println("    low := 0")
			fmt.Println("    high := len(data) - 1")
			fmt.Println("\n    for low <= high {")
			fmt.Println("        mid := low + (high-low)/2 // Mencegah overflow pada bilangan besar")
			fmt.Println("\n        if data[mid] == target {")
			fmt.Println("            return mid // Element found")
			fmt.Println("        } else if data[mid] < target {")
			fmt.Println("            low = mid + 1 // Search in the right half")
			fmt.Println("        } else {")
			fmt.Println("            high = mid - 1 // Search in the left half")
			fmt.Println("        }")
			fmt.Println("    }")
			fmt.Println("    return -1 // Element not found")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggilnya:")
			fmt.Println("// sortedNumbers := []int{5, 8, 10, 15, 20}")
			fmt.Println("// index := binarySearch(sortedNumbers, 10) // index akan menjadi 2")
			fmt.Println("// index = binarySearch(sortedNumbers, 7) // index akan menjadi -1")
			fmt.Println("```")
		case 3:
			fmt.Println("Binary Search adalah algoritma yang sangat efisien, tetapi memiliki satu syarat utama:")
			fmt.Println("\n- **Efisiensi Waktu (Time Complexity):**")
			fmt.Println("    - **Best Case:** O(1) - Jika elemen yang dicari adalah elemen tengah.")
			fmt.Println("    - **Worst Case & Average Case:** O(log n) - Waktu yang dibutuhkan meningkat secara logaritmis terhadap ukuran daftar. Ini jauh lebih cepat daripada O(n) untuk daftar besar. Misalnya, untuk 1 juta elemen, log2(1.000.000) adalah sekitar 20 langkah, bandingkan dengan 1 juta langkah pada Sequential Search.")
			fmt.Println("\n- **Prasyarat Daftar Terurut:** Keuntungan efisiensi ini hanya bisa didapatkan jika daftar sudah diurutkan. Jika daftar tidak terurut, kamu perlu mengurutkannya terlebih dahulu (yang memiliki overhead waktu tersendiri, biasanya O(n log n) atau lebih buruk).")
			fmt.Println("\nBinary Search adalah algoritma pilihan untuk pencarian di kumpulan data besar yang terurut, seperti dalam database, *dictionary*, atau sistem pencarian indeks.")
			fmt.Println("\n---")
			fmt.Println("### Latihan Sederhana Binary Search")
			fmt.Println("Pikirkan bagaimana kamu akan memodifikasi fungsi `binarySearch` di atas agar juga mengembalikan `true` jika elemen ditemukan, dan `false` jika tidak ditemukan.")
			fmt.Println("```go")
			fmt.Println("func apakahAda(data []int, target int) bool {")
			fmt.Println("    // Tulis kodemu di sini")
			fmt.Println("    // Hint: Pastikan data sudah terurut!")
			fmt.Println("    return false // Ganti ini dengan logika pencarianmu")
			fmt.Println("}")
			fmt.Println("```")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			return true // Kembali ke MainMenu jika ada kesalahan halaman
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
				return true        // Kembali ke MainMenu
			} else if strings.ToLower(Choice) == "s" && halamanSekarang == totalHalaman {
				// Handle redirection to binary search quiz directly from here if 's' is chosen
				atribut.ClearScreen()
				shouldStayInMenu := binary.Soal1BinSearch()
				if !shouldStayInMenu {
					return false // Exit program
				}
				isStudying = false // Stop studying loop to prevent re-displaying this section
				inputValid = true  // Exit input loop
				return true        // Return to MainMenu after quiz
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y', 'n', atau 's' (jika di halaman terakhir).")
			}
		}
	}
	return true // Kembali ke MainMenu setelah selesai belajar
}
