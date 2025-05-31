package p_searching 

import (
	"Project-Alpro/atribut" 
	"fmt"
	"strings"
)


func MainMenu() {
	var pilihan string
	var pilihan2 string 

	fmt.Println(`
======================================
Selamat Datang di Pembelajaran Searching
======================================
1. Apa itu Sequential Search
2. Apa itu Binary Search
3. Soal Searching
4. Keluar
    `)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Scanln() 
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarSequentialSearch() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			isInputValid := false
			for !isInputValid {
				fmt.Scan(&pilihan2)
				fmt.Scanln() 
				if strings.ToLower(pilihan2) == "y" {
					isInputValid = true 
					MainMenu()          
					return              
				} else if strings.ToLower(pilihan2) == "n" {
					isInputValid = true 
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
		if belajarBinarySearch() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			isInputValid := false
			for !isInputValid {
				fmt.Scan(&pilihan2)
				fmt.Scanln() 
				if strings.ToLower(pilihan2) == "y" {
					isInputValid = true 
					MainMenu()         
					return             
				} else if strings.ToLower(pilihan2) == "n" {
					isInputValid = true
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
					fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
				}
			}
		} else { 
			return
		}
	case "3": 
		atribut.ClearScreen()
		soalFunction() 
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		isInputValid := false
		for !isInputValid {
			fmt.Scan(&pilihan2)
			fmt.Scanln() 
			if strings.ToLower(pilihan2) == "y" {
				isInputValid = true 
				MainMenu()          
				return           
			} else if strings.ToLower(pilihan2) == "n" {
				isInputValid = true 
				return
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
			}
		}
	case "4": 
		fmt.Println("Terima kasih telah menggunakan program pembelajaran searching.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', '3', atau '4'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln() 
		MainMenu()  
	}
}

func belajarSequentialSearch() bool {
	halamanSekarang := 1
	totalHalaman := 3        
	var berhenti bool = true 
	var Choice string       

	for berhenti { 
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
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			isInputValid := false
			for !isInputValid {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() 

				if strings.ToLower(Choice) == "y" {
					halamanSekarang++
					isInputValid = true 
				} else if strings.ToLower(Choice) == "n" {
					atribut.ClearScreen()
					berhenti = false    
					isInputValid = true 
					MainMenu()          
					return false        
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else { 
			isInputValid := false
			for !isInputValid {
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln()

				if strings.ToLower(Choice) == "s" {
					soalFunction()      
					berhenti = false    
					isInputValid = true 
					return true         
				} else if strings.ToLower(Choice) == "n" {
					berhenti = false    
					isInputValid = true 
					MainMenu()          
					atribut.ClearScreen()
					return false 
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 's' atau 'n'.")
				}
			}
		}
	}
	return true // Default return: if loop finishes (e.g., all pages shown and user chose 's' or 'n')
}

// belajarBinarySearch manages the display of Binary Search material.
// It returns true if the user wants to return to the MainMenu, false if it handled redirection.
func belajarBinarySearch() bool {
	halamanSekarang := 1
	totalHalaman := 3        // Adjusted total pages for Binary Search
	var berhenti bool = true // Inisialisasi 'berhenti' sebagai true
	var Choice string        // Local variable for Choice

	for berhenti { // Loop dikendalikan oleh variabel 'berhenti'
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
			fmt.Println("\n---") // Garis pemisah untuk visualisasi
			fmt.Println("### Latihan Sederhana Binary Search")
			fmt.Println("Pikirkan bagaimana kamu akan memodifikasi fungsi `binarySearch` di atas agar juga mengembalikan `true` jika elemen ditemukan, dan `false` jika tidak ditemukan.")
			fmt.Println("```go")
			fmt.Println("func apakahAda(data []int, target int) bool {")
			fmt.Println("    // Tulis kodemu di sini")
			fmt.Println("    // Hint: Pastikan data sudah terurut!")
			fmt.Println("    return false // Ganti ini dengan logika pencarianmu")
			fmt.Println("}")
			fmt.Println("```")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			isInputValid := false
			for !isInputValid {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() // Clear newline

				if strings.ToLower(Choice) == "y" {
					halamanSekarang++
					isInputValid = true // Input valid, keluar dari loop
				} else if strings.ToLower(Choice) == "n" {
					atribut.ClearScreen()
					berhenti = false    // Mengatur 'berhenti' menjadi false untuk keluar dari loop utama
					isInputValid = true // Input valid, keluar dari loop
					MainMenu()          // Recursive call to MainMenu
					return false        // Return false untuk mengindikasikan MainMenu sudah dipanggil
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else { // On the last page
			isInputValid := false
			for !isInputValid {
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln() // Clear newline

				if strings.ToLower(Choice) == "s" {
					soalFunction()      // Memanggil fungsi soalFunction
					berhenti = false    // Mengatur 'berhenti' menjadi false untuk keluar dari loop utama
					isInputValid = true // Input valid, keluar dari loop
					return true         // Return true to go back to MainMenu after quiz
				} else if strings.ToLower(Choice) == "n" {
					berhenti = false    // Mengatur 'berhenti' menjadi false untuk keluar dari loop utama
					isInputValid = true // Input valid, keluar dari loop
					MainMenu()          // Rekursi
					atribut.ClearScreen()
					return false // Return false untuk mengindikasikan MainMenu sudah dipanggil
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 's' atau 'n'.")
				}
			}
		}
	}
	return true // Default return: if loop finishes (e.g., all pages shown and user chose 's' or 'n')
}

// soalFunction is a placeholder for the quiz section.
// This function can be reused by both sequential and binary search learning modules.
func soalFunction() {
	atribut.ClearScreen()
	fmt.Println("==================")
	fmt.Println(" 	SOAL SEARCHING") // Generic name for searching quiz
	fmt.Println("==================")
	fmt.Println("COMING SOON!")
	fmt.Println("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln() // Wait for Enter
}
