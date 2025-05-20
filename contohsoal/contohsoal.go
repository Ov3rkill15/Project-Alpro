package contohsoal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var Choice2 string

type Soal struct {
	Pertanyaan string   `json:"pertanyaan"`
	TestCases  []string `json:"test_cases"`
}

func loadSoal(path string) (*Soal, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var soal Soal
	err = json.Unmarshal(data, &soal)
	if err != nil {
		return nil, err
	}
	return &soal, nil
}

func openNotepadPlusPlus(filePath string) {
	exePath := "C:\\Program Files\\Notepad++\\notepad++.exe"
	cmd := exec.Command(exePath, filePath)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Gagal membuka Notepad++:", err)
	}
}

func compileAndRun() (string, error) {
	cmd := exec.Command("go", "run", "jawaban.go")
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func MainMenu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("=== PILIH SOAL ===")
	fmt.Println("1. Soal Pengantar")
	fmt.Println("2. Soal Algoritma")
	fmt.Print("Masukkan pilihan: ")
	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	var soalPath string
	switch pilihan {
	case "1":
		soalPath = "D:\\Matkul smester 2\\Project-Alpro\\contohsoal\\soal_pengpro.json"
	case "2":
		soalPath = "D:\\Matkul smester 2\\Project-Alpro\\contohsoal\\soal_alpro.json"
	default:
		fmt.Println("Pilihan tidak valid")
	}

	soal, err := loadSoal(soalPath)
	if err != nil {
		fmt.Println("Gagal memuat soal:", err)
		return
	}

	// Auto tulis soal di file jawaban.go sebagai komentar
	jawabanContent := "// " + soal.Pertanyaan + "\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\t// Tulis jawabanmu di sini\n}\n"
	os.WriteFile("jawaban.go", []byte(jawabanContent), 0644)
	defer os.Remove("jawaban.go")
	fmt.Println("Membuka Notepad++ untuk menjawab soal...")
	time.Sleep(1 * time.Second)
	openNotepadPlusPlus("jawaban.go")

	for {
		fmt.Print("\nTekan ENTER untuk cek jawaban atau 'x' untuk keluar: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "x" {
			fmt.Println("Keluar dari mode soal...")
			return
		}

		output, err := compileAndRun()
		if err != nil {
			fmt.Println("Terjadi error saat menjalankan kode:\n", err)
			continue
		}

		match := false
		for _, tc := range soal.TestCases {
			if strings.Contains(output, tc) {
				match = true
				break
			}
		}

		if match {
			fmt.Println("\n✅ Jawaban BENAR! Output sesuai test case.")
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			defer os.Remove("jawaban.go")
			fmt.Scan(&Choice2)
			bufio.NewReader(os.Stdin).ReadString('\n')
			if Choice2 == "y" {
				MainMenu()
			} else {
				return
			}
		} else {
			fmt.Println("\n❌ Output tidak sesuai dengan test case.")
			fmt.Println("Output kamu:\n", output)
			fmt.Println("Test case yang diharapkan:")
			for _, t := range soal.TestCases {
				fmt.Println("-", t)
			}
		}
	}
}
