package p_function

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

var pilihan string
var pilihan2 string

func Submenu() {
	fmt.Println(`
==================
Selamat Datang di Pembelajaran function
==================
1. apa itu function anjay
2. soal function
3. Keluar
	`)
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)
}

func MainMenu() {
	atribut.ClearScreen()
	Submenu()
	switch pilihan {
	case "1":
		atribut.ClearScreen()
		belajarFunction()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}

	case "2":
		atribut.ClearScreen()
		soalFunction()
		fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			MainMenu()
		} else {
			return
		}
	case "3":
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
}

func belajarFunction() {
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
			fmt.Println("    // Ini adalah badan fungsi")
			fmt.Println("    // Semua kode yang melakukan tugas fungsi ada di sini")
			fmt.Println("    // ...")
			fmt.Println("    return nilai1, nilai2 // Jika fungsi mengembalikan nilai")
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
			fmt.Println("    fmt.Printf(\"Halo, %s! Selamat datang di program kami.\\n\", nama)")
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
		}

		fmt.Println("\n------------------------------------")

		if halamanSekarang < totalHalaman {
			fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "y" {
				halamanSekarang++
			} else {
				atribut.ClearScreen()
				MainMenu()
			}
		} else {
			fmt.Print("Materi selesai! Mau lanjut ke soa (k), atau kembali ke menu utama (n)? ")
			fmt.Scan(&Choice)
			fmt.Scanln()

			if strings.ToLower(Choice) == "k" {
				soalFunction()
				berhenti = false
			} else if strings.ToLower(Choice) == "n" {
				MainMenu()
				berhenti = false
				atribut.ClearScreen()
			} else {
				var i int
				for i = 0; i < 2; i++ {
					fmt.Println("Pilihan tidak valid. Silakan masukkan 'k' atau 'n'.")
					fmt.Scan(&Choice)
				}
				MainMenu()
			}

		}
	}
}

func soalFunction() {
	fmt.Println("COMING SOON!")
}
