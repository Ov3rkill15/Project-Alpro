
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
	fmt.Scanln(&Choice)

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
		fmt.Println("4. Kembali ke Menu Utama") // Adjusted option number
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			displayQuizScores("original") // Tampilan asli
		case "2":
			displayQuizScores("descending") // Urut menurun
		case "3":
			displayQuizScores("ascending") // Urut menaik
		case "4": // Adjusted case for returning to main menu
			fmt.Println("Kembali ke menu utama...")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Println("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func binarySearchAsc(students [dataquiz.NMAX]atribut.Quiz, count int, target int) int {
	left := 0
	right := count - 1
	foundIndex := -1
	foundMatch := false

	for left <= right && !foundMatch {
		mid := (left + right) / 2
		if students[mid].TotalScore == target {
			foundIndex = mid
			foundMatch = true
		} else if students[mid].TotalScore < target {
			left = mid + 1 // Cari di KANAN
		} else { // students[mid].TotalScore > target
			right = mid - 1 // Cari di KIRI
		}
	}
	return foundIndex
}

// binarySearchDesc melakukan binary search pada array yang diurutkan menurun.
// Mengembalikan indeks pertama yang ditemukan atau -1 jika tidak ada.
func binarySearchDesc(students [dataquiz.NMAX]atribut.Quiz, count int, target int) int {
	left := 0
	right := count - 1
	foundIndex := -1
	foundMatch := false

	for left <= right && !foundMatch {
		mid := (left + right) / 2
		if students[mid].TotalScore == target {
			foundIndex = mid
			foundMatch = true
		} else if students[mid].TotalScore > target { // PENTING: Perubahan di sini untuk Descending!
			left = mid + 1 // Jika nilai tengah > target, target harus di KANAN (nilai lebih kecil)
		} else { // students[mid].TotalScore < target
			right = mid - 1 // Jika nilai tengah < target, target harus di KIRI (nilai lebih besar)
		}
	}
	return foundIndex
}

