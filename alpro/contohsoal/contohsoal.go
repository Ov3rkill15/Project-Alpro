package contohsoal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Choice2 adalah variabel global untuk menyimpan pilihan pengguna setelah jawaban benar.
const nmax int = 5

// Soal merepresentasikan struktur soal dengan pertanyaan dan daftar test case.
type Soal struct {
	Materi     string       // Menambahkan field Materi untuk deskripsi soal
	Pertanyaan string       `json:"pertanyaan"`
	TestCases  [nmax]string `json:"test_cases"` // Menggunakan slice untuk daftar test case, ini adalah cara standar di Go.
}

// openNotepadPlusPlus membuka file yang diberikan menggunakan Notepad++.
// Pastikan path ke notepad++.exe sudah benar di sistem Anda.
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

// MainMenu adalah fungsi utama yang mengelola alur soal dan pengecekan jawaban.
func Mainmenu() {
	var Choice2 string
	reader := bufio.NewReader(os.Stdin)

	// Definisi semua soal secara langsung di dalam kode Go
	daftarSoal := [10]Soal{
		{
			Materi:     "Function (Sederhana)",
			Pertanyaan: "Buatlah sebuah fungsi Go bernama `hitungJumlah` yang menerima dua bilangan bulat ($a$ dan $b$) sebagai input, lalu mengembalikan hasil penjumlahan kedua bilangan tersebut. Cetak hasil pemanggilan fungsi ini dengan input 5 dan 3.",
			TestCases:  [nmax]string{"8"},
		},
		{
			Materi:     "Procedure (Sederhana)",
			Pertanyaan: "Buatlah sebuah fungsi Go bernama `sapaPengguna` yang menerima sebuah string (nama) sebagai input dan mencetak \"Halo, [nama]!\" ke konsol. Panggil fungsi ini dengan nama \"Budi\".",
			TestCases:  [nmax]string{"Halo, Budi!"},
		},
		{
			Materi:     "Rekursif (Sederhana)",
			Pertanyaan: "Buatlah sebuah fungsi rekursif Go bernama `faktorial` yang menghitung nilai faktorial dari sebuah bilangan bulat non-negatif $n$. Cetak hasil faktorial dari 4.",
			TestCases:  [nmax]string{"24"},
		},
		{
			Materi:     "Tipe Bentukan (Menengah)",
			Pertanyaan: "Definisikan sebuah `struct` Go bernama `Mahasiswa` dengan field `Nama` (string) dan `NIM` (string). Buat sebuah instance `Mahasiswa` dengan Nama \"Dewi\" dan NIM \"123456789\". Cetak `NIM` dari instance tersebut.",
			TestCases:  [nmax]string{"123456789"},
		},
		{
			Materi:     "Array (Menengah)",
			Pertanyaan: "Buatlah sebuah array Go berukuran 5 yang berisi bilangan bulat: 10, 20, 30, 40, 50. Cetak elemen pada indeks ke-2 dari array tersebut.",
			TestCases:  [nmax]string{"30"},
		},
		{
			Materi:     "Pencarian Nilai Max (Menengah)",
			Pertanyaan: "Diberikan sebuah array bilangan bulat: [15, 7, 22, 10, 30, 5]. Temukan dan cetak nilai maksimum dalam array tersebut.",
			TestCases:  [nmax]string{"30"},
		},
		{
			Materi:     "Sequential Search (Menengah)",
			Pertanyaan: "Diberikan sebuah array bilangan bulat: [5, 12, 8, 20, 15, 3]. Lakukan pencarian sekuensial untuk nilai 20. Jika ditemukan, cetak \"Nilai ditemukan!\". Jika tidak, cetak \"Nilai tidak ditemukan.\".",
			TestCases:  [nmax]string{"Nilai ditemukan!"},
		},
		{
			Materi:     "Binary Search (Menengah)",
			Pertanyaan: "Diberikan sebuah array bilangan bulat yang sudah terurut: [10, 20, 30, 40, 50, 60]. Lakukan pencarian biner untuk nilai 40. Jika ditemukan, cetak \"Nilai ditemukan!\". Jika tidak, cetak \"Nilai tidak ditemukan.\".",
			TestCases:  [nmax]string{"Nilai ditemukan!"},
		},
		{
			Materi:     "Selection Sort (Menengah)",
			Pertanyaan: "Diberikan sebuah array bilangan bulat: [64, 25, 12, 22, 11]. Terapkan algoritma Selection Sort untuk mengurutkan array ini secara ascending. Cetak array setelah diurutkan.",
			TestCases:  [nmax]string{"[11 12 22 25 64]"},
		},
		{
			Materi:     "Insertion Sort (Menengah)",
			Pertanyaan: "Diberikan sebuah array bilangan bulat: [12, 11, 13, 5, 6]. Terapkan algoritma Insertion Sort untuk mengurutkan array ini secara ascending. Cetak array setelah diurutkan.",
			TestCases:  [nmax]string{"[5 6 11 12 13]"},
		},
	}

	fmt.Println()
	fmt.Println("=== PILIH SOAL ===")
	i := 0
	for i < len(daftarSoal) {
		s := daftarSoal[i]
		fmt.Printf("%d. %s\n", i+1, s.Materi)
		i++
	}
	fmt.Print("Masukkan pilihan: ")
	pilihanStr, _ := reader.ReadString('\n')
	pilihanStr = strings.TrimSpace(pilihanStr)

	pilihanInt := -1
	fmt.Sscanf(pilihanStr, "%d", &pilihanInt)

	if pilihanInt < 1 || pilihanInt > len(daftarSoal) {
		fmt.Println("Pilihan tidak valid")
		return
	}

	soalTerpilih := daftarSoal[pilihanInt-1]

	// Auto tulis soal di file jawaban.go sebagai komentar
	jawabanContent := "// " + soalTerpilih.Pertanyaan + "\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\t// Tulis jawabanmu di sini\n}\n"
	os.WriteFile("jawaban.go", []byte(jawabanContent), 0644)
	fmt.Println("Membuka Notepad++ untuk menjawab soal...")
	time.Sleep(1 * time.Second)
	openNotepadPlusPlus("jawaban.go")

	for {
		fmt.Print("\nTekan ENTER untuk cek jawaban atau 'x' untuk keluar: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "x" {
			fmt.Println("Keluar dari mode soal...")
			os.Remove("jawaban.go") // Hapus file jawaban.go saat keluar.
			fmt.Print("\nTekan ENTER: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

		}

		output, _ := compileAndRun()
		match := false
		for i := 0; i < len(soalTerpilih.TestCases); i++ {
			tc := soalTerpilih.TestCases[i]
			if strings.Contains(output, tc) {
				match = true
			}
		}

		if match {
			fmt.Println("\n✅ Jawaban BENAR! Output sesuai test case.")
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			fmt.Scan(&Choice2)
			// Membersihkan buffer stdin setelah fmt.Scan
			bufio.NewReader(os.Stdin).ReadString('\n')
			if strings.ToLower(Choice2) == "y" {
				os.Remove("jawaban.go") // Hapus file sebelum memanggil MainMenu lagi.
				Mainmenu()
			} else {
				os.Remove("jawaban.go") // Hapus file jawaban.go saat kembali ke menu utama.
				return
			}
		} else {
			fmt.Println("\n❌ Output tidak sesuai dengan test case.")
			fmt.Println("Output kamu:\n", output)
			fmt.Println("Test case yang diharapkan:")
			for _, t := range soalTerpilih.TestCases {
				fmt.Println("-", t)
			}
		}
	}
}
