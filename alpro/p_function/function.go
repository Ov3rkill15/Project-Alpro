package p_function

import (
	soal "Project-Alpro/alpro/praktikum/Fungsi" // Asumsi ini adalah package soal untuk fungsi
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Function.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	// isRunning akan menjadi false jika pengguna memilih untuk keluar (0).
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
=====================================
Selamat Datang di Pembelajaran Function
=====================================
1. Apa itu Function
2. Referensi Soal Fungsi
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarFunction akan mengelola navigasi internalnya sendiri
			// dan akan kembali ke MainMenu saat selesai atau pengguna memilih 'n' dari dalamnya.
			shouldStayInMenu := belajarFunction()
			if !shouldStayInMenu {
				// Jika belajarFunction mengindikasikan keluar dari program, hentikan loop MainMenu.
				// Dalam implementasi belajarFunction yang kita buat, ini berarti program akan keluar.
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// Asumsi: soal.Soal1fungsi() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal1fungsi() // Panggil fungsi soal dari package 'soal'
			if !shouldStayInMenu {
				// Jika soal.Soal1fungsi() mengindikasikan keluar dari program, hentikan loop MainMenu.
				fmt.Println("Terima kasih telah menggunakan program pembelajaran function.")
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran function.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarFunction mengelola materi pembelajaran fungsi.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program (misalnya dari halaman terakhir).
func belajarFunction() bool {
	halamanSekarang := 1
	totalHalaman := 3 // Ada 3 halaman materi
	var Choice string

	// isStudying akan menjadi false jika pengguna memilih 'n' untuk kembali ke menu utama.
	isStudying := true

	for isStudying { // Loop tak terbatas sampai pengguna memutuskan untuk keluar dari materi ini
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
			fmt.Println("    fmt.Printf(\"Halo, %s! Selamat datang di program kami.\\n\", nama)")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggil (menggunakan) fungsi ini:")
			fmt.Println("// sapaPengguna(\"Budi\") // Ini akan mencetak \"Halo, Budi! Selamat datang di program kami.\"")
			fmt.Println("// sapaPengguna(\"Ani\")  // Ini akan mencetak \"Halo, Ani! Selamat datang di program kami.\"")
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
			fmt.Println("    fmt.Printf(\"Luas persegi panjang dengan panjang 10 dan lebar 5 adalah: %d\\n\", luas1)")
			fmt.Println("\n    luas2 := hitungLuasPersegiPanjang(7, 3)")
			fmt.Println("    fmt.Printf(\"Luas persegi panjang dengan panjang 7 dan lebar 3 adalah: %d\\n\", luas2)")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDalam contoh di atas:")
			fmt.Println("* Kita mendefinisikan fungsi bernama `hitungLuasPersegiPanjang` yang menerima dua *integer* sebagai **input** (`panjang` dan `lebar`).")
			fmt.Println("* Fungsi ini **mengembalikan** satu *integer* yang merupakan hasil perkalian `panjang` dan `lebar`.")
			fmt.Println("* Di dalam fungsi `main`, kita **memanggil** `hitungLuasPersegiPanjang` beberapa kali dengan nilai yang berbeda, menunjukkan bagaimana fungsi dapat digunakan kembali.")
			fmt.Println("\nIni adalah contoh dasar bagaimana fungsi membantu kita mengorganisir kode dan menghindari pengulangan.")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			// Jika halaman tidak valid, kita akan menganggap pengguna ingin kembali ke menu utama.
			return true // Kembali ke MainMenu
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya (outer loop)
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
			}
		}
	}
	// Setelah loop isStudying berakhir (karena Choice adalah 'n'), kita mengembalikan true
	// yang menandakan bahwa MainMenu harus terus berjalan.
	return true
}
