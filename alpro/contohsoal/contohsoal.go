package contohsoal // Changed package name

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Soal merepresentasikan struktur soal dengan pertanyaan dan daftar test case.
type Soal struct {
	Materi     string
	Pertanyaan string
	// Menggunakan array tetap untuk test cases karena batasan "tanpa slice".
	// TestCases mungkin perlu nmax yang lebih kecil atau diisi dengan string kosong
	// jika jumlah test case bervariasi per soal.
	TestCases    [10]string // Menggunakan array tetap dengan ukuran yang cukup
	NumTestCases int        // Menyimpan jumlah test case yang valid
}

// openNotepadPlusPlus membuka file yang diberikan menggunakan Notepad++.
func openNotepadPlusPlus(filePath string) {
	exePath := "C:\\Program Files\\Notepad++\\notepad++.exe" // Sesuaikan path ini jika Notepad++ Anda diinstal di lokasi lain.
	cmd := exec.Command(exePath, filePath)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Gagal membuka Notepad++:", err)
	}
}

// compileAndRun mengkompilasi dan menjalankan file jawaban.go, lalu mengembalikan outputnya.
func compileAndRun() (string, error) {
	cmd := exec.Command("go", "run", "jawaban.go")
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// StartSoalMenu adalah fungsi utama yang mengelola alur soal dan pengecekan jawaban.
// Ini sekarang menerima pointer ke float64 (nilaiTemp) dan mengembalikan boolean.
func StartSoalMenu(nilaiTemp *float64) bool { // Renamed and added return value
	var Choice2 string
	reader := bufio.NewReader(os.Stdin)

	// Definisi semua soal secara langsung di dalam kode Go
	// Pastikan TestCases dan NumTestCases diisi dengan benar.
	daftarSoal := [10]Soal{ // Menggunakan array tetap dengan ukuran NMAX
		{
			Materi:       "Function (Sederhana)",
			Pertanyaan:   "Buatlah sebuah fungsi Go bernama `hitungJumlah` yang menerima dua bilangan bulat ($a$ dan $b$) sebagai input, lalu mengembalikan hasil penjumlahan kedua bilangan tersebut. Cetak hasil pemanggilan fungsi ini dengan input 5 dan 3.",
			TestCases:    [10]string{"8"},
			NumTestCases: 1,
		},
		{
			Materi:       "Procedure (Sederhana)",
			Pertanyaan:   "Buatlah sebuah fungsi Go bernama `sapaPengguna` yang menerima sebuah string (nama) sebagai input dan mencetak \"Halo, [nama]!\" ke konsol. Panggil fungsi ini dengan nama \"Budi\".",
			TestCases:    [10]string{"Halo, Budi!"},
			NumTestCases: 1,
		},
		{
			Materi:       "Rekursif (Sederhana)",
			Pertanyaan:   "Buatlah sebuah fungsi rekursif Go bernama `faktorial` yang menghitung nilai faktorial dari sebuah bilangan bulat non-negatif $n$. Cetak hasil faktorial dari 4.",
			TestCases:    [10]string{"24"},
			NumTestCases: 1,
		},
		{
			Materi:       "Tipe Bentukan (Menengah)",
			Pertanyaan:   "Definisikan sebuah `struct` Go bernama `Mahasiswa` dengan field `Nama` (string) dan `NIM` (string). Buat sebuah instance `Mahasiswa` dengan Nama \"Dewi\" dan NIM \"123456789\". Cetak `NIM` dari instance tersebut.",
			TestCases:    [10]string{"123456789"},
			NumTestCases: 1,
		},
		{
			Materi:       "Array (Menengah)",
			Pertanyaan:   "Buatlah sebuah array Go berukuran 5 yang berisi bilangan bulat: 10, 20, 30, 40, 50. Cetak elemen pada indeks ke-2 dari array tersebut.",
			TestCases:    [10]string{"30"},
			NumTestCases: 1,
		},
		{
			Materi:       "Pencarian Nilai Max (Menengah)",
			Pertanyaan:   "Diberikan sebuah array bilangan bulat: [15, 7, 22, 10, 30, 5]. Temukan dan cetak nilai maksimum dalam array tersebut.",
			TestCases:    [10]string{"30"},
			NumTestCases: 1,
		},
		{
			Materi:       "Sequential Search (Menengah)",
			Pertanyaan:   "Diberikan sebuah array bilangan bulat: [5, 12, 8, 20, 15, 3]. Lakukan pencarian sekuensial untuk nilai 20. Jika ditemukan, cetak \"Nilai ditemukan!\". Jika tidak, cetak \"Nilai tidak ditemukan.\".",
			TestCases:    [10]string{"Nilai ditemukan!"},
			NumTestCases: 1,
		},
		{
			Materi:       "Binary Search (Menengah)",
			Pertanyaan:   "Diberikan sebuah array bilangan bulat yang sudah terurut: [10, 20, 30, 40, 50, 60]. Lakukan pencarian biner untuk nilai 40. Jika ditemukan, cetak \"Nilai ditemukan!\". Jika tidak, cetak \"Nilai tidak ditemukan.\".",
			TestCases:    [10]string{"Nilai ditemukan!"},
			NumTestCases: 1,
		},
		{
			Materi:       "Selection Sort (Menengah)",
			Pertanyaan:   "Diberikan sebuah array bilangan bulat: [64, 25, 12, 22, 11]. Terapkan algoritma Selection Sort untuk mengurutkan array ini secara ascending. Cetak array setelah diurutkan.",
			TestCases:    [10]string{"[11 12 22 25 64]"},
			NumTestCases: 1,
		},
		{
			Materi:       "Insertion Sort (Menengah)",
			Pertanyaan:   "Diberikan sebuah array bilangan bulat: [12, 11, 13, 5, 6]. Terapkan algoritma Insertion Sort untuk mengurutkan array ini secara ascending. Cetak array setelah diurutkan.",
			TestCases:    [10]string{"[5 6 11 12 13]"},
			NumTestCases: 1,
		},
	}

	// Untuk mendapatkan jumlah soal yang valid, Anda perlu loop ini secara manual
	numValidSoal := 0
	i := 0
	for i < len(daftarSoal) {
		if daftarSoal[i].Materi != "" { // Asumsi soal valid jika Materi tidak kosong
			numValidSoal++
		}
		i++
	}

	fmt.Println()
	fmt.Println("=== PILIH SOAL ===")
	i = 0
	for i < numValidSoal { // Loop hanya sampai jumlah soal yang valid
		s := daftarSoal[i]
		fmt.Printf("%d. %s\n", i+1, s.Materi)
		i++
	}
	fmt.Println("0. Keluar")
	fmt.Print("Masukkan pilihan: ") // Menggunakan Print bukan Println agar input sebaris

	fmt.Println()
	fmt.Printf("Nilai saat ini: %.2f\n", *nilaiTemp)

	pilihanStr, _ := reader.ReadString('\n')
	pilihanStr = strings.TrimSpace(pilihanStr)

	pilihanInt := -1
	fmt.Sscanf(pilihanStr, "%d", &pilihanInt)

	if pilihanInt == 0 { // Jika pilih 0, keluar dari soal menu
		return false // Mengembalikan false untuk keluar dari fungsi ini
	}
	if pilihanInt < 1 || pilihanInt > numValidSoal { // Cek validitas pilihan
		fmt.Println("Pilihan tidak valid.")
		fmt.Println("Tekan ENTER untuk melanjutkan...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return StartSoalMenu(nilaiTemp) // Kembali ke menu soal jika pilihan tidak valid
	}

	soalTerpilih := daftarSoal[pilihanInt-1]

	// Auto tulis soal di file jawaban.go sebagai komentar
	jawabanContent := "// " + soalTerpilih.Pertanyaan + "\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\t// Tulis jawabanmu di sini\n}\n"
	os.WriteFile("jawaban.go", []byte(jawabanContent), 0644)
	fmt.Println("Membuka Notepad++ untuk menjawab soal...")
	time.Sleep(1 * time.Second)
	openNotepadPlusPlus("jawaban.go")

	var nilai float64 = 10.0 // Nilai awal untuk soal ini

	isAttemptingQuiz := true // Flag untuk mengontrol loop pengerjaan kuis

	for isAttemptingQuiz {
		fmt.Print("\nTekan ENTER untuk cek jawaban atau 'x' untuk keluar: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "x" {
			fmt.Println("Keluar dari mode soal...")
			os.Remove("jawaban.go")
			fmt.Print("\nTekan ENTER: ")
			bufio.NewReader(os.Stdin).ReadString('\n')
			isAttemptingQuiz = false // Hentikan loop pengerjaan kuis
			return false             // Mengembalikan false agar tidak kembali ke menu soal secara rekursif
		}

		output, _ := compileAndRun()

		isMatch := false // Flag untuk menunjukkan apakah output cocok dengan test case
		tcIndex := 0
		for tcIndex < soalTerpilih.NumTestCases && !isMatch {
			if strings.TrimSpace(output) == strings.TrimSpace(soalTerpilih.TestCases[tcIndex]) {
				isMatch = true
			}
			tcIndex++
		}

		if isMatch {
			fmt.Println("\n✅ Jawaban BENAR! Output sesuai test case.")
			if nilai < 0 { // Pastikan nilai tidak negatif
				nilai = 0
			}
			*nilaiTemp += nilai // Tambahkan nilai soal ini ke TotalScore siswa
			fmt.Printf("Nilai kamu untuk soal ini: %.2f\n", nilai)
			fmt.Printf("Total nilai sementara: %.2f\n", *nilaiTemp)

			// Setelah jawaban benar, berikan opsi kembali ke menu utama atau pilih soal lain
			fmt.Println("Mau pilih materi lain (y) atau kembali ke menu utama (n)?")

			isChoiceValid := false // Flag untuk validasi input Y/N setelah jawaban benar
			for !isChoiceValid {
				fmt.Scan(&Choice2)
				bufio.NewReader(os.Stdin).ReadString('\n') // Bersihkan buffer

				if strings.ToLower(Choice2) == "y" {
					os.Remove("jawaban.go")
					isChoiceValid = true
					isAttemptingQuiz = false        // Keluar dari loop pengerjaan kuis
					return StartSoalMenu(nilaiTemp) // Kembali ke menu soal (rekursif)
				} else if strings.ToLower(Choice2) == "n" {
					os.Remove("jawaban.go")
					isChoiceValid = true
					isAttemptingQuiz = false // Keluar dari loop pengerjaan kuis
					return true              // Mengembalikan true untuk kembali ke MainMenu di quiz package
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else {
			nilai -= nilai * 0.1 // Pengurangan 10%
			fmt.Println("\n❌ Output tidak sesuai dengan test case.")
			fmt.Printf("Nilai dikurangi 10%%, sisa nilai: %.2f\n", nilai)
			fmt.Println("Output kamu:\n", output)
			fmt.Println("Test case yang diharapkan:")
			tcPrintIndex := 0
			for tcPrintIndex < soalTerpilih.NumTestCases {
				fmt.Println("-", soalTerpilih.TestCases[tcPrintIndex])
				tcPrintIndex++
			}
		}
	}
	os.Remove("jawaban.go") // Pastikan file dihapus jika loop pengerjaan kuis selesai (misal karena 'x')
	return false            // Jika keluar dari loop pengerjaan kuis tanpa memilih y/n, anggap ingin keluar dari soal menu
}
