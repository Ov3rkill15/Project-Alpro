package p_function

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

func MainMenu() {
	var pilihan string
	var pilihan2 string

	fmt.Println(`
=====================================
Selamat Datang di Pembelajaran Function
=====================================
1. Apa itu Function
0. Keluar
    `)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Scanln()
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarFunction() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			berhenti2 := false
			for !berhenti2 {
				fmt.Scan(&pilihan2)
				fmt.Scanln()
				if strings.ToLower(pilihan2) == "y" {
					MainMenu()
				} else if strings.ToLower(pilihan2) == "n" {
					return
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
					fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
				}
			}
		} else {
			return
		}
	case "0":
		fmt.Println("Terima kasih telah menggunakan program pembelajaran function.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '3'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
		MainMenu()
	}
}

func belajarFunction() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true
	var Choice string

	for berhenti {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Sub-Function di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan kamu seorang koki ahli yang sedang menyiapkan hidangan kompleks. Daripada menulis setiap langkah kecil untuk setiap hidangan dari awal (misalnya, 'potong bawang', 'tumis bawang putih', 'rebus air'), kamu mungkin punya resep mini untuk tugas-tugas umum. Misalnya, 'Saus Tomat Dasar' atau 'Sayuran Panggang Sempurna'. Kamu menyiapkan ini sekali, lalu tinggal merujuknya kapan pun hidangan membutuhkan.")
			fmt.Println("\nDalam pemrograman, sebuah **fungsi** itu persis seperti resep mini itu. Ini adalah blok kode mandiri yang dirancang untuk melakukan tugas spesifik, yang didefinisikan dengan baik. Kamu menulis kode untuk tugas itu sekali di dalam fungsi, memberinya nama, lalu kamu bisa **memanggil** (atau mengeksekusi) fungsi itu kapan pun dan di mana pun kamu membutuhkannya dalam programmu.")
			fmt.Println("\n**Kenapa Fungsi Sangat Penting?**")
			fmt.Println("- **Reusabilitas:** Ini keuntungan besar! Kamu menulis kode sekali, menyimpannya dalam fungsi, dan kemudian bisa menggunakannya berkali-kali di seluruh programmu. Ini menghemat waktu pengembangan, mengurangi kemungkinan kesalahan, dan membuat kodemu lebih efisien.")
			fmt.Println("- **Modularitas & Organisasi:** Fungsi membantumu memecah masalah besar dan kompleks menjadi bagian-bagian yang lebih kecil dan mudah dikelola. Setiap fungsi bisa fokus melakukan satu hal dengan baik. Ini membuat kodemu jauh lebih mudah dipahami, di-debug, dan dipelihara.")
			fmt.Println("- **Keterbacaan:** Fungsi dengan nama yang baik membuat kodemu hampir seperti dokumentasi diri sendiri. Saat kamu melihat `hitungTotalHarga()` atau `kirimNotifikasiEmail()`, kamu langsung tahu apa yang dimaksudkan oleh bagian kode itu, bahkan tanpa melihat detail internalnya.")
		case 2:
			fmt.Println("Mari kita lihat struktur dasar sebuah fungsi di Go. Ini cukup lugas:")
			fmt.Println("\n```go")
			fmt.Println("func namaFungsi(parameter1 tipeData, parameter2 tipeData) (nilaiKembali1 tipeData, nilaiKembali2 tipeData) {")
			fmt.Println("    // Ini adalah badan fungsi")
			fmt.Println("    // Semua kode yang melakukan tugas fungsi ada di sini")
			fmt.Println("    // ...")
			fmt.Println("    return nilai1, nilai2 // Jika fungsi mengembalikan nilai")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nInilah arti setiap bagian:")
			fmt.Println("- **`func` keyword:** Ini cara kamu mendeklarasikan fungsi di Go.")
			fmt.Println("- **`namaFungsi`:** Nama unik yang kamu berikan pada fungsimu. Pilih nama yang deskriptif!")
			fmt.Println("- **`()` (Tanda Kurung):** Ini untuk **parameter** (input). Parameter adalah 'bahan' yang kamu berikan agar fungsi bisa bekerja. Jika fungsi tidak butuh input, tanda kurungnya tetap kosong.")
			fmt.Println("- **`(nilaiKembali1 tipeData)` (Return Values):** Ini opsional, yang menentukan nilai apa yang akan **dikembalikan** fungsi setelah selesai. Fungsi Go bisa mengembalikan nol, satu, atau bahkan beberapa nilai sekaligus!")
			fmt.Println("- **`{}` (Kurung Kurawal):** Ini adalah **badan fungsi**, tempat kamu menulis kode sebenarnya yang melakukan tugas fungsi.")
			fmt.Println("\n**Contoh Sederhana:**")
			fmt.Println("```go")
			fmt.Println("func sapaPengguna(nama string) { // sapaPengguna adalah fungsi, 'nama' adalah parameter string")
			fmt.Println("    fmt.Printf(\"Halo, (%)s! Selamat datang di program kami.\\n\", nama)")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggil (menggunakan) fungsi ini:")
			fmt.Println("// sapaPengguna(\"Budi\") // Ini akan mencetak \"Halo, Budi! Selamat datang di program kami.\"")
			fmt.Println("// sapaPengguna(\"Ani\")  // Ini akan mencetak \"Halo, Ani! Selamat datang di program kami.\"")
			fmt.Println("```")
		case 3:
			fmt.Println("Fungsi adalah tulang punggung dari setiap program Go yang terstruktur dengan baik. Mereka memungkinkan kita membangun aplikasi kompleks dari unit-unit kecil yang mudah diuji dan dikelola.")
			fmt.Println("\nGo memiliki beberapa fitur yang sangat kuat terkait fungsi:")
			fmt.Println("- **Multiple Return Values:** Salah satu fitur unggulan Go adalah kemampuannya mengembalikan lebih dari satu nilai dari satu fungsi. Ini sangat berguna, seringkali memungkinkanmu mengembalikan hasil dan juga status *error* dari sebuah operasi sekaligus.")
			fmt.Println("- **Functions as First-Class Citizens:** Di Go, fungsi bisa diperlakukan seperti variabel lainnya. Kamu bisa meneruskan fungsi sebagai argumen ke fungsi lain, menyimpannya di variabel, atau bahkan mengembalikannya dari fungsi lain! Ini membuka pola desain yang kuat seperti *callback* dan *higher-order functions*.")
			fmt.Println("- **`defer` Statements:** Kamu bisa menggunakan kata kunci `defer` di dalam fungsi untuk memastikan bahwa panggilan fungsi akan dieksekusi tepat sebelum fungsi yang mengelilinginya selesai, terlepas dari bagaimana fungsi tersebut keluar (normal atau karena *error*). Ini umum digunakan untuk membersihkan sumber daya, seperti menutup *file* atau koneksi *database*.")
			fmt.Println("\nDengan memahami dan menguasai fungsi, kamu akan bisa menulis kode Go yang tidak hanya efisien, tetapi juga mudah dibaca, dipelihara, dan diskalakan. Fungsi adalah salah satu konsep paling fundamental yang harus kamu pahami saat belajar Go!")
			fmt.Println("\n---")

			fmt.Println("### Contoh Penggunaan Fungsi Sederhana")
			fmt.Println("\nMari kita lihat contoh fungsi sederhana yang menghitung luas persegi panjang:")
			fmt.Println("```go")
			fmt.Println("package main")
			fmt.Println("\nimport \"fmt\"")
			fmt.Println("\n// hitungLuasPersegiPanjang adalah fungsi yang menerima dua parameter (panjang dan lebar)")
			fmt.Println("// dan mengembalikan satu nilai (luas) bertipe int.")
			fmt.Println("func hitungLuasPersegiPanjang(panjang int, lebar int) int {")
			fmt.Println("    luas := panjang * lebar")
			fmt.Println("    return luas")
			fmt.Println("}")
			fmt.Println("\nfunc main() {")
			fmt.Println("    // Memanggil fungsi hitungLuasPersegiPanjang")
			fmt.Println("    luas1 := hitungLuasPersegiPanjang(10, 5)")
			fmt.Println("    fmt.Printf(\"Luas persegi panjang dengan panjang 10 dan lebar 5 adalah: (%)d\\n\", luas1) // Output: 50")
			fmt.Println("\n    luas2 := hitungLuasPersegiPanjang(7, 3)")
			fmt.Println("    fmt.Printf(\"Luas persegi panjang dengan panjang 7 dan lebar 3 adalah: (%)d\\n\", luas2) // Output: 21")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDalam contoh di atas:")
			fmt.Println("* Kita mendefinisikan fungsi bernama `hitungLuasPersegiPanjang` yang menerima dua *integer* sebagai **input** (`panjang` dan `lebar`).")
			fmt.Println("* Fungsi ini **mengembalikan** satu *integer* yang merupakan hasil perkalian `panjang` dan `lebar`.")
			fmt.Println("* Di dalam fungsi `main`, kita **memanggil** `hitungLuasPersegiPanjang` beberapa kali dengan nilai yang berbeda, menunjukkan bagaimana fungsi dapat digunakan kembali.")
			fmt.Println("\nIni adalah contoh dasar bagaimana fungsi membantu kita mengorganisir kode dan menghindari pengulangan.")
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln()

				if strings.ToLower(Choice) == "y" {
					halamanSekarang++
					berhenti2 = true
				} else if strings.ToLower(Choice) == "n" {
					atribut.ClearScreen()
					berhenti = false
					berhenti2 = true
					MainMenu()
					return false
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
				}
			}
		} else {
			berhenti2 := false
			for !berhenti2 {
				fmt.Print("Materi selesai! Mau lanjut ke soal (s), atau kembali ke menu utama (n)? ")
				fmt.Scan(&Choice)
				fmt.Scanln()
				if strings.ToLower(Choice) == "n" {
					berhenti = false
					berhenti2 = true
					MainMenu()
					atribut.ClearScreen()
					return false
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 's' atau 'n'.")
				}
			}
		}
	}
	return true
}
