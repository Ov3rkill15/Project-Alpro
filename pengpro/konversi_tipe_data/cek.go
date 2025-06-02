package konversi_tipe_data

import (
	"Project-Alpro/atribut"          // Untuk fungsi ClearScreen
	"Project-Alpro/pengpro/quiz" // Untuk mengakses DataQuiz, NMAX, dan fungsi pembantu
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Konversi Tipe Data.
// Menerima pointer ke array DataQuiz untuk mengelola data siswa.
func MainMenu(quizDataArray *[quiz.NMAX]atribut.Quiz) { // Menambahkan parameter quizDataArray
	var choice string
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SELAMAT DATANG!           ")
		fmt.Println("====================================")
		fmt.Println("Pilih materi yang ingin Anda pelajari:")
		fmt.Println("1. Konversi Tipe Data Go") // Mengubah teks menu
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			Submenu() // Panggilan Submenu tanpa parameter
		case "2":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Submenu menampilkan submenu untuk materi Konversi Tipe Data.
func Submenu() { // Tidak ada parameter di sini
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("   SUBMENU MATERI KONVERSI TIPE DATA") // Mengubah teks menu
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Konversi Tipe Data Go")   // Mengubah teks menu
		fmt.Println("2. Kerjakan Kuis Konversi Tipe Data Go") // Mengubah teks menu
		fmt.Println("3. Kembali ke Menu Utama")
		// Opsi "Daftar / Masukkan Data Baru" dipindahkan ke pengpro.Submenu
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriKonversiTipeData() { // Mengubah nama fungsi materi
				quiz.HandleQuizStart(QuizKonversiTipeData, "Konversi Tipe Data") // Mengubah nama fungsi kuis dan nama materi
			}
		case "2":
			quiz.HandleQuizStart(QuizKonversiTipeData, "Konversi Tipe Data") // Mengubah nama fungsi kuis dan nama materi
		case "3":
			fmt.Println("Kembali ke menu utama...")
			berhenti2 = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// PromptKembaliToMain DIHAPUS dari sini karena sudah ada di package quiz
// func PromptKembaliToMain() { ... }

// MateriKonversiTipeData menampilkan materi Konversi Tipe Data secara berhalaman.
// Mengembalikan true jika pengguna ingin melanjutkan ke kuis, false jika kembali ke submenu.
func MateriKonversiTipeData() bool { // Mengubah nama fungsi dari MateriGoLanguage
	halamanSekarang := 1
	totalHalaman := 3
	var choice string

	readingMaterial := true

	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Konversi Tipe Data pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, **konversi tipe data** (type conversion) adalah proses mengubah nilai dari satu tipe data ke tipe data lain.")
			fmt.Println("Go adalah bahasa yang *strongly typed*, artinya Anda harus secara eksplisit mengonversi tipe data jika ingin menggunakannya dalam operasi yang melibatkan tipe yang berbeda.")
			fmt.Println("Konversi tipe data dasar dilakukan dengan menulis nama tipe target dalam tanda kurung, diikuti dengan variabel atau nilai yang ingin dikonversi.")
			fmt.Println("Contoh konversi dari `int` ke `float64`:")
			fmt.Println("  `var angkaInt int = 10`")
			fmt.Println("  `var angkaFloat float64 = float64(angkaInt)`")
			fmt.Println("  `fmt.Printf(\"Angka Int: %T %v\\n\", angkaInt, angkaInt)`")
			fmt.Println("  `fmt.Printf(\"Angka Float: %T %v\\n\", angkaFloat, angkaFloat)`")
		case 2:
			fmt.Println("Perhatikan bahwa saat mengonversi tipe data, ada kemungkinan **kehilangan presisi** atau data jika tipe target tidak dapat menampung nilai dari tipe asalnya.")
			fmt.Println("Contoh kehilangan presisi dari `float64` ke `int`:")
			fmt.Println("  `var nilaiFloat float64 = 10.75`")
			fmt.Println("  `var nilaiInt int = int(nilaiFloat)`")
			fmt.Println("  `fmt.Printf(\"Nilai Float: %T %v\\n\", nilaiFloat, nilaiFloat)`")
			fmt.Println("  `fmt.Printf(\"Nilai Int: %T %v\\n\", nilaiInt, nilaiInt)`")
			fmt.Println("\nKonversi tipe data numerik yang lebih besar ke yang lebih kecil juga dapat menyebabkan *overflow* atau nilai yang tidak terduga jika nilainya melebihi kapasitas tipe target.")
			fmt.Println("Contoh konversi dari `int32` ke `int8` (hati-hati jika nilai terlalu besar):")
			fmt.Println("  `var besarAngka int32 = 130`")
			fmt.Println("  `var kecilAngka int8 = int8(besarAngka)`")
			fmt.Println("  `fmt.Printf(\"Besar Angka: %T %v\\n\", besarAngka, besarAngka)`")
			fmt.Println("  `fmt.Printf(\"Kecil Angka: %T %v\\n\", kecilAngka, kecilAngka)`")
		case 3:
			fmt.Println("Konversi juga bisa dilakukan pada tipe **boolean** ke numerik, meskipun hasilnya selalu 0 atau 1 (untuk `false` dan `true`).")
			fmt.Println("Contoh konversi `bool` ke `int`:")
			fmt.Println("  `var isTrue bool = true`")
			fmt.Println("  `var intFromBool int = int(isTrue)` // Tidak diizinkan langsung di Go, ini hanya konsep.")
			fmt.Println("  `// Di Go, Anda perlu logika if/else atau operator ternary (jika ada) untuk konversi ini.`")
			fmt.Println("  `// Contoh: intFromBool := 0; if isTrue { intFromBool = 1 }`")
			fmt.Println("\nPenting untuk selalu berhati-hati saat melakukan konversi tipe data, terutama jika ada potensi kehilangan informasi.")
			fmt.Println("Pahami implikasi dari setiap konversi agar program berjalan sesuai harapan.")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			validInput := false
			for !validInput {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke submenu (n)? ")
				fmt.Scanln(&choice)

				if strings.ToLower(choice) == "y" {
					halamanSekarang++
					validInput = true
				} else if strings.ToLower(choice) == "n" {
					readingMaterial = false
					return false
				} else {
					fmt.Println("Pilihan tidak valid. Silakan masukkan 'y' atau 'n'.")
				}
			}
		} else { // Di halaman terakhir materi
			validInput := false
			for !validInput {
				fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke submenu (n)? ")
				fmt.Scanln(&choice)

				if strings.ToLower(choice) == "k" {
					readingMaterial = false
					return true
				} else if strings.ToLower(choice) == "n" {
					readingMaterial = false
					return false
				} else {
					fmt.Println("Pilihan tidak valid. Silakan masukkan 'k' atau 'n'.")
				}
			}
		}
	}
	return false
}

// QuizKonversiTipeData menjalankan kuis Konversi Tipe Data dan memperbarui skor siswa.
// Menerima pointer ke array DataQuiz dan indeks siswa yang sedang kuis.
func QuizKonversiTipeData(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) { // Mengubah nama fungsi dari QuizGoLanguage
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Konversi Tipe Data pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1 (Updated for Type Conversion without strconv)
	fmt.Println("\n1. Dalam Go, apakah konversi tipe data harus dilakukan secara eksplisit? (Ya/Tidak)")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "ya" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'Ya'.")
	}

	// Pertanyaan 2 (Updated for Type Conversion without strconv)
	fmt.Println("\n2. Apa risiko utama saat mengonversi tipe data `float64` ke `int`?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "kehilangan presisi") || strings.Contains(strings.ToLower(jawaban), "data desimal hilang") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah kehilangan presisi (data desimal hilang).")
	}

	// Pertanyaan 3 (Updated for Type Conversion without strconv)
	fmt.Println("\n3. Jika Anda mengonversi tipe data numerik yang besar ke tipe data numerik yang lebih kecil, apa yang mungkin terjadi pada nilainya?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "overflow") || strings.Contains(strings.ToLower(jawaban), "nilai tidak terduga") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah overflow atau nilai tidak terduga.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	// Perbarui skor siswa di array DataQuiz global melalui pointer
	if index >= 0 && index < quiz.NMAX {
		// Mengakses field skor spesifik untuk Konversi Tipe Data
		oldScoreKonversiTipeData := quizDataArray[index].KonversiTipeDataScore
		quizDataArray[index].KonversiTipeDataScore = currentQuizScore
		scoreChange := currentQuizScore - oldScoreKonversiTipeData
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Konversi Tipe Data' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package quiz
}

