package perulangan

import (
	"Project-Alpro/atribut"          // Untuk fungsi ClearScreen
	"Project-Alpro/pengpro/quiz" // Untuk mengakses DataQuiz, NMAX, dan fungsi pembantu

	// Untuk mengakses struct Quiz
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Perulangan.
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
		fmt.Println("1. Perulangan (Looping) Go") // Mengubah teks menu
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

// Submenu menampilkan submenu untuk materi Perulangan.
func Submenu() { // Tidak ada parameter di sini
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("   SUBMENU MATERI PERULANGAN (LOOPING)") // Mengubah teks menu
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Perulangan (Looping) Go")   // Mengubah teks menu
		fmt.Println("2. Kerjakan Kuis Perulangan (Looping) Go") // Mengubah teks menu
		fmt.Println("3. Kembali ke Menu Utama")
		// Opsi "Daftar / Masukkan Data Baru" sudah dipindahkan ke pengpro.Submenu
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriPerulangan() { // Mengubah nama fungsi materi
				quiz.HandleQuizStart(QuizPerulangan, "Perulangan") // Mengubah nama fungsi kuis dan nama materi
			}
		case "2":
			quiz.HandleQuizStart(QuizPerulangan, "Perulangan") // Mengubah nama fungsi kuis dan nama materi
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

// MateriPerulangan menampilkan materi Perulangan (Looping) Go secara berhalaman.
// Mengembalikan true jika pengguna ingin melanjutkan ke kuis, false jika kembali ke submenu.
func MateriPerulangan() bool { // Mengubah nama fungsi dari MateriGoLanguage
	halamanSekarang := 1
	totalHalaman := 3
	var choice string

	readingMaterial := true

	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Perulangan (Looping) pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, **perulangan** digunakan untuk menjalankan blok kode berulang kali. Go hanya memiliki satu struktur perulangan: `for` loop, yang sangat fleksibel.")
			fmt.Println("Bentuk paling umum dari `for` loop adalah **`for` loop dengan tiga komponen**: inisialisasi, kondisi, dan post-statement.")
			fmt.Println("Contoh:")
			fmt.Println("```go")
			fmt.Println("for i := 0; i < 5; i++ {")
			fmt.Println("    fmt.Println(\"Angka:\", i)")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("Ini akan mencetak angka 0 hingga 4. `i := 0` (inisialisasi), `i < 5` (kondisi), `i++` (post-statement).")
		case 2:
			fmt.Println("Anda juga bisa menggunakan `for` loop seperti **`while` loop** di bahasa lain, hanya dengan kondisi:")
			fmt.Println("```go")
			fmt.Println("sum := 1")
			fmt.Println("for sum < 100 {")
			fmt.Println("    sum += sum")
			fmt.Println("}")
			fmt.Println("fmt.Println(sum)")
			fmt.Println("```")
			fmt.Println("Dan bentuk **infinite loop** (perulangan tak terbatas) tanpa kondisi:") // Menghilangkan referensi break
			fmt.Println("```go")
			fmt.Println("k := 0")
			fmt.Println("for { // Loop tanpa batas")
			fmt.Println("    fmt.Println(\"Perulangan ke\", k)")
			fmt.Println("    k++")
			fmt.Println("    // Biasanya dihentikan dengan kondisi internal atau dari luar")
			fmt.Println("    // if k > 2 { break } // Jika break dilarang, perlu cara lain untuk keluar loop tak terbatas ini")
			fmt.Println("}")
			fmt.Println("```")
		case 3:
			fmt.Println("Go juga memiliki `for...range` loop, yang digunakan untuk mengiterasi elemen-elemen dalam struktur data seperti array, slice, string, dan map.")
			fmt.Println("Contoh `for...range` dengan slice:")
			fmt.Println("```go")
			fmt.Println("kata := []string{\"apel\", \"jeruk\", \"mangga\"}")
			fmt.Println("for index, value := range kata {")
			fmt.Println("    fmt.Printf(\"Index: %d, Nilai: %s\\n\", index, value)")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("Memahami perulangan sangat penting untuk memproses daftar data atau mengulang tugas-tugas tertentu dalam program Anda.") // Menghilangkan referensi break/continue
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

// QuizPerulangan menjalankan kuis Perulangan (Looping) dan memperbarui skor siswa.
// Menerima pointer ke array DataQuiz dan indeks siswa yang sedang kuis.
func QuizPerulangan(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) { // Mengubah nama fungsi dari QuizGoLanguage
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Perulangan (Looping) pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Go hanya memiliki satu kata kunci untuk perulangan. Apa itu?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "for" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'for'.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Bentuk `for` loop seperti apa yang digunakan untuk menghentikan perulangan secara paksa, jika kita diperbolehkan menggunakan kata kunci tersebut?") // Menyesuaikan teks pertanyaan
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "break" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'break'.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Bentuk `for` loop yang digunakan untuk mengiterasi elemen dalam array atau slice adalah `for...` apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "range" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'range'.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	// Perbarui skor siswa di array DataQuiz global melalui pointer
	if index >= 0 && index < quiz.NMAX {
		// Mengakses field skor spesifik untuk Perulangan
		oldScorePerulangan := quizDataArray[index].PerulanganScore
		quizDataArray[index].PerulanganScore = currentQuizScore
		scoreChange := currentQuizScore - oldScorePerulangan
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Perulangan' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package dataquiz
}
