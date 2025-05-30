package pengpro

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

var Choice int
var Choice2 string

func MainMenu() {
	var berhenti bool
	berhenti = true
	atribut.ClearScreen()
	for berhenti {
	Submenu()
		if Choice == 0 {
			fmt.Println("Terima kasih telah menggunakan program PENGPRO!")
			berhenti = false
		}
	
}

func Submenu() {
	welcome := figure.NewFigure("PENGPRO", "doom", true).String()
	fmt.Print("\033[32m") // Set warna hijau
	fmt.Print(welcome)    // Cetak teks ASCII
	fmt.Print("\033[0m")
	fmt.Println("== Pengenalan Pemrograman ==")
	fmt.Println("1. Apa itu Pemrograman bahasa Go?")
	fmt.Println("2. Pengenalan tipe data(Number, String, Boolean)")
	fmt.Println("3. Variable dan Constant")
	fmt.Println("4. Konversi tipe data")
	fmt.Println("5. Operasi aritmatika dan logika")
	fmt.Println("6. Perulangan(For Loop, While Loop, Repeat Until)")
	fmt.Println("7. Percabangan(If Else, Switch Case)")
	fmt.Print("Pilih materi: ")
	fmt.Scan(&Choice)
	switch Choice {
	case 1:
		MateriGoLanguage() // Panggil fungsi utama materi Go
	case 2:
		MateriTipeData()
	case 3:
		MateriVariableConstant()
	case 4:
		MateriKonversiTipeData()
	case 5:
		MateriOperasiAritmatikaLogika()
	case 6:
		MateriPerulangan()
	case 7:
		MateriPercabangan()
	case 0:
		return
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		fmt.Println("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
		fmt.Scanln()
	}
}

func MateriGoLanguage() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
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
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizGoLanguage()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
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
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "google" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah Google.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Sebutkan salah satu fitur utama Go yang terkait dengan konkurensi (pemrosesan paralel).")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.Contains(strings.ToLower(jawaban), "goroutine") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah goroutine.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Go sering digunakan untuk pengembangan apa yang terkait dengan API dan layanan-layanan kecil?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
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

func MateriTipeData() {
	halamanSekarang := 1
	totalHalaman := 4
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Pengenalan Tipe Data (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam pemrograman, tipe data adalah klasifikasi yang memberi tahu kompiler atau interpreter bagaimana seorang programmer bermaksud menggunakan data tersebut.")
			fmt.Println("Setiap variabel yang Anda deklarasikan harus memiliki tipe data, yang menentukan jenis nilai apa yang dapat disimpan variabel itu.")
			fmt.Println("\nDi Go, ada beberapa tipe data dasar yang sering digunakan:")
			fmt.Println("1. Number (Angka)")
			fmt.Println("2. String (Teks)")
			fmt.Println("3. Boolean (Logika)")
		case 2:
			fmt.Println("--- 1. Tipe Data Number (Angka) ---")
			fmt.Println("Tipe data ini digunakan untuk menyimpan nilai numerik (angka).")
			fmt.Println("Ada dua kategori utama untuk angka:")
			fmt.Println("a. Integer (Bilangan Bulat)")
			fmt.Println("   Digunakan untuk angka tanpa desimal (..., -2, -1, 0, 1, 2, ...).")
			fmt.Println("   Go memiliki berbagai ukuran integer, seperti: ")
			fmt.Println("   - `int`: Ukuran tergantung arsitektur CPU (32-bit atau 64-bit). Ini yang paling umum.")
			fmt.Println("   - `int8`, `int16`, `int32`, `int64`: Integer dengan ukuran bit spesifik (signed).")
			fmt.Println("   - `uint8`, `uint16`, `uint32`, `uint64`: Integer tanpa tanda (unsigned, hanya positif).")
			fmt.Println("   Contoh: `var umur int = 30`, `var jumlahBarang int8 = 120`")
		case 3:
			fmt.Println("--- 1. Tipe Data Number (Angka) Lanjutan ---")
			fmt.Println("b. Floating-Point (Bilangan Desimal)")
			fmt.Println("   Digunakan untuk angka dengan bagian desimal.")
			fmt.Println("   Go memiliki dua tipe floating-point:")
			fmt.Println("   - `float32`: Presisi tunggal, cocok untuk kebutuhan yang tidak memerlukan akurasi tinggi.")
			fmt.Println("   - `float64`: Presisi ganda, ini adalah tipe float yang paling umum dan direkomendasikan karena akurasinya.")
			fmt.Println("   Contoh: `var harga float64 = 15500.75`, `var phi float32 = 3.14`")
			fmt.Println("\n--- 2. Tipe Data String (Teks) ---")
			fmt.Println("   Digunakan untuk menyimpan urutan karakter (teks).")
			fmt.Println("   Nilai string harus diapit oleh tanda kutip ganda (`\"`).")
			fmt.Println("   String di Go adalah *immutable*, artinya setelah dibuat, isinya tidak bisa diubah.")
			fmt.Println("   Contoh: `var namaDepan string = \"Budi\"`, `pesan := \"Selamat datang di Go!\"`")
		case 4:
			fmt.Println("--- 3. Tipe Data Boolean (Logika) ---")
			fmt.Println("   Tipe data ini sangat sederhana, hanya memiliki dua nilai: `true` (benar) atau `false` (salah).")
			fmt.Println("   Digunakan untuk menyimpan hasil dari operasi logika atau kondisi, seperti perbandingan.")
			fmt.Println("   Contoh: `var isLogin bool = true`, `isDewasa := (umur >= 18)`")
			fmt.Println("\nMemahami tipe data adalah dasar penting dalam pemrograman karena menentukan operasi apa yang bisa dilakukan pada data dan seberapa banyak memori yang akan dialokasikan.")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizTipeData()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

func QuizTipeData() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Pengenalan Tipe Data ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Tipe data yang digunakan untuk menampung bilangan bulat?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "integer" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah integer.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Jika kita memiliki sebuah variabel x bernilai 4.5, maka termasuk tipe data apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.Contains(strings.ToLower(jawaban), "float") || strings.Contains(strings.ToLower(jawaban), "desimal") || strings.Contains(strings.ToLower(jawaban), "floating-point") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah floating-point (atau float/desimal).")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Jika ingin menyatakan sebuah variabel true atau false, menggunakan tipe data?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.Contains(strings.ToLower(jawaban), "boolean") || strings.Contains(strings.ToLower(jawaban), "bool") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah boolean.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", score)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func MateriVariableConstant() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Variabel dan Konstanta (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam Go, variabel dan konstanta adalah cara untuk menyimpan data dalam program Anda.  Perbedaan utamanya adalah:")
			fmt.Println("\n--- Variabel ---")
			fmt.Println("Variabel adalah 'wadah' yang menyimpan nilai yang *bisa berubah* selama program berjalan.")
			fmt.Println("Anda harus mendeklarasikan variabel sebelum menggunakannya.")
			fmt.Println("Ada beberapa cara mendeklarasikan variabel di Go:")
			fmt.Println("\n1. Menggunakan `var`:")
			fmt.Println("   - `var namaTipe namaVariabel` (Deklarasi tanpa nilai awal)")
			fmt.Println("     Contoh: `var nama string`")
			fmt.Println("   - `var namaVariabel tipe = nilai` (Deklarasi dengan nilai awal)")
			fmt.Println("     Contoh: `var umur int = 30`")

		case 2:
			fmt.Println("--- Variabel Lanjutan ---")
			fmt.Println("   - `var namaVariabel = nilai` (Go menyimpulkan tipe data)")
			fmt.Println("     Contoh: `var pesan = \"Halo\"`")
			fmt.Println("\n2. Menggunakan *short variable declaration* (`:=`):")
			fmt.Println("   - Hanya bisa digunakan di dalam fungsi.")
			fmt.Println("   - `namaVariabel := nilai` (Go menyimpulkan tipe data)")
			fmt.Println("     Contoh: `nama := \"Budi\"`, `age := 30`")
			fmt.Println("\n--- Konstanta ---")
			fmt.Println("Konstanta adalah nilai yang *tidak bisa diubah* setelah dideklarasikan. Konstanta berguna untuk nilai-nilai yang diketahui dan tidak akan berubah, seperti nilai pi atau pesan tetap.")

		case 3:
			fmt.Println("--- Konstanta Lanjutan ---")
			fmt.Println("Anda mendeklarasikan konstanta menggunakan kata kunci `const`.")
			fmt.Println("Contoh:")
			fmt.Println("   - `const namaKonstanta tipe = nilai`")
			fmt.Println("     Contoh: `const PI float64 = 3.14159`")
			fmt.Println("   - `const namaKonstanta = nilai` (Go menyimpulkan tipe data)")
			fmt.Println("     Contoh: `const GREETING = \"Selamat datang\"`")
			fmt.Println("\nBeberapa aturan dan praktik terbaik:")
			fmt.Println("- Nama variabel dan konstanta harus deskriptif (memberi tahu apa yang disimpan).")
			fmt.Println("- Gunakan *camelCase* untuk nama variabel (misalnya, `namaDepan`, `jumlahBarang`).")
			fmt.Println("- Gunakan huruf kapital semua dengan garis bawah untuk konstanta (misalnya, `PI`, `MAX_UKURAN`).")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizVarConstant()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

func QuizVarConstant() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Variabel dan Konstanta ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut dengan singkat.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Wadah yang menyimpan nilai yang dapat berubah, disebut apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "variable" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah variable.")
	}

	fmt.Println("\n2. Jika kita ingin membuat variabel memiliki nilai yang tetap dan tidak dapat diubah, diperlukan perintah apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.Contains(strings.ToLower(jawaban), "const") || strings.Contains(strings.ToLower(jawaban), "konstanta") {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang benar adalah const atau konstanta.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 2 poin.\n", score)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func MateriKonversiTipeData() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Konversi Tipe Data (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Konversi tipe data adalah proses mengubah nilai dari satu tipe data ke tipe data lain.")
			fmt.Println("Ini sering diperlukan ketika Anda ingin melakukan operasi yang melibatkan tipe data yang berbeda, atau ketika Anda perlu menyimpan data dalam format tertentu.")
			fmt.Println("Go adalah bahasa yang *strongly typed*, yang berarti Anda harus secara eksplisit melakukan konversi tipe data (tidak otomatis di banyak kasus).")

			fmt.Println("\n--- Konversi Antar Tipe Angka ---")
			fmt.Println("Anda bisa mengkonversi antara tipe integer dan float.")
			fmt.Println("Syntax: `targetTipe(variabel)`")
			fmt.Println("Contoh:")
			fmt.Println("   `var i int = 10`")
			fmt.Println("   `var f float64 = float64(i)` // Konversi int ke float64 (hasil: 10.0)")
			fmt.Println("   `var k int = int(f)`       // Konversi float64 ke int (hasil: 10, bagian desimal hilang)")

		case 2:
			fmt.Println("--- Konversi String ke Number ---")
			fmt.Println("Ini seringkali memerlukan penanganan kesalahan karena string mungkin bukan format angka yang valid.")
			fmt.Println("Go menyediakan package `strconv` untuk tujuan ini.")
			fmt.Println("   - `strconv.Atoi(s string) (int, error)`: Konversi string ke integer.")
			fmt.Println("     Contoh:")
			fmt.Println("       `strNum := \"123\"`")
			fmt.Println("       `num, err := strconv.Atoi(strNum)` // num akan 123, err akan nil jika berhasil")
			fmt.Println("       `if err != nil { fmt.Println(\"Error konversi string ke int:\", err) }`")
			fmt.Println("   - `strconv.ParseFloat(s string, bitSize int) (float64, error)`: Konversi string ke float.")
			fmt.Println("     Contoh:")
			fmt.Println("       `strFloat := \"3.14\"`")
			fmt.Println("       `pi, err := strconv.ParseFloat(strFloat, 64)` // 64 untuk float64")

		case 3:
			fmt.Println("--- Konversi Number ke String ---")
			fmt.Println("Untuk mengubah angka menjadi string, Anda juga bisa menggunakan package `strconv`.")
			fmt.Println("   - `strconv.Itoa(i int) string`: Konversi integer ke string.")
			fmt.Println("     Contoh: `angka := 45; strAngka := strconv.Itoa(angka)` // strAngka akan \"45\"")
			fmt.Println("   - `strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string`:")
			fmt.Println("     Konversi float ke string. `fmt` bisa 'f', 'e', 'g'; `prec` jumlah digit desimal.")
			fmt.Println("     Contoh: `val := 12.34; strVal := strconv.FormatFloat(val, 'f', 2, 64)` // strVal akan \"12.34\"")
			fmt.Println("\nPenting untuk diingat bahwa konversi tipe data bisa menyebabkan kehilangan data (misalnya, saat mengkonversi float ke integer, bagian desimal akan hilang) atau kesalahan (jika format tidak valid).")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizKonversiTipeData()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

func QuizKonversiTipeData() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Konversi Tipe Data ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Untuk mengubah nilai `int` menjadi `float64`, Anda menulis `float64(nilaiInt)`. Ini disebut konversi tipe data apa?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "eksplisit" || strings.ToLower(jawaban) == "explicit" {
		fmt.Println("Benar! (atau secara eksplisit)")
		score++
	} else {
		fmt.Println("Salah. Ini adalah konversi eksplisit.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Package apa di Go yang digunakan untuk mengkonversi string menjadi angka (misalnya integer atau float)?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "strconv" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Package yang benar adalah `strconv`.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Jika Anda mengkonversi `float64(3.9)` menjadi `int`, berapa nilai yang akan dihasilkan?")
	fmt.Print("Jawaban Anda (angka): ")
	var jawabanInt int
	fmt.Scan(&jawabanInt)
	fmt.Scanln()
	if jawabanInt == 3 {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Nilai desimal akan dipotong, menjadi 3.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari 3 poin.\n", score)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func MateriOperasiAritmatikaLogika() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Operasi Aritmatika dan Logika (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam pemrograman, operasi aritmatika dan logika adalah dasar untuk melakukan perhitungan dan membuat keputusan.")
			fmt.Println("\n--- Operasi Aritmatika ---")
			fmt.Println("Digunakan untuk melakukan perhitungan matematis.")
			fmt.Println("Operator umum di Go:")
			fmt.Println("   - Penjumlahan: `+` (Contoh: `5 + 3 = 8`)")
			fmt.Println("   - Pengurangan: `-` (Contoh: `10 - 4 = 6`)")
			fmt.Println("   - Perkalian: `*` (Contoh: `2 * 6 = 12`)")
			fmt.Println("   - Pembagian: `/` (Contoh: `15 / 3 = 5`, `10 / 3 = 3` (untuk int), `10.0 / 3.0 = 3.33` (untuk float))")
			fmt.Println("   - Modulus: `%` (Sisa hasil bagi. Contoh: `10 % 3 = 1`)")

		case 2:
			fmt.Println("--- Operasi Aritmatika Lanjutan ---")
			fmt.Println("Prioritas operasi sama seperti matematika (perkalian/pembagian/modulus didahulukan sebelum penjumlahan/pengurangan).")
			fmt.Println("Anda bisa menggunakan tanda kurung `()` untuk mengubah prioritas.")
			fmt.Println("Contoh:")
			fmt.Println("   `hasil1 := 2 + 3 * 4`   // Hasil: 14 (3*4 dulu)")
			fmt.Println("   `hasil2 := (2 + 3) * 4` // Hasil: 20 (2+3 dulu)")

			fmt.Println("\n--- Operasi Logika ---")
			fmt.Println("Digunakan untuk menggabungkan atau memanipulasi nilai boolean (`true` atau `false`).")
			fmt.Println("Hasil dari operasi logika juga selalu `true` atau `false`.")
			fmt.Println("Operator umum di Go:")
			fmt.Println("   - AND: `&&` (True jika kedua operand true. Contoh: `(A && B)`)")
			fmt.Println("   - OR:  `||` (True jika salah satu operand true. Contoh: `(A || B)`)")
			fmt.Println("   - NOT: `!` (Membalik nilai boolean. Contoh: `!A`)")

		case 3:
			fmt.Println("--- Operasi Logika Lanjutan ---")
			fmt.Println("Contoh penggunaan:")
			fmt.Println("   `usia := 20`")
			fmt.Println("   `punyaSIM := true`")
			fmt.Println("   `bisaMengemudi := (usia >= 18 && punyaSIM)` // true jika usia >= 18 DAN punyaSIM true")
			fmt.Println("   `isLibur := true`")
			fmt.Println("   `goToSekolah := !isLibur` // false jika isLibur true")
			fmt.Println("\nMemahami operasi ini sangat penting untuk kontrol alur program (percabangan dan perulangan).")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizOperasiAritmatikaLogika()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

// --- FUNGSI KUIS KHUSUS MATERI OPERASI ARITMATIKA DAN LOGIKA ---
func QuizOperasiAritmatikaLogika() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Operasi Aritmatika dan Logika ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Berapa hasil dari `10 + 2 * 3`?")
	fmt.Print("Jawaban Anda (angka): ")
	var jawabanInt int
	fmt.Scan(&jawabanInt)
	fmt.Scanln()
	if jawabanInt == 16 { // 2*3 = 6, lalu 10+6 = 16
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Perhatikan prioritas operasi.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Operator Go yang digunakan untuk 'sisa hasil bagi' adalah?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if jawaban == "%" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Operatornya adalah `%`.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Jika `nilaiA` adalah `true` dan `nilaiB` adalah `false`, apa hasil dari `nilaiA && nilaiB`?")
	fmt.Print("Jawaban Anda (true/false): ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "false" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Operator AND (&&) membutuhkan kedua sisi bernilai true untuk menghasilkan true.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari %d poin.\n", score, 3)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func MateriPerulangan() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Perulangan (For Loop, While Loop, Repeat Until) (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Perulangan adalah struktur kontrol yang memungkinkan Anda menjalankan blok kode berulang kali.")
			fmt.Println("Ini sangat berguna ketika Anda perlu melakukan tugas yang sama berkali-kali tanpa harus menulis kode yang sama berulang-ulang.")
			fmt.Println("\nGo hanya memiliki satu konstruksi perulangan: `for`.")
			fmt.Println("Namun, `for` loop di Go sangat fleksibel dan dapat digunakan untuk mengimplementasikan berbagai jenis perulangan yang ditemukan di bahasa lain.")
			fmt.Println("\n--- 1. `for` Loop Klasik (Seperti C/Java) ---")
			fmt.Println("Paling umum untuk mengulang sejumlah kali yang diketahui.")
			fmt.Println("Syntax: `for inisialisasi; kondisi; post-statement { ...kode... }`")
			fmt.Println("Contoh:")
			fmt.Println("   `for i := 0; i < 5; i++ {`")
			fmt.Println("   `   fmt.Println(\"Iterasi ke-\", i)`")
			fmt.Println("   `}`")

		case 2:
			fmt.Println("--- 2. `for` Loop sebagai `while` Loop ---")
			fmt.Println("Digunakan ketika Anda ingin mengulang selama suatu kondisi tertentu masih benar, dan jumlah iterasi tidak diketahui di awal.")
			fmt.Println("Syntax: `for kondisi { ...kode... }`")
			fmt.Println("Contoh:")
			fmt.Println("   `nilai := 0`")
			fmt.Println("   `for nilai < 5 {`")
			fmt.Println("   `   fmt.Println(\"Nilai sekarang:\", nilai)`")
			fmt.Println("   `   nilai++`")
			fmt.Println("   `}`")

			fmt.Println("\n--- 3. `for` Loop Tanpa Kondisi (Infinite Loop / Seperti `repeat until`) ---")
			fmt.Println("Perulangan ini akan berjalan terus-menerus kecuali jika Anda secara eksplisit menghentikannya dengan `break` atau `return`.")
			fmt.Println("Syntax: `for { ...kode... }`")
			fmt.Println("Contoh:")
			fmt.Println("   `count := 0`")
			fmt.Println("   `for {`")
			fmt.Println("   `   fmt.Println(\"Loop tak terbatas...\")`")
			fmt.Println("   `   count++`")
			fmt.Println("   `   if count >= 3 {`")
			fmt.Println("   `       break // Hentikan loop setelah 3 iterasi`")
			fmt.Println("   `   }`")
			fmt.Println("   `}`")

		case 3:
			fmt.Println("--- Kata Kunci Penting dalam Perulangan ---")
			fmt.Println("   - `break`: Menghentikan eksekusi loop saat ini dan melanjutkan kode setelah loop.")
			fmt.Println("   - `continue`: Melewatkan sisa iterasi saat ini dan melanjutkan ke iterasi berikutnya dari loop.")
			fmt.Println("Contoh `continue`:")
			fmt.Println("   `for i := 0; i < 5; i++ {`")
			fmt.Println("   `   if i == 2 {`")
			fmt.Println("   `       continue // Lewati iterasi saat i adalah 2`")
			fmt.Println("   `   }`")
			fmt.Println("   `   fmt.Println(\"Angka:\", i)`")
			fmt.Println("   `}`")
			fmt.Println("   // Output: Angka: 0, Angka: 1, Angka: 3, Angka: 4")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizPerulangan()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

// --- FUNGSI KUIS KHUSUS MATERI PERULANGAN ---
func QuizPerulangan() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Perulangan ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Berapa kali kode di bawah ini akan mencetak 'Halo'?")
	fmt.Println("   `for i := 0; i < 3; i++ { fmt.Println(\"Halo\") }`")
	fmt.Print("Jawaban Anda (angka): ")
	var jawabanInt int
	fmt.Scan(&jawabanInt)
	fmt.Scanln()
	if jawabanInt == 3 {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Perulangan akan berjalan 3 kali (i=0, 1, 2).")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Keyword apa yang digunakan untuk menghentikan loop sepenuhnya?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "break" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Keyword yang benar adalah `break`.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Go hanya memiliki satu jenis konstruksi perulangan. Apa keyword yang digunakan?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "for" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Keyword yang benar adalah `for`.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari %d poin.\n", score, 3)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func MateriPercabangan() {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Percabangan (If Else, Switch Case) (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Percabangan adalah struktur kontrol yang memungkinkan program Anda membuat keputusan dan menjalankan blok kode yang berbeda berdasarkan kondisi tertentu.")
			fmt.Println("Ini fundamental untuk menciptakan logika yang dinamis dan adaptif dalam aplikasi Anda.")

			fmt.Println("\n--- 1. `if`, `else if`, dan `else` ---")
			fmt.Println("Digunakan untuk mengeksekusi blok kode jika suatu kondisi benar, dan blok lain jika kondisi salah.")
			fmt.Println("Syntax:")
			fmt.Println("   `if kondisi1 {`")
			fmt.Println("   `   // Kode jika kondisi1 benar`")
			fmt.Println("   `} else if kondisi2 {`")
			fmt.Println("   `   // Kode jika kondisi1 salah TAPI kondisi2 benar`")
			fmt.Println("   `} else {`")
			fmt.Println("   `   // Kode jika semua kondisi di atas salah`")
			fmt.Println("   `}`")
			fmt.Println("\nContoh:")
			fmt.Println("   `nilai := 80`")
			fmt.Println("   `if nilai >= 75 {`")
			fmt.Println("   `   fmt.Println(\"Selamat, Anda Lulus!\")`")
			fmt.Println("   `} else {`")
			fmt.Println("   `   fmt.Println(\"Maaf, Anda Gagal.\")`")
			fmt.Println("   `}`")

		case 2:
			fmt.Println("--- 2. `switch` Statement ---")
			fmt.Println("Menyediakan cara yang lebih ringkas dan bersih untuk menangani beberapa kondisi dibandingkan dengan serangkaian `if else if` yang panjang.")
			fmt.Println("Go's `switch` tidak memerlukan `break` secara eksplisit antar `case` (ini adalah 'fallthrough' secara implisit hanya jika Anda menambahkannya).")
			fmt.Println("Syntax:")
			fmt.Println("   `switch ekspresi {`")
			fmt.Println("   `case nilai1:`")
			fmt.Println("   `   // Kode jika ekspresi == nilai1`")
			fmt.Println("   `case nilai2, nilai3:`") // Multiple values in one case
			fmt.Println("   `   // Kode jika ekspresi == nilai2 atau nilai3`")
			fmt.Println("   `default:`")
			fmt.Println("   `   // Kode jika tidak ada case yang cocok`")
			fmt.Println("   `}`")
			fmt.Println("\nContoh:")
			fmt.Println("   `hari := \"Selasa\"`")
			fmt.Println("   `switch hari {`")
			fmt.Println("   `case \"Senin\":`")
			fmt.Println("   `   fmt.Println(\"Awal Pekan, semangat!\")`")
			fmt.Println("   `case \"Sabtu\", \"Minggu\":`")
			fmt.Println("   `   fmt.Println(\"Waktunya libur!\")`")
			fmt.Println("   `default:`")
			fmt.Println("   `   fmt.Println(\"Hari kerja biasa.\")`")
			fmt.Println("   `}`")

		case 3:
			fmt.Println("--- `switch` Tanpa Ekspresi (Mirip `if/else if`) ---")
			fmt.Println("Anda bisa menggunakan `switch` tanpa ekspresi, di mana setiap `case` adalah kondisi boolean.")
			fmt.Println("Ini sering membuat kode lebih mudah dibaca daripada serangkaian `if-else if`.")
			fmt.Println("Contoh:")
			fmt.Println("   `umur := 18`")
			fmt.Println("   `switch {`")
			fmt.Println("   `case umur < 13:`")
			fmt.Println("   `   fmt.Println(\"Anak-anak\")`")
			fmt.Println("   `case umur >= 13 && umur < 20:`")
			fmt.Println("   `   fmt.Println(\"Remaja\")`")
			fmt.Println("   `default:`")
			fmt.Println("   `   fmt.Println(\"Dewasa\")`")
			fmt.Println("   `}`")

			fmt.Println("\n--- Pentingnya Percabangan ---")
			fmt.Println("Percabangan memungkinkan program Anda untuk menjadi 'pintar', bereaksi terhadap input pengguna, status sistem, atau data eksternal.")
			fmt.Println("Ini adalah blok bangunan esensial untuk hampir setiap aplikasi yang kompleks.")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "y" {
				halamanSekarang++
			} else {
				berhenti = false
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke kuis (k) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice2)
			fmt.Scanln()

			if strings.ToLower(Choice2) == "k" {
				QuizPercabangan()
				berhenti = false
			} else {
				berhenti = false
			}
		}
	}
	return
}

func QuizPercabangan() {
	atribut.ClearScreen()
	fmt.Println("=== Kuis: Percabangan (If Else, Switch Case) ===")
	fmt.Println("Jawablah pertanyaan-pertanyaan berikut.")

	score := 0
	var jawaban string

	// Pertanyaan 1
	fmt.Println("\n1. Struktur kontrol apa yang digunakan untuk mengeksekusi blok kode jika suatu kondisi **benar**, dan blok lain jika kondisi **salah**?")
	fmt.Println("   a. For Loop")
	fmt.Println("   b. If Else")
	fmt.Println("   c. Switch Case")
	fmt.Println("   d. GoTo")
	fmt.Print("Jawaban Anda (a/b/c/d): ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "b" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Jawaban yang tepat adalah If Else.")
	}

	// Pertanyaan 2
	fmt.Println("\n2. Dalam Go, apakah Anda perlu menggunakan keyword `break` secara eksplisit di setiap `case` dalam `switch` statement untuk mencegah 'fallthrough' (lanjut ke case berikutnya)?")
	fmt.Println("   a. Ya, selalu")
	fmt.Println("   b. Tidak, Go secara implisit menambahkan break")
	fmt.Println("   c. Tergantung versi Go")
	fmt.Println("   d. Hanya jika ada multiple values di satu case")
	fmt.Print("Jawaban Anda (a/b/c/d): ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "b" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Go secara implisit menambahkan `break` untuk setiap `case`. Anda perlu menggunakan `fallthrough` jika ingin perilaku sebaliknya.")
	}

	// Pertanyaan 3
	fmt.Println("\n3. Perhatikan kode berikut:")
	fmt.Println("   ```go")
	fmt.Println("   umur := 15")
	fmt.Println("   switch {")
	fmt.Println("   case umur < 10:")
	fmt.Println("       fmt.Println(\"Anak-anak\")")
	fmt.Println("   case umur >= 10 && umur < 20:")
	fmt.Println("       fmt.Println(\"Remaja\")")
	fmt.Println("   default:")
	fmt.Println("       fmt.Println(\"Dewasa\")")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   Apa output dari kode di atas?")
	fmt.Print("Jawaban Anda: ")
	fmt.Scan(&jawaban)
	fmt.Scanln()
	if strings.ToLower(jawaban) == "remaja" {
		fmt.Println("Benar!")
		score++
	} else {
		fmt.Println("Salah. Umur 15 masuk dalam kategori 'Remaja'.")
	}

	fmt.Println("\n==============================")
	fmt.Printf("Kuis selesai! Anda mendapatkan %d dari %d poin.\n", score, 3)
	fmt.Println("==============================")

	PromptKembaliToMain()
}

func PromptKembaliToMain() {
	fmt.Println("\n------------------------------------")
	fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
	fmt.Scan(&Choice2)
	fmt.Scanln() // Consume the newline after reading Choice2
	if strings.ToLower(Choice2) == "y" {
		// Do nothing, the outer loop in main will call Submenu() again
	} else {
		Choice = 0 // Set Choice to 0 to exit the main loop
	}
}
