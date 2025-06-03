package p_maxmin

import (
	soal "Project-Alpro/alpro/praktikum/Searching"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

func MainMenu() {
	var pilihan string
	var pilihan2 string

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
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarMaxMin() {
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
		} else {
			return
		}
	case "2":
		atribut.ClearScreen()
		if soal.Soal1searching() {
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
	case "3":
		atribut.ClearScreen()
		if soal.Soal2searching() {
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
		fmt.Println("Terima kasih telah menggunakan program pembelajaran Max & Min.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
		MainMenu()
	}
}

func belajarMaxMin() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true
	var Choice string
	for berhenti {
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
			fmt.Println("### Latihan: Pikirkan Edge Cases")
			fmt.Println("Bagaimana fungsi `FindMaxMin` akan berperilaku jika array hanya memiliki satu elemen? Apakah logikanya masih benar?")
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
				fmt.Print("kembali ke menu utama (n)")
				fmt.Scan(&Choice)
				fmt.Scanln()
				if strings.ToLower(Choice) == "n" {
					berhenti = false
					berhenti2 = true
					MainMenu()
					atribut.ClearScreen()
					return false
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		}
	}
	return true
}
