package pengpro

import (
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro/konversi_tipe_data"
	"Project-Alpro/pengpro/operasi_matematika_logika"
	pengenalan "Project-Alpro/pengpro/pengenalan_Go"
	"Project-Alpro/pengpro/percabangan"
	"Project-Alpro/pengpro/perulangan"
	dataquiz "Project-Alpro/pengpro/quiz"
	"Project-Alpro/pengpro/tipe_data_go"
	"Project-Alpro/pengpro/variable_constant"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure" // Untuk tampilan ASCII art
)

func MainMenu() {
	var Choice int
	berhenti := true // Flag untuk mengontrol loop menu utama
	atribut.ClearScreen()
	for berhenti {
		// Submenu akan menampilkan opsi dan mengembalikan pilihan pengguna.
		// Jika pengguna memilih 0, aplikasi akan keluar.
		Choice = Submenu()
		if Choice == 0 {
			fmt.Println("Terima kasih telah menggunakan program PENGPRO!")
			berhenti = false // Menghentikan loop utama aplikasi
		} else {
			// Jika pilihan bukan 0, bersihkan layar dan tampilkan menu utama lagi.
			atribut.ClearScreen()
		}
	}
}

// Submenu menampilkan daftar materi utama dan mengarahkan ke menu materi spesifik.
// Mengembalikan pilihan pengguna. Jika 0, mengindikasikan ingin keluar dari aplikasi.
func Submenu() int {
	var Choice int
	welcome := figure.NewFigure("PENGPRO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")  // Reset warna
	fmt.Println("== Pengenalan Pemrograman ==")
	fmt.Println("1. Apa itu Pemrograman bahasa Go?") // OPSI PERTAMA
	fmt.Println("2. Pengenalan tipe data (Number, String, Boolean)")
	fmt.Println("3. Variable dan Constant")
	fmt.Println("4. Konversi tipe data")
	fmt.Println("5. Operasi aritmatika dan logika")
	fmt.Println("6. Perulangan (For Loop, While Loop, Repeat Until)")
	fmt.Println("7. Percabangan (If Else, Switch Case)")
	fmt.Println("8. Tampilkan Skor Kuis")
	fmt.Println("9. Daftar Siswa Baru")
	fmt.Println("0. Keluar Aplikasi")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&Choice) // Menggunakan Scanln untuk membaca seluruh baris dan membersihkan buffer

	switch Choice {
	case 1:
		pengenalan.MainMenu(&dataquiz.DataQuiz)
	case 2:
		tipe_data_go.MainMenu(&dataquiz.DataQuiz)
	case 3:
		variable_constant.MainMenu(&dataquiz.DataQuiz)
	case 4:
		konversi_tipe_data.MainMenu(&dataquiz.DataQuiz)
	case 5:
		operasi_matematika_logika.MainMenu(&dataquiz.DataQuiz)
	case 6:
		perulangan.MainMenu(&dataquiz.DataQuiz)
	case 7:
		percabangan.MainMenu(&dataquiz.DataQuiz)
	case 8:
		displayQuizScores()
	case 9:
		dataquiz.HandleNewStudentRegistration()
	case 0:
		return 0
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		fmt.Println("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}
	return Choice
}

func displayQuizScores() {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          SKOR KUIS SISWA           ")
	fmt.Println("====================================")
	// Sesuaikan header ini dengan semua field skor yang ada di types.Quiz
	fmt.Printf("%-15s %-10s %-8s %-8s %-12s %-8s %-8s %-10s %-8s %s\n",
		"Nama", "ID", "Total", "GoLang", "Percabangan", "KonvTD", "OpML", "Perulangan", "TDGo", "VarConst")
	fmt.Println(strings.Repeat("-", 120)) // Sesuaikan panjang garis agar sesuai header

	for i := 0; i < dataquiz.NMAX; i++ {
		if dataquiz.DataQuiz[i].Nama != "" { // Hanya tampilkan entri yang memiliki nama
			fmt.Printf("%-15s %-10s %-8d %-8d %-12d %-8d %-8d %-10d %-8d %d\n",
				dataquiz.DataQuiz[i].Nama,
				dataquiz.DataQuiz[i].ID,
				dataquiz.DataQuiz[i].TotalScore,
				dataquiz.DataQuiz[i].GoLanguageScore,
				dataquiz.DataQuiz[i].PercabanganScore,
				dataquiz.DataQuiz[i].KonversiTipeDataScore,
				dataquiz.DataQuiz[i].OperasiMatematikaLogikaScore,
				dataquiz.DataQuiz[i].PerulanganScore,
				dataquiz.DataQuiz[i].TipeDataGoScore,
				dataquiz.DataQuiz[i].VariableConstantScore,
			)
		}
	}
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("Tekan Enter untuk kembali ke menu utama...")
	fmt.Scanln()
}
