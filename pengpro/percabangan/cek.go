package percabangan

import (
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro/quiz" // Import dataquiz package
	"fmt"
	"strings"
)

// MainMenu for Percabangan
func MainMenu(quizDataArray *[quiz.NMAX]atribut.Quiz) { // Tetap menerima array
	var choice string
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SELAMAT DATANG!           ")
		fmt.Println("====================================")
		fmt.Println("Pilih materi yang ingin Anda pelajari:")
		fmt.Println("1. Percabangan Go")
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			Submenu() // Tidak perlu meneruskan quizDataArray ke Submenu lagi
		case "2":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Submenu for Percabangan
func Submenu() { // Tidak lagi menerima quizDataArray
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("       SUBMENU MATERI PERCABANGAN   ")
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Percabangan Go")
		fmt.Println("2. Kerjakan Kuis Percabangan Go")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Println("4. Daftar / Masukkan Data Baru")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriPercabanganGo() {
				quiz.HandleQuizStart(QuizPercabanganGo, "Percabangan")
			}
		case "2":
			quiz.HandleQuizStart(QuizPercabanganGo, "Percabangan")
		case "3":
			fmt.Println("Kembali ke menu utama...")
			berhenti2 = false
		case "4":
			quiz.HandleNewStudentRegistration()
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// MateriPercabanganGo (tidak berubah)
func MateriPercabanganGo() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var choice string
	readingMaterial := true
	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Percabangan `if` pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)
		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, **percabangan `if`** digunakan untuk mengeksekusi blok kode tertentu berdasarkan suatu kondisi. Ini adalah salah satu struktur kontrol paling fundamental.")
			fmt.Println("Sintaks dasar `if` adalah:")
			fmt.Println("```go")
			fmt.Println("if kondisi {")
			fmt.Println("    // kode yang dieksekusi jika kondisi true")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("Contoh sederhana:")
			fmt.Println("```go")
			fmt.Println("nilai := 80")
			fmt.Println("if nilai > 70 {")
			fmt.Println("    fmt.Println(\"Selamat, Anda lulus!\")")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("Kondisi harus berupa ekspresi boolean (`true` atau `false`).")
		case 2:
			fmt.Println("Anda dapat menambahkan blok `else` untuk mengeksekusi kode jika kondisi `if` adalah `false`:")
			fmt.Println("```go")
			fmt.Println("umur := 17")
			fmt.Println("if umur >= 18 {")
			fmt.Println("    fmt.Println(\"Anda dewasa.\")")
			fmt.Println("} else {")
			fmt.Println("    fmt.Println(\"Anda belum dewasa.\")")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nUntuk multiple kondisi, gunakan `else if`:")
			fmt.Println("```go")
			fmt.Println("suhu := 25")
			fmt.Println("if suhu < 0 {")
			fmt.Println("    fmt.Println(\"Sangat dingin.\")")
			fmt.Println("} else if suhu >= 0 && suhu <= 20")
			fmt.Println("    fmt.Println(\"Dingin.\")")
			fmt.Println(" else {")
			fmt.Println("    fmt.Println(\"Hangat.\")")
			fmt.Println("}")
			fmt.Println("```")
		case 3:
			fmt.Println("Go juga memungkinkan **`if` dengan *short statement***. Ini berguna untuk mendeklarasikan variabel dan menguji kondisinya dalam satu baris, dengan cakupan variabel hanya di dalam blok `if` atau `else`.")
			fmt.Println("Contoh `if` dengan short statement:")
			fmt.Println("```go")
			fmt.Println("if angka := 10; angka % 2 == 0 {")
			fmt.Println("    fmt.Println(\"Angka genap.\")")
			fmt.Println(" else ")
			fmt.Println("    fmt.Println(\"Angka ganjil.\")")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("Variabel `angka` hanya dapat diakses di dalam blok `if` dan `else`. Ini membantu menjaga kode tetap bersih dan mencegah polusi *namespace*.")
			fmt.Println("Pahami penggunaan percabangan `if` untuk membuat program Anda dapat merespons berbagai skenario input atau kondisi.")
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
		} else {
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

// QuizPercabanganGo (berubah sedikit, parameter pertama sekarang berupa pointer ke array DataQuiz)
func QuizPercabanganGo(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Percabangan `if` pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Struktur kontrol apa yang digunakan untuk mengeksekusi blok kode berdasarkan suatu kondisi di Go?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "if" || strings.Contains(strings.ToLower(jawaban), "percabangan") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'if' atau 'percabangan if'.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Untuk menangani kondisi alternatif jika kondisi `if` pertama `false`, blok apa yang Anda gunakan?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "else" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'else'.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Apa nama fitur `if` di Go yang memungkinkan deklarasi variabel dan pengujian kondisi dalam satu baris?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "short statement") || strings.Contains(strings.ToLower(jawaban), "short declaration") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'short statement' atau 'short declaration'.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	if index >= 0 && index < quiz.NMAX {
		oldScorePercabangan := quizDataArray[index].PercabanganScore
		quizDataArray[index].PercabanganScore = currentQuizScore
		scoreChange := currentQuizScore - oldScorePercabangan
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Percabangan' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package dataquiz
}
