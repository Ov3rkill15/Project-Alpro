package pengenalan_Go

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)


func MainMenu() {
	var choice string
	// Use a boolean flag to control the main menu loop
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SELAMAT DATANG!           ")
		fmt.Println("====================================")
		fmt.Println("Pilih materi yang ingin Anda pelajari:")
		fmt.Println("1. Pengenalan Bahasa Go")
		fmt.Println("2. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			Submenu() // Go to the Go Language submenu
		case "2":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			berhenti = false // Set flag to false to exit the loop
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			// The loop will repeat because berhenti is still true
		}
	}
}


func Submenu() {
	var choice string
	// Use a boolean flag to control the submenu loop
	berhenti2 := true
	for berhenti2 {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          SUBMENU MATERI GO         ")
		fmt.Println("====================================")
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Baca Materi Go Language")
		fmt.Println("2. Kerjakan Kuis Go Language")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			// Call MateriGoLanguage; if it returns true, go to QuizGoLanguage
			if MateriGoLanguage() {
				QuizGoLanguage()
			}
			
		case "2":
			QuizGoLanguage() // Start the Go language quiz
		case "3":
			fmt.Println("Kembali ke menu utama...")
			berhenti2 = false // Set flag to false to exit the loop
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			// The loop will repeat because berhenti2 is still true
		}
	}
}

// PromptKembaliToMain prompts the user to return to the submenu.
func PromptKembaliToMain() {
	var choice string
	// Use a boolean flag to control the prompt loop
	berhenti3 := true
	for berhenti3 {
		fmt.Print("\nTekan 'm' untuk kembali ke submenu: ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "m" {
			berhenti3 = false // Set flag to false to exit the loop
		} else {
			fmt.Println("Pilihan tidak valid. Silakan masukkan 'm'.")
			// The loop will repeat because berhenti3 is still true
		}
	}
	// After the loop, the function returns naturally
}

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
			fmt.Println("- Command-Line Interface (CLI) tools")
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


func QuizGoLanguage() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Apa itu Pemrograman Bahasa Go? ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Bahasa Go dikembangkan oleh perusahaan apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.ToLower(jawaban) == "google" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah Google.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Sebutkan salah satu fitur utama Go yang terkait dengan konkurensi (pemrosesan paralel).")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "goroutine") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah goroutine.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Go sering digunakan untuk pengembangan apa yang terkait dengan API dan layanan-layanan kecil?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scanln(&jawaban)
	if strings.Contains(strings.ToLower(jawaban), "microservices") || strings.Contains(strings.ToLower(jawaban), "api") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah microservices atau API.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", score)
	fmt.Println("==============================")
	PromptKembaliToMain()
}
