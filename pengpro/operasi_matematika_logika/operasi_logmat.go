package operasi_matematika_logika

import (
	"Project-Alpro/atribut"          // Untuk fungsi ClearScreen
	"Project-Alpro/pengpro/quiz" // Untuk mengakses DataQuiz, NMAX, dan fungsi pembantu
	"fmt"
	"strings"
)

// MainMenu menampilkan menu utama untuk materi Operasi Aritmatika & Logika.
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
		fmt.Println("1. Operasi Aritmatika & Logika Go") // Mengubah teks menu
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

// Submenu menampilkan submenu untuk materi Operasi Aritmatika & Logika.
func Submenu() { // Tidak ada parameter di sini
	var choice string
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println(" SUBMENU MATERI OPERASI ARITMATIKA & LOGIKA") // Mengubah teks menu
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Operasi Aritmatika & Logika Go")   // Mengubah teks menu
		fmt.Println("2. Kerjakan Kuis Operasi Aritmatika & Logika Go") // Mengubah teks menu
		fmt.Println("3. Kembali ke Menu Utama")
		// Opsi "Daftar / Masukkan Data Baru" dipindahkan ke pengpro.Submenu
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if MateriOperasiMatematikaLogika() { // Mengubah nama fungsi materi
				quiz.HandleQuizStart(QuizOperasiMatematikaLogika, "Operasi Aritmatika & Logika") // Mengubah nama fungsi kuis dan nama materi
			}
		case "2":
			quiz.HandleQuizStart(QuizOperasiMatematikaLogika, "Operasi Aritmatika & Logika") // Mengubah nama fungsi kuis dan nama materi
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

// MateriOperasiMatematikaLogika menampilkan materi Operasi Aritmatika dan Logika secara berhalaman.
// Mengembalikan true jika pengguna ingin melanjutkan ke kuis, false jika kembali ke submenu.
func MateriOperasiMatematikaLogika() bool { // Mengubah nama fungsi dari MateriGoLanguage
	halamanSekarang := 1
	totalHalaman := 3
	var choice string

	readingMaterial := true

	for readingMaterial {
		atribut.ClearScreen()
		fmt.Printf("=== Operasi Aritmatika dan Logika pada Bahasa Go (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, Anda dapat melakukan berbagai **operasi aritmatika** untuk perhitungan matematis.")
			fmt.Println("Operator aritmatika dasar meliputi:")
			fmt.Println("  - Penjumlahan: `+` (misal: `5 + 3 = 8`)")
			fmt.Println("  - Pengurangan: `-` (misal: `10 - 4 = 6`)")
			fmt.Println("  - Perkalian: `*` (misal: `2 * 6 = 12`)")
			fmt.Println("  - Pembagian: `/` (misal: `15 / 3 = 5`)")
			fmt.Println("  - Modulus (sisa bagi): `%` (misal: `10 % 3 = 1`)")
			fmt.Println("Contoh penggunaan:")
			fmt.Println("  `hasilTambah := 7 + 2`")
			fmt.Println("  `hasilBagi := 10.0 / 4.0`")
			fmt.Println("Pastikan tipe data yang dioperasikan kompatibel. Pembagian integer akan menghasilkan integer.")
		case 2:
			fmt.Println("Sekarang mari kita bahas **operasi logika**.")
			fmt.Println("Operator logika digunakan untuk menggabungkan atau memanipulasi nilai boolean (`true` atau `false`).")
			fmt.Println("Operator logika dasar meliputi:")
			fmt.Println("  - AND (Konjungsi): `&&`")
			fmt.Println("    - `true && true` hasilnya `true`")
			fmt.Println("    - `true && false` hasilnya `false`")
			fmt.Println("    - `false && false` hasilnya `false`")
			fmt.Println("  - OR (Disjungsi): `||`")
			fmt.Println("    - `true || true` hasilnya `true`")
			fmt.Println("    - `true || false` hasilnya `true`")
			fmt.Println("    - `false || false` hasilnya `false`")
		case 3:
			fmt.Println("  - NOT (Negasi): `!`")
			fmt.Println("    - `!true` hasilnya `false`")
			fmt.Println("    - `!false` hasilnya `true`")
			fmt.Println("Contoh penggunaan operator logika:")
			fmt.Println("  `isRaining := true`")
			fmt.Println("  `hasUmbrella := false`")
			fmt.Println("  `canGoOut := isRaining && hasUmbrella`")
			fmt.Println("  `needsCoat := isRaining || !hasUmbrella`")
			fmt.Println("Operasi logika sangat penting untuk membuat keputusan (branching) dalam program Anda.")
			fmt.Println("Dengan memahami operator ini, Anda dapat mengontrol alur program berdasarkan kondisi yang berbeda.")
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

// QuizOperasiMatematikaLogika menjalankan kuis Operasi Aritmatika dan Logika dan memperbarui skor siswa.
// Menerima pointer ke array DataQuiz dan indeks siswa yang sedang kuis.
func QuizOperasiMatematikaLogika(quizDataArray *[quiz.NMAX]atribut.Quiz, index int) { // Mengubah nama fungsi dari QuizGoLanguage
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Operasi Aritmatika dan Logika pada Bahasa Go ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	currentQuizScore := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Operator `%` dalam Go digunakan untuk operasi apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "modulus") || strings.Contains(strings.ToLower(jawaban), "sisa bagi") {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah modulus (sisa bagi).")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Jika `kondisiA` adalah `true` dan `kondisiB` adalah `false`, apa hasil dari `kondisiA && kondisiB`?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "false" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah 'false'.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Operator logika apa yang digunakan untuk membalikkan nilai boolean?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "!" {
		fmt.Println("Benar!")
		currentQuizScore++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah '!'.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", currentQuizScore)
	fmt.Println("==============================")

	// Perbarui skor siswa di array DataQuiz global melalui pointer
	if index >= 0 && index < quiz.NMAX {
		// Mengakses field skor spesifik untuk Operasi Matematika & Logika
		oldScoreOperasiMatematikaLogika := quizDataArray[index].OperasiMatematikaLogikaScore
		quizDataArray[index].OperasiMatematikaLogikaScore = currentQuizScore
		scoreChange := currentQuizScore - oldScoreOperasiMatematikaLogika
		quizDataArray[index].TotalScore += scoreChange

		fmt.Printf("Skor %s (ID: %s) untuk 'Operasi Aritmatika & Logika' telah diperbarui: %d\n", quizDataArray[index].Nama, quizDataArray[index].ID, currentQuizScore)
		fmt.Printf("Total skor kumulatif Anda: %d\n", quizDataArray[index].TotalScore)
	} else {
		fmt.Println("Error: Indeks data kuis tidak valid. Skor tidak dapat disimpan.")
	}

	quiz.PromptKembaliToMain() // Memanggil PromptKembaliToMain dari package dataquiz
}
