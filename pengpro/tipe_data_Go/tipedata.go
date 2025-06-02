package tipe_data_go

import (
	"Project-Alpro/atribut"          // Untuk fungsi ClearScreen
	"Project-Alpro/pengpro/quiz" // Untuk mengakses DataQuiz, NMAX, dan fungsi pembantu
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Tipe Data Go.
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
		fmt.Println("1. Tipe Data Go Language") // Mengubah teks menu
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

// Submenu menampilkan submenu untuk materi Tipe Data Go.
func Submenu() { // Tidak ada parameter di sini
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("   SUBMENU MATERI TIPE DATA GO   ") // Mengubah teks menu
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Tipe Data Go Language")   // Mengubah teks menu
		fmt.Println("2. Kerjakan Kuis Tipe Data Go Language") // Mengubah teks menu
		fmt.Println("3. Kembali ke Menu Utama")
		// Opsi "Daftar / Masukkan Data Baru" sudah dipindahkan ke pengpro.Submenu
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriTipeDataGo() { // Mengubah nama fungsi materi
				quiz.HandleQuizStart(QuizTipeDataGo, "Tipe Data Go") // Mengubah nama fungsi kuis dan nama materi
			}
		case "2":
			quiz.HandleQuizStart(QuizTipeDataGo, "Tipe Data Go") // Mengubah nama fungsi kuis dan nama materi
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

// MateriTipeDataGo menampilkan materi Tipe Data Go secara berhalaman.
// Mengembalikan true jika pengguna ingin melanjutkan ke kuis, false jika kembali ke submenu.
func MateriTipeDataGo() bool { // Mengubah nama fungsi dari MateriGoLanguage
	halamanSekarang := 1
	totalHalaman := 3
	var choice string

	readingMaterial := true

	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Tipe Data pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Tipe data adalah klasifikasi yang memberitahu kompiler bagaimana programmer bermaksud untuk menggunakan data. Dalam Go, setiap variabel harus memiliki tipe data.")
			fmt.Println("Go memiliki beberapa tipe data dasar, termasuk:")
			fmt.Println("- Tipe data **Numerik**: Digunakan untuk angka. Contoh: `int`, `float64`.")
			fmt.Println("- Tipe data **Boolean**: Digunakan untuk nilai benar atau salah. Contoh: `bool`.")
		case 2:
			fmt.Println("- Tipe data **String**: Digunakan untuk teks atau urutan karakter. Contoh: `string`.")
			fmt.Println("Mari kita lihat beberapa contoh tipe data numerik:")
			fmt.Println("  - `int`: Bilangan bulat (integer), ukurannya tergantung arsitektur sistem (32-bit atau 64-bit).")
			fmt.Println("  - `int8`, `int16`, `int32`, `int64`: Bilangan bulat dengan ukuran bit yang spesifik.")
			fmt.Println("  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`: Bilangan bulat non-negatif (unsigned integer).")
			fmt.Println("  - `float32`, `float64`: Bilangan desimal (floating-point numbers), `float64` adalah yang paling umum digunakan.")
		case 3:
			fmt.Println("Untuk tipe data **String**, nilai string diapit oleh tanda kutip ganda (\"). Contoh: `var nama string = \"Budi\"`.")
			fmt.Println("Untuk tipe data **Boolean**, hanya ada dua nilai: `true` atau `false`. Contoh: `var isStudent bool = true`.")
			fmt.Println("Memahami tipe data sangat penting untuk menulis kode yang efisien dan bebas error.")
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
		} else { // On the last page of the material
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

// QuizTipeDataGo menjalankan kuis Tipe Data Go dan memperbarui skor siswa.
// Menerima pointer ke array DataQuiz dan indeks siswa yang sedang kuis.
func QuizTipeDataGo(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) { // Mengubah nama fungsi dari QuizGoLanguage
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Tipe Data pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Tipe data apa yang digunakan untuk bilangan bulat dalam Go?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "int") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah int (atau int8, int16, dll.).")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Untuk nilai benar atau salah (true/false), tipe data apa yang digunakan?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "bool" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah bool.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Tipe data apa yang digunakan untuk menyimpan teks dalam Go?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "string" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah string.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	// Perbarui skor siswa di array DataQuiz global melalui pointer
	if index >= 0 && index < quiz.NMAX {
		// Mengakses field skor spesifik untuk Tipe Data Go
		oldScoreTipeDataGo := quizDataArray[index].TipeDataGoScore
		quizDataArray[index].TipeDataGoScore = currentQuizScore
		scoreChange := currentQuizScore - oldScoreTipeDataGo
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Tipe Data Go' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package dataquiz
}
