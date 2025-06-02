package p_golanguange

import (
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro/quiz"
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Go Language.
// Menerima pointer ke array DataQuiz untuk mengelola data siswa.
func MainMenu(quizDataArray *[quiz.NMAX]atribut.Quiz) {
	var choice string
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SELAMAT DATANG!           ")
		fmt.Println("====================================")
		fmt.Println("Pilih materi yang ingin Anda pelajari:")
		fmt.Println("1. Pengenalan Bahasa Go")
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			Submenu() // Panggilan Submenu SUDAH BENAR TANPA PARAMETER
		case "2":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// Submenu menampilkan submenu untuk materi Go Language.
func Submenu() { // Definisi fungsi ini TANPA PARAMETER sudah benar
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SUBMENU MATERI GO         ")
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Go Language")
		fmt.Println("2. Kerjakan Kuis Go Language")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriGoLanguage() {
				quiz.HandleQuizStart(QuizGoLanguage, "Pengenalan Go")
			}
		case "2":
			quiz.HandleQuizStart(QuizGoLanguage, "Pengenalan Go")
		case "3":
			fmt.Println("Kembali ke menu utama...")
			berhenti2 = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// MateriGoLanguage (tidak berubah)
func MateriGoLanguage() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var choice string
	readingMaterial := true
	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Pemrograman Bahasa Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)
		switch halamanSekarang {
		case 1:
			fmt.Println("Bahasa pemrograman Go, atau sering disebut Golang, adalah bahasa pemrograman yang dikembangkan oleh Google. Golang pertama kali dirilis pada tahun 2009 oleh tiga insinyur Google:")
			fmt.Println("Robert Griesemer, Rob Pike, dan Ken Thompson. Tujuan utama dari pengembangan Golang adalah untuk menciptakan bahasa yang efisien, sederhana, dan mudah dipelajari, serta mendukung concurrency dengan baik.")
			fmt.Println("Beberapa fitur utamanya termasuk konkurensi (goroutine), garbage collection, dan sintaksis yang mirip C.")
		case 2:
			fmt.Println("Bahasa Go sangat populer untuk:")
			fmt.Println("- Pengembangan Web (dengan framework seperti Gin atau Echo)")
			fmt.Println("- Microservices dan API")
		case 3:
			fmt.Println("- Sistem jaringan dan distributed systems")
			fmt.Println("Go juga memiliki komunitas yang berkembang pesat dan ekosistem yang kaya dengan banyak *library* dan *tool*.")
			fmt.Println("Sintaksisnya yang bersih dan sederhana membuatnya mudah dipelajari, terutama bagi yang sudah familiar dengan C/C++ atau Java.")
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

// QuizGoLanguage (tidak berubah)
func QuizGoLanguage(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Apa itu Pemrograman Bahasa Go? ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Bahasa Go dikembangkan oleh perusahaan apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "google" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah Google.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Sebutkan salah satu fitur utama Go yang terkait dengan konkurensi (pemrosesan paralel).")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "goroutine") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah goroutine.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Go sering digunakan untuk pengembangan apa yang terkait dengan API dan layanan-layanan kecil?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "microservices") || strings.Contains(strings.ToLower(jawaban), "api") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah microservices atau API.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	if index >= 0 && index < quiz.NMAX {
		oldScoreGoLanguage := quizDataArray[index].GoLanguageScore
		quizDataArray[index].GoLanguageScore = currentQuizScore
		scoreChange := currentQuizScore - oldScoreGoLanguage
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Pengenalan Go' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain()
}
