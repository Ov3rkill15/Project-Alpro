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
	Judul     string   `json:"judul"`
	Deskripsi string   `json:"deskripsi"`
	TestCases []string `json:"test_cases"`
	Outputs   []string `json:"outputs"`
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

func tulisSoalKeJawaban(path string, soal *Soal) error {
	template := fmt.Sprintf(`// Judul: %s
// Deskripsi: %s

package main

import "fmt"

func main() {
	// Tulis jawaban kamu di sini
}
`, soal.Judul, soal.Deskripsi)

	return os.WriteFile(path, []byte(template), 0644)
}

func MainMenu() {
	soalPath := `D:\Matkul smester 2\New folder\soal1.json`
	jawabanPath := `D:\Matkul smester 2\New folder\jawaban.go`

	soal, err := loadSoal(soalPath)
	if err != nil {
		panic("Gagal membaca file soal: " + err.Error())
	}

	err = tulisSoalKeJawaban(jawabanPath, soal)
	if err != nil {
		fmt.Println("Gagal menulis ke jawaban.go:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Buka Notepad
		fmt.Println("\nMembuka jawaban.go di Notepad...")
		exec.Command("notepad", jawabanPath).Run()

		// Tunggu input Enter dari user
		fmt.Print("Tekan ENTER setelah kamu selesai mengerjakan dan menutup Notepad...")
		_, _ = reader.ReadString('\n') // Tidak perlu menggunakan variabel 'err'

		// Jalankan kode
		cmd := exec.Command("go", "run", jawabanPath)
		output, err := cmd.CombinedOutput()

		// Tidak gunakan err jika tidak penting
		if err != nil {
			fmt.Println("Ada error saat menjalankan kode:", err)
		}

		cleanOutput := strings.TrimSpace(string(output))
		cleanExpected := strings.TrimSpace(strings.Join(soal.Outputs, "\n"))

		fmt.Println("\nOutput dari program kamu:")
		fmt.Println(cleanOutput)

		// Cek apakah output sudah benar
		if cleanOutput == cleanExpected {
			fmt.Println("\n✅ Jawaban BENAR! Selamat!")
			fmt.Println("Tetap di soal ini(y), Menu soal lain(r) atau kembali ke menu utama(m)?")
			fmt.Scan(&Choice2)
			for Choice2 == "y" {
				MainMenu()
				fmt.Println("Tetap di Fibonacci(y), Menu rekursif lain(r) atau kembali ke menu utama(m)?")
				fmt.Scan(&Choice2)
			}
			if Choice2 == "r" {
				clearScreen()
				return
			} else if Choice2 == "m" {
				clearScreen()
				fmt.Println("Kembali ke menu utama...")
				for range 10 {
					time.Sleep(100 * time.Millisecond)
					fmt.Print(".")
				}
				return
			}
		} else {
			fmt.Println("\n❌ Jawaban SALAH! Coba lagi ya.")
			fmt.Println("Output yang diharapkan:")
			fmt.Println(cleanExpected)
			fmt.Println("\n----------------------------------------")
			fmt.Print("Tekan ENTER untuk membuka kembali Notepad dan perbaiki jawaban...")
			_, _ = reader.ReadString('\n') // Tidak perlu menggunakan variabel 'err'
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
