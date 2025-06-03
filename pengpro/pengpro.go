package pengpro

import (
	"Project-Alpro/atribut"
	"Project-Alpro/pengpro/konversi_tipe_data"
	"Project-Alpro/pengpro/operasi_matematika_logika"
	pengenalan "Project-Alpro/pengpro/pengenalan_Go"
	"Project-Alpro/pengpro/percabangan"
	"Project-Alpro/pengpro/perulangan"
	dataquiz "Project-Alpro/pengpro/quiz"
	tipe_data_go "Project-Alpro/pengpro/tipe_data_Go"
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
	fmt.Println("8. Daftar Siswa Baru")   // Opsi daftar siswa di sini
	fmt.Println("9. Tampilkan Skor Kuis") // Opsi untuk menampilkan submenu skor
	fmt.Println("0. Keluar Aplikasi")     // Opsi keluar aplikasi
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&Choice) // Menggunakan Scanln untuk membaca seluruh baris dan membersihkan buffer

	switch Choice {
	case 1:
		p_golanguange.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 2:
		tipe_data_go.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 3:
		variable_constant.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 4:
		konversi_tipe_data.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 5:
		operasi_matematika_logika.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 6:
		perulangan.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 7:
		percabangan.MainMenu(&dataquiz.DataQuiz) // Meneruskan pointer ke array DataQuiz global
	case 8: // Daftar Siswa Baru
		dataquiz.HandleNewStudentRegistration()
	case 9: // Masuk ke submenu tampilan skor
		DisplayScoresMenu()
	case 0:
		return 0 // Mengindikasikan ingin keluar dari aplikasi
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		fmt.Println("Tekan Enter untuk melanjutkan...")
		fmt.Scanln() // Membersihkan buffer input setelah pesan kesalahan
	}
	return Choice // Mengembalikan pilihan untuk melanjutkan loop di MainMenu
}

// DisplayScoresMenu menampilkan submenu untuk pilihan tampilan skor.
func DisplayScoresMenu() {
	var choice string
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          MENU TAMPILAN SKOR        ")
		fmt.Println("====================================")
		fmt.Println("Pilih opsi tampilan skor:")
		fmt.Println("1. Daftar Nilai (Urutan Asli)")
		fmt.Println("2. Daftar Nilai (Urut Menurun berdasarkan Total Poin)")
		fmt.Println("3. Daftar Nilai (Urut Menaik berdasarkan Total Poin)")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			displayQuizScores("original") // Tampilan asli
		case "2":
			displayQuizScores("descending") // Urut menurun
		case "3":
			displayQuizScores("ascending") // Urut menaik
		case "4":
			fmt.Println("Kembali ke menu utama...")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Println("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

// displayQuizScores menampilkan semua skor kuis untuk semua siswa.
// Menerima parameter `sortOrder` untuk menentukan urutan tampilan:
// "original", "descending", "ascending".
func displayQuizScores(sortOrder string) {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          SKOR KUIS SISWA           ")

	// 1. Buat array lokal baru untuk menampung data yang akan diurutkan.
	// Inisialisasi dengan semua slot kosong (zero value).
	var studentsToSort [dataquiz.NMAX]atribut.Quiz
	activeCount := 0 // Menghitung berapa banyak siswa aktif yang disalin

	// 2. Salin hanya siswa yang aktif ke array lokal.
	// Ini juga akan "memadatkan" data ke awal array studentsToSort.
	for i := 0; i < dataquiz.NMAX; i++ {
		if dataquiz.DataQuiz[i].Nama != "" {
			studentsToSort[activeCount] = dataquiz.DataQuiz[i]
			activeCount++
		}
	}

	// 3. Lakukan pengurutan pada bagian yang berisi siswa aktif dari studentsToSort
	switch sortOrder {
	case "descending":
		fmt.Println("     (Diurutkan Menurun berdasarkan Total Poin) ")
		// Insertion Sort untuk urut menurun, disesuaikan dengan format while-loop
		// Kamus: i, pass : integer; temp : types.Quiz
		pass := 1                // pass <- 1
		for pass < activeCount { // while pass < activeCount do (setara dengan pass <= n-1 di pseudocode jika n=activeCount)
			// { Pencarian indeks yang tepat untuk elemen }
			i := pass
			temp := studentsToSort[pass] // temp <- A[pass]

			// while i > 0 and A[i-1] < temp.TotalScore do (untuk menurun)
			for i > 0 && studentsToSort[i-1].TotalScore < temp.TotalScore {
				studentsToSort[i] = studentsToSort[i-1] // A[i] <- A[i-1]
				i--                                     // i <- i - 1
			}
			// { Menempatkan elemen pada lokasi tersebut}
			studentsToSort[i] = temp // A[i] <- temp

			pass++ // pass <- pass + 1
		}
	case "ascending":
		fmt.Println("     (Diurutkan Menaik berdasarkan Total Poin) ")
		pass := 0                  // pass <- 0
		for pass < activeCount-1 { // while pass <= n-1 do (simulasi)
			// { 1. Pencarian nilai idx ekstrim (minimum) }
			idx := pass           // idx <- pass
			i := pass + 1         // i <- pass + 1
			for i < activeCount { // while i < n do (simulasi)
				if studentsToSort[i].TotalScore < studentsToSort[idx].TotalScore { // Kondisi untuk mencari MINIMUM
					idx = i
				}
				i++ // i <- i + 1
			}
			// { 2. Pertukaran }
			temp := studentsToSort[idx]                // temp <- A[idx]
			studentsToSort[idx] = studentsToSort[pass] // A[idx] <- A[pass]
			studentsToSort[pass] = temp                // A[pass] <- temp

			pass++ // pass <- pass + 1
		}
	default: // "original" atau jika ada nilai sortOrder lain yang tidak dikenali
		fmt.Println("     (Urutan Asli) ")
	}
	fmt.Println("====================================")

	// 4. Tampilkan data
	fmt.Printf("%-15s %-10s %-8s %-8s %-12s %-8s %-8s %-10s %-8s %s\n",
		"Nama", "ID", "Total", "GoLang", "Percabangan", "KonvTD", "OpML", "Perulangan", "TDGo", "VarConst")
	fmt.Println(strings.Repeat("-", 120))

	if activeCount == 0 { // Cek apakah ada siswa aktif yang ditemukan
		fmt.Println("Belum ada siswa yang terdaftar atau mengambil kuis.")
	} else {
		if sortOrder == "original" {
			// Tampilkan dari array global asli untuk menjaga urutan awal
			for i := 0; i < dataquiz.NMAX; i++ {
				if dataquiz.DataQuiz[i].Nama != "" {
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
		} else {
			// Tampilkan dari array lokal yang sudah diurutkan (studentsToSort)
			// Hanya loop sebanyak activeCount untuk menghindari slot kosong
			for i := 0; i < activeCount; i++ {
				student := studentsToSort[i] // Ambil data dari array yang diurutkan
				fmt.Printf("%-15s %-10s %-8d %-8d %-12d %-8d %-8d %-10d %-8d %d\n",
					student.Nama,
					student.ID,
					student.TotalScore,
					student.GoLanguageScore,
					student.PercabanganScore,
					student.KonversiTipeDataScore,
					student.OperasiMatematikaLogikaScore,
					student.PerulanganScore,
					student.TipeDataGoScore,
					student.VariableConstantScore,
				)
			}
		}
	}

	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("Tekan Enter untuk kembali ke menu sebelumnya...")
	fmt.Scanln()
}