// displayQuizScores menampilkan semua skor kuis untuk semua siswa.
// Menerima parameter `sortOrder` untuk menentukan urutan tampilan:
// "original", "descending", "ascending".
// Setelah menampilkan, fungsi ini juga menawarkan opsi untuk mencari siswa berdasarkan skor.
func displayQuizScores(sortOrder string) {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          SKOR KUIS SISWA           ")

	// 1. Buat array lokal baru untuk menampung data yang akan diurutkan.
	var studentsToSort [dataquiz.NMAX]atribut.Quiz
	activeCount := 0

	// 2. Salin hanya siswa yang aktif ke array lokal.
	for i := 0; i < dataquiz.NMAX; i++ {
		if dataquiz.DataQuiz[i].Nama != "" {
			studentsToSort[activeCount] = dataquiz.DataQuiz[i]
			activeCount++
		}
	}

	// 3. Lakukan pengurutan untuk DISPLAY
	switch sortOrder {
	case "descending":
		fmt.Println("     (Diurutkan Menurun berdasarkan Total Poin) ")
		pass := 1
		for pass < activeCount {
			i := pass
			temp := studentsToSort[pass]
			for i > 0 && studentsToSort[i-1].TotalScore < temp.TotalScore {
				studentsToSort[i] = studentsToSort[i-1]
				i--
			}
			studentsToSort[i] = temp
			pass++
		}
	case "ascending":
		fmt.Println("     (Diurutkan Menaik berdasarkan Total Poin) ")
		for pass := 0; pass < activeCount-1; pass++ {
			idxMin := pass
			for i := pass + 1; i < activeCount; i++ {
				if studentsToSort[i].TotalScore < studentsToSort[idxMin].TotalScore {
					idxMin = i
				}
			}
			temp := studentsToSort[pass]
			studentsToSort[pass] = studentsToSort[idxMin]
			studentsToSort[idxMin] = temp
		}
	default: // "original"
		fmt.Println("     (Urutan Asli) ")
		for pass := 0; pass < activeCount-1; pass++ {
			idxMin := pass
			for i := pass + 1; i < activeCount; i++ {
				if studentsToSort[i].TotalScore < studentsToSort[idxMin].TotalScore {
					idxMin = i
				}
			}
			temp := studentsToSort[pass]
			studentsToSort[pass] = studentsToSort[idxMin]
			studentsToSort[idxMin] = temp
		}
	}
	fmt.Println("====================================")

	// 4. Tampilkan data (sesuai sortOrder untuk display)
	fmt.Printf("%-15s %-10s %-8s %-8s %-12s %-8s %-8s %-10s %-8s %s\n",
		"Nama", "ID", "Total", "GoLang", "Percabangan", "KonvTD", "OpML", "Perulangan", "TDGo", "VarConst")
	fmt.Println(strings.Repeat("-", 120))

	if activeCount == 0 {
		fmt.Println("Belum ada siswa yang terdaftar atau mengambil kuis.")
	} else {
		for i := 0; i < activeCount; i++ {
			student := studentsToSort[i]
			fmt.Printf("%-15s %-10s %-8d %-8d %-12d %-8d %-8d %-10d %-8d %d\n",
				student.Nama, student.ID, student.TotalScore,
				student.GoLanguageScore, student.PercabanganScore,
				student.KonversiTipeDataScore, student.OperasiMatematikaLogikaScore,
				student.PerulanganScore, student.TipeDataGoScore,
				student.VariableConstantScore,
			)
		}
	}

	fmt.Println(strings.Repeat("-", 120))

	// --- Mulai penambahan binary search di sini ---
	var searchChoice string
	fmt.Println("\n====================================")
	fmt.Println("      Opsi Tambahan Setelah Tampilan Skor    ")
	fmt.Println("====================================")
	fmt.Println("1. Cari Nama Siswa berdasarkan Total Poin (Binary Search)")
	fmt.Println("2. Kembali ke Menu Sebelumnya")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&searchChoice)

	if searchChoice == "1" {
		var searchScore int
		fmt.Print("Masukkan total poin yang dicari: ")
		fmt.Scanln(&searchScore)

		foundIndex := -1
		if sortOrder == "ascending" || sortOrder == "original" { // "original" sudah diurutkan menaik di atas
			foundIndex = binarySearchAsc(studentsToSort, activeCount, searchScore)
		} else if sortOrder == "descending" {
			foundIndex = binarySearchDesc(studentsToSort, activeCount, searchScore)
		}

		if foundIndex != -1 {
			fmt.Printf("\n--- Hasil Pencarian untuk Total Poin %d ---\n", searchScore)
			fmt.Println("Nama Siswa:")

			if sortOrder == "ascending" || sortOrder == "original" {
				// Scan ke kiri (karena terurut menaik)
				i := foundIndex
				for i >= 0 && studentsToSort[i].TotalScore == searchScore {
					fmt.Printf("- %s (ID: %s)\n", studentsToSort[i].Nama, studentsToSort[i].ID)
					i--
				}
				// Scan ke kanan
				i = foundIndex + 1
				for i < activeCount && studentsToSort[i].TotalScore == searchScore {
					fmt.Printf("- %s (ID: %s)\n", studentsToSort[i].Nama, studentsToSort[i].ID)
					i++
				}
			} else if sortOrder == "descending" {
				// Scan ke kiri (karena terurut menurun)
				i := foundIndex
				for i < activeCount && studentsToSort[i].TotalScore == searchScore { // for descending, scan right for smaller index values
					fmt.Printf("- %s (ID: %s)\n", studentsToSort[i].Nama, studentsToSort[i].ID)
					i++
				}
				// Scan ke kanan (karena terurut menurun)
				i = foundIndex - 1 // Start from left of foundIndex
				for i >= 0 && studentsToSort[i].TotalScore == searchScore {
					fmt.Printf("- %s (ID: %s)\n", studentsToSort[i].Nama, studentsToSort[i].ID)
					i--
				}
			}

		} else {
			fmt.Printf("\nTidak ada siswa dengan total poin %d ditemukan.\n", searchScore)
		}
	} else if searchChoice == "2" {
		fmt.Println("Kembali ke menu sebelumnya...")
	} else {
		fmt.Println("Pilihan tidak valid. Kembali ke menu sebelumnya...")
	}
	// --- Akhir penambahan binary search ---

	fmt.Println("Tekan Enter untuk kembali ke menu sebelumnya...")
	fmt.Scanln()
}