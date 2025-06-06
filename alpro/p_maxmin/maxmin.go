package p_maxmin

import (
	soal "Project-Alpro/alpro/praktikum/Searching" // Asumsi ini adalah package soal untuk Searching (Max/Min)
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Max & Min Array.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
============================================
Selamat Datang di Pembelajaran Max & Min Array
============================================
1. Mencari Nilai Maksimum & Minimum
2. Referensi Soal Max
3. Referensi Soal Min
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarMaxMin akan mengelola navigasi internalnya sendiri
			// dan akan kembali ke MainMenu saat selesai atau pengguna memilih 'n' dari dalamnya.
			shouldStayInMenu := belajarMaxMin()
			if !shouldStayInMenu {
				// Jika belajarMaxMin mengindikasikan keluar dari program, hentikan loop MainMenu.
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// Asumsi: soal.Soal1searching() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal1searching() // Panggil fungsi soal dari package 'soal'
			if !shouldStayInMenu {
				// Jika soal.Soal1searching() mengindikasikan keluar dari program, hentikan loop MainMenu.
				fmt.Println("Terima kasih telah menggunakan program pembelajaran Max & Min.")
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "3":
			// Asumsi: soal.Soal2searching() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal2searching() // Panggil fungsi soal dari package 'soal'
			if !shouldStayInMenu {
				// Jika soal.Soal2searching() mengindikasikan keluar dari program, hentikan loop MainMenu.
				fmt.Println("Terima kasih telah menggunakan program pembelajaran Max & Min.")
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran Max & Min.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', '3', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarMaxMin mengelola materi pembelajaran mencari nilai maksimum dan minimum.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program (misalnya dari halaman terakhir).
func belajarMaxMin() bool {
	halamanSekarang := 1
	totalHalaman := 3 // Ada 3 halaman materi
	var Choice string

	// isStudying akan menjadi false jika pengguna memilih 'n' untuk kembali ke menu utama.
	isStudying := true

	for isStudying { // Loop tak terbatas sampai pengguna memutuskan untuk keluar dari materi ini
		atribut.ClearScreen()
		fmt.Printf("=== Mencari Nilai Maksimum & Minimum di Array (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Mencari nilai maksimum (terbesar) atau minimum (terkecil) dalam sebuah kumpulan data adalah operasi dasar yang sering dilakukan. Dengan menggunakan fungsi, kita bisa mengemas logika pencarian ini menjadi unit yang dapat digunakan kembali dan lebih rapi.")
			fmt.Println("\n**Mengapa Menggunakan Fungsi untuk Max/Min?**")
			fmt.Println("- **Reusabilitas:** Tulis kode pencarian max/min sekali, lalu gunakan berulang kali untuk berbagai array.")
			fmt.Println("- **Modularitas:** Pisahkan logika pencarian dari logika utama program, membuat kode lebih mudah dipahami.")
			fmt.Println("- **Keterbacaan:** Nama fungsi yang jelas (misalnya `findMax`, `findMin`) langsung menunjukkan tujuannya.")
			fmt.Println("\n**Konsep Dasar Implementasi dalam Fungsi:**")
			fmt.Println("1.  Fungsi akan menerima array (atau slice) sebagai input (parameter).")
			fmt.Println("2.  Fungsi akan mengembalikan nilai maksimum atau minimum (atau keduanya) sebagai output.")
			fmt.Println("3.  Logika pencarian akan dilakukan di dalam fungsi.")
		case 2:
			fmt.Println("Berikut adalah contoh implementasi fungsi `FindMax` untuk mencari nilai maksimum dalam slice integer di Go:")
			fmt.Println("\n```go")
			fmt.Println("func FindMax(numbers []int) int {")
			fmt.Println("    if len(numbers) == 0 {")
			fmt.Println("        // Handle case of empty slice. For simplicity, return 0 or a sentinel value.")
			fmt.Println("        return 0 // Atau bisa juga panic(\"slice is empty\")")
			fmt.Println("    }")
			fmt.Println("\n    maxVal := numbers[0] // Asumsikan elemen pertama adalah nilai maksimum sementara")
			fmt.Println("    for i := 1; i < len(numbers); i++ {")
			fmt.Println("        if numbers[i] > maxVal {")
			fmt.Println("            maxVal = numbers[i] // Update jika ditemukan nilai yang lebih besar")
			fmt.Println("        }")
			fmt.Println("    }")
			fmt.Println("    return maxVal")
			fmt.Println("}")
			fmt.Println("\n// Contoh penggunaan:")
			fmt.Println("// data := []int{15, 2, 30, 8, 25}")
			fmt.Println("// maxResult := FindMax(data) // maxResult akan menjadi 30")
			fmt.Println("```")
			fmt.Println("\nDan untuk fungsi `FindMin` (mencari nilai minimum):")
			fmt.Println("\n```go")
			fmt.Println("func FindMin(numbers []int) int {")
			fmt.Println("    if len(numbers) == 0 {")
			fmt.Println("        return 0")
			fmt.Println("    }")
			fmt.Println("\n    minVal := numbers[0] // Asumsikan elemen pertama adalah nilai minimum sementara")
			fmt.Println("    for i := 1; i < len(numbers); i++ {")
			fmt.Println("        if numbers[i] < minVal {")
			fmt.Println("            minVal = numbers[i] // Update jika ditemukan nilai yang lebih kecil")
			fmt.Println("        }")
			fmt.Println("    }")
			fmt.Println("    return minVal")
			fmt.Println("}")
			fmt.Println("\n// Contoh penggunaan:")
			fmt.Println("// data := []int{15, 2, 30, 8, 25}")
			fmt.Println("// minResult := FindMin(data) // minResult akan menjadi 2")
			fmt.Println("```")
		case 3:
			fmt.Println("Seringkali, kita perlu mencari nilai maksimum dan minimum secara bersamaan dari sebuah array. Kita bisa melakukannya dalam satu fungsi untuk efisiensi.")
			fmt.Println("\n**Implementasi Fungsi `FindMaxMin`:**")
			fmt.Println("```go")
			fmt.Println("func FindMaxMin(numbers []int) (int, int) {")
			fmt.Println("    if len(numbers) == 0 {")
			fmt.Println("        return 0, 0 // Atau handle error")
			fmt.Println("    }")
			fmt.Println("\n    maxVal := numbers[0]")
			fmt.Println("    minVal := numbers[0]")
			fmt.Println("\n    for i := 1; i < len(numbers); i++ {")
			fmt.Println("        if numbers[i] > maxVal {")
			fmt.Println("            maxVal = numbers[i]")
			fmt.Println("        } else if numbers[i] < minVal {")
			fmt.Println("            minVal = numbers[i]")
			fmt.Println("        }")
			fmt.Println("    }")
			fmt.Println("    return maxVal, minVal")
			fmt.Println("}")
			fmt.Println("\n// Contoh penggunaan:")
			fmt.Println("// data := []int{15, 2, 30, 8, 25}")
			fmt.Println("// max, min := FindMaxMin(data) // max=30, min=2")
			fmt.Println("```")
			fmt.Println("\nDengan menggabungkannya dalam satu fungsi, kita hanya perlu melakukan satu iterasi penuh melalui array, yang lebih efisien daripada dua iterasi terpisah untuk mencari max dan min secara terpisah.")
			fmt.Println("\n---")
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
