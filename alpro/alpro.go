package alpro

import (
	"Project-Alpro/alpro/contohsoal"
	"Project-Alpro/alpro/p_array"
	"Project-Alpro/alpro/p_function"
	"Project-Alpro/alpro/p_maxmin"
	"Project-Alpro/alpro/p_procedure"
	"Project-Alpro/alpro/p_rekursif"
	"Project-Alpro/alpro/p_searching"
	"Project-Alpro/alpro/p_sorting"
	p_tipeBentukan "Project-Alpro/alpro/p_tipebentukan"
	"Project-Alpro/alpro/praktikum"
	"Project-Alpro/alpro/quiz" // Import package quiz
	"Project-Alpro/atribut"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// MainMenu adalah fungsi utama aplikasi.
func MainMenu() int {
	reader := bufio.NewReader(os.Stdin)
	var stop bool = true

	// --- PENTING: Inisialisasi data siswa di awal program ---
	// Panggil ini sekali untuk mengisi studentsData di package quiz
	quiz.InitStudentsData()
	// --- Akhir bagian inisialisasi ---

	for stop {
		atribut.ClearScreen()
		Submenu() // Menampilkan menu utama
		fmt.Print("Masukkan pilihan: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			atribut.ClearScreen()
			p_function.MainMenu()
		case "2":
			atribut.ClearScreen()
			p_procedure.MainMenu()
		case "3":
			atribut.ClearScreen()
			p_rekursif.MainMenu()
		case "4":
			atribut.ClearScreen()
			p_tipeBentukan.MainMenu()
		case "5":
			atribut.ClearScreen()
			p_array.MainMenu()
		case "6":
			atribut.ClearScreen()
			p_maxmin.MainMenu()
		case "7":
			atribut.ClearScreen()
			p_searching.MainMenu()
		case "8":
			atribut.ClearScreen()
			p_sorting.MainMenu()
		case "9":
			atribut.ClearScreen()
			praktikum.MainMenu()
		case "10": // Opsi "contoh soal"
			atribut.ClearScreen()
			// Gunakan quiz.StartQuizSession untuk memulai kuis.
			// Ini akan menangani pengambilan ID siswa dan meneruskan pointer skor.
			quiz.StartQuizSession(contohsoal.StartSoalMenu, "Contoh Soal Umum")
		case "11": // Opsi "scores"
			atribut.ClearScreen()
			quiz.DisplayScoresMenu()
		case "12": // Opsi "Register"
			atribut.ClearScreen()
			// quiz.InitStudentsData() // Tidak perlu dipanggil lagi di sini, sudah di awal MainMenu
			quiz.RegisterNewStudentFlow()
		case "0": // Opsi "Keluar"
			atribut.ClearScreen()
			atribut.Loading(1200) // Animasi loading
			stop = false          // Menghentikan loop menu utama
			return 0
		default: // Pilihan tidak valid
			atribut.ClearScreen()
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			atribut.Loading(1200) // Animasi loading
		}
	}
	return 0
}

// Submenu menampilkan menu utama aplikasi.
func Submenu() {
	welcome := figure.NewFigure("ALPRO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")  // Reset warna
	fmt.Println(`
====================================
Selamat Datang di Algoritma dan Pemrograman
====================================
1. Pembelajaran Function
2. Pembelajaran Procedure
3. Pembelajaran Rekursif
4. Pembelajaran TipeBentukan
5. Pembelajaran Array
6. Pembelajaran maxmin
7. Pembelajaran Searching
8. Pembelajaran Sorting
9. Referensi Soal
10. Contoh Soal
11. Lihat Skor Siswa
12. Daftar/Registrasi Siswa Baru

0. Keluar
     `)
}
