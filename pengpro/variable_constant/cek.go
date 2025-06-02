package variable_constant

import (
	"Project-Alpro/atribut"          // Untuk fungsi ClearScreen
	"Project-Alpro/pengpro/quiz" // Untuk mengakses DataQuiz, NMAX, dan fungsi pembantu
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Variabel dan Konstanta.
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
		fmt.Println("1. Variabel dan Konstanta Go") // Mengubah teks menu
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

// Submenu menampilkan submenu untuk materi Variabel dan Konstanta.
func Submenu() { // Tidak ada parameter di sini
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("   SUBMENU MATERI VARIABEL DAN KONSTANTA") // Mengubah teks menu
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Variabel dan Konstanta Go")   // Mengubah teks menu
		fmt.Println("2. Kerjakan Kuis Variabel dan Konstanta Go") // Mengubah teks menu
		fmt.Println("3. Kembali ke Menu Utama")
		// Opsi "Daftar / Masukkan Data Baru" sudah dipindahkan ke pengpro.Submenu
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriVariableConstant() { // Mengubah nama fungsi materi
				quiz.HandleQuizStart(QuizVariableConstant, "Variabel dan Konstanta") // Mengubah nama fungsi kuis dan nama materi
			}
		case "2":
			quiz.HandleQuizStart(QuizVariableConstant, "Variabel dan Konstanta") // Mengubah nama fungsi kuis dan nama materi
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

// MateriVariableConstant menampilkan materi Variabel dan Konstanta secara berhalaman.
// Mengembalikan true jika pengguna ingin melanjutkan ke kuis, false jika kembali ke submenu.
func MateriVariableConstant() bool { // Mengubah nama fungsi dari MateriGoLanguage
	halamanSekarang := 1
	totalHalaman := 3
	var choice string

	readingMaterial := true

	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Variabel dan Konstanta pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam pemrograman Go, **variabel** digunakan untuk menyimpan nilai data yang dapat berubah selama program berjalan.")
			fmt.Println("Anda mendeklarasikan variabel dengan kata kunci `var`, diikuti nama variabel, tipe data, dan nilai opsional.")
			fmt.Println("Contoh deklarasi variabel:")
			fmt.Println("  `var nama string = \"Alice\"`")
			fmt.Println("  `var umur int = 30`")
			fmt.Println("Go juga memiliki cara singkat untuk mendeklarasikan dan menginisialisasi variabel menggunakan `:=` (disebut *short declaration operator*).")
			fmt.Println("Contoh short declaration:")
			fmt.Println("  `kota := \"New York\"`")
			fmt.Println("  `harga := 99.99`")
		case 2:
			fmt.Println("Variabel yang dideklarasikan tetapi tidak diinisialisasi akan secara otomatis diberi nilai nol (zero value) sesuai tipe datanya.")
			fmt.Println("Contoh zero value:")
			fmt.Println("  - `int`: 0")
			fmt.Println("  - `float`: 0.0")
			fmt.Println("  - `bool`: `false`")
			fmt.Println("  - `string`: `\"\"` (string kosong)")
			fmt.Println("\nSekarang mari kita bahas **konstanta**.")
			fmt.Println("Konstanta adalah nilai yang tidak dapat diubah setelah dideklarasikan. Konstanta dideklarasikan dengan kata kunci `const`.")
			fmt.Println("Konstanta harus diinisialisasi saat deklarasi.")
		case 3:
			fmt.Println("Contoh deklarasi konstanta:")
			fmt.Println("  `const PI = 3.14`")
			fmt.Println("  `const BAHASA = \"GoLang\"`")
			fmt.Println("Perbedaan utama antara variabel dan konstanta adalah bahwa nilai variabel dapat diubah, sedangkan nilai konstanta tidak bisa.")
			fmt.Println("Konstanta dalam Go dapat berupa karakter, string, boolean, atau nilai numerik.")
			fmt.Println("Menggunakan konstanta membantu membuat kode lebih aman dan mudah dibaca untuk nilai-nilai yang tetap.")
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

// QuizVariableConstant menjalankan kuis Variabel dan Konstanta dan memperbarui skor siswa.
// Menerima pointer ke array DataQuiz dan indeks siswa yang sedang kuis.
func QuizVariableConstant(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) { // Mengubah nama fungsi dari QuizGoLanguage
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Variabel dan Konstanta pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Kata kunci apa yang digunakan untuk mendeklarasikan variabel di Go?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "var" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'var'.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Operator apa yang digunakan untuk deklarasi variabel singkat (short declaration) di Go?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if jawaban == ":=" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah ':='. ")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Nilai yang tidak dapat diubah setelah dideklarasikan disebut apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "konstanta") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah konstanta.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	// Perbarui skor siswa di array DataQuiz global melalui pointer
	if index >= 0 && index < quiz.NMAX {
		// Mengakses field skor spesifik untuk Variabel dan Konstanta
		oldScoreVariableConstant := quizDataArray[index].VariableConstantScore
		quizDataArray[index].VariableConstantScore = currentQuizScore
		scoreChange := currentQuizScore - oldScoreVariableConstant
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Variabel dan Konstanta' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package dataquiz
}
