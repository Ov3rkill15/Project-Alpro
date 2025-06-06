package p_procedure

import (
	soal "Project-Alpro/alpro/praktikum/Prosedur" // Asumsi ini adalah package soal untuk Prosedur
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Procedure.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
=====================================
Selamat Datang di Pembelajaran Procedure
=====================================
1. Apa itu Procedure
2. Referensi Soal Procedure
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()
		atribut.ClearScreen() // Bersihkan layar setelah pilihan diinput

		switch pilihan {
		case "1":
			// belajarProcedure akan mengelola navigasi internalnya sendiri
			// dan akan kembali ke MainMenu saat selesai atau pengguna memilih 'n' dari dalamnya.
			shouldStayInMenu := belajarProcedure()
			if !shouldStayInMenu {
				// Jika belajarProcedure mengindikasikan keluar dari program, hentikan loop MainMenu.
				// Dalam implementasi belajarProcedure yang kita buat, ini berarti program akan keluar.
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// Asumsi: soal.Soal1prosedur() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program sepenuhnya.
			shouldStayInMenu := soal.Soal1prosedur() // Panggil fungsi soal dari package 'soal'
			if !shouldStayInMenu {
				// Jika soal.Soal1prosedur() mengindikasikan keluar dari program, hentikan loop MainMenu.
				fmt.Println("Terima kasih telah menggunakan program pembelajaran procedure.")
				isRunning = false
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0":
			fmt.Println("Terima kasih telah menggunakan program pembelajaran procedure.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarProcedure mengelola materi pembelajaran prosedur.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program (misalnya dari halaman terakhir).
func belajarProcedure() bool {
	halamanSekarang := 1
	totalHalaman := 3 // Ada 3 halaman materi
	var Choice string

	// isStudying akan menjadi false jika pengguna memilih 'n' untuk kembali ke menu utama.
	isStudying := true

	for isStudying { // Loop tak terbatas sampai pengguna memutuskan untuk keluar dari materi ini
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Procedure di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Dalam pemrograman, terkadang kita hanya ingin melakukan serangkaian tindakan tanpa mengharapkan hasil atau nilai kembali. Bayangkan kamu memberikan instruksi kepada asistenmu untuk melakukan tugas, seperti 'bersihkan meja' atau 'nyalakan lampu'. Kamu tidak mengharapkan asistenmu memberikan laporan atau nilai kembali; kamu hanya ingin tugas itu diselesaikan.")
			fmt.Println("\nDalam Go, fungsi yang dirancang untuk melakukan tugas tanpa mengembalikan nilai disebut **prosedur**. Secara teknis, ini masih sebuah **fungsi**, tapi kita menggunakannya hanya untuk efek samping (misalnya, mencetak sesuatu ke konsol, mengubah status, atau memodifikasi data yang diteruskan melalui pointer), bukan untuk menghasilkan nilai yang bisa kita simpan atau gunakan di bagian lain program.")
			fmt.Println("\n**Kapan Kita Menggunakan Prosedur (Fungsi Tanpa Return)?**")
			fmt.Println("- **Menampilkan Informasi:** Seperti mencetak pesan ke layar, log, atau menulis ke file.")
			fmt.Println("- **Modifikasi State:** Mengubah nilai variabel global, memodifikasi elemen dalam slice atau map (jika diteruskan secara referensi), atau berinteraksi dengan hardware.")
			fmt.Println("- **Memicu Aksi:** Mengirim notifikasi, menyimpan data ke database, atau memulai proses lain.")
			fmt.Println("- **Inisialisasi:** Menyiapkan konfigurasi awal atau sumber daya.")
		case 2:
			fmt.Println("Struktur dasar sebuah prosedur (fungsi tanpa return) di Go sangat mirip dengan fungsi biasa, hanya saja bagian return valuenya dihilangkan.")
			fmt.Println("\n```go")
			fmt.Println("func namaProsedur(parameter1 tipeData, parameter2 tipeData) {")
			fmt.Println("    // Ini adalah badan prosedur")
			fmt.Println("    // Semua kode yang melakukan tugas ada di sini")
			fmt.Println("    // ...")
			fmt.Println("    // Tidak ada pernyataan 'return' jika tidak ada nilai yang dikembalikan")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nInilah arti setiap bagian:")
			fmt.Println("- **`func` keyword:** Sama seperti fungsi biasa, ini cara kamu mendeklarasikan prosedur di Go.")
			fmt.Println("- **`namaProsedur`:** Nama unik yang kamu berikan pada prosedurmu. Pilih nama yang deskriptif untuk tugas yang dilakukannya!")
			fmt.Println("- **`()` (Tanda Kurung):** Ini untuk **parameter** (input). Kamu bisa memberikan 'bahan' agar prosedur bisa bekerja. Jika tidak butuh input, tanda kurungnya tetap kosong.")
			fmt.Println("- **Tidak ada Return Values:** Ini yang membedakannya. Setelah tanda kurung parameter, tidak ada deklarasi tipe data untuk nilai yang dikembalikan. Ini berarti prosedur ini tidak akan mengembalikan nilai apa pun.")
			fmt.Println("- **`{}` (Kurung Kurawal):** Ini adalah **badan prosedur**, tempat kamu menulis kode sebenarnya yang melakukan tugasnya.")
			fmt.Println("\n**Contoh Prosedur Sederhana:**")
			fmt.Println("```go")
			fmt.Println("func tampilkanPesanSelamatDatang(nama string) { // tidak ada tipe return")
			fmt.Println("    fmt.Printf(\"Halo, %s! Selamat datang di aplikasi kami.\\n\", nama)")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggil (menggunakan) prosedur ini:")
			fmt.Println("// tampilkanPesanSelamatDatang(\"Budi\") // Ini akan mencetak \"Halo, Budi! Selamat datang di aplikasi kami.\"")
			fmt.Println("// tampilkanPesanSelamatDatang(\"Siti\")  // Ini akan mencetak \"Halo, Siti! Selamat datang di aplikasi kami.\"")
			fmt.Println("```")
		case 3:
			fmt.Println("Prosedur (fungsi tanpa return) adalah elemen penting dalam struktur program Go, memungkinkan kita untuk mengisolasi dan mengatur kode yang melakukan aksi atau efek samping.")
			fmt.Println("\nMereka sering digunakan untuk tugas-tugas seperti:")
			fmt.Println("- **Pencetakan dan Logging:** Menampilkan informasi debugging atau pesan pengguna.")
			fmt.Println("- **Manipulasi Data In-Place:** Jika kamu memiliki pointer ke suatu struktur data, prosedur bisa memodifikasi data yang ditunjuk secara langsung.")
			fmt.Println("- **Interaksi Eksternal:** Mengirim data ke jaringan, menulis ke database, atau memodifikasi file.")
			fmt.Println("- **Pembersihan Sumber Daya:** Menggunakan `defer` dalam prosedur untuk memastikan penutupan file atau koneksi setelah operasi selesai.")
			fmt.Println("\nMemahami perbedaan antara fungsi yang mengembalikan nilai dan prosedur yang tidak mengembalikan nilai sangatlah penting. Keduanya adalah 'fungsi', tetapi tujuannya berbeda dalam desain programmu.")
			fmt.Println("\n---")

			fmt.Println("### Contoh Penggunaan Prosedur Sederhana")
			fmt.Println("\nMari kita lihat contoh prosedur sederhana yang menampilkan menu aplikasi:")
			fmt.Println("```go")
			fmt.Println("package main")
			fmt.Println("\nimport \"fmt\"")
			fmt.Println("\n// tampilkanMenu adalah prosedur yang mencetak opsi menu ke konsol.")
			fmt.Println("// Prosedur ini tidak menerima parameter dan tidak mengembalikan nilai.")
			fmt.Println("func tampilkanMenu() {")
			fmt.Println("    fmt.Println(\"------------------------\")")
			fmt.Println("    fmt.Println(\"|      MENU APLIKASI   |\")")
			fmt.Println("    fmt.Println(\"------------------------\")")
			fmt.Println("    fmt.Println(\"| 1. Tambah Data       |\")")
			fmt.Println("    fmt.Println(\"| 2. Lihat Data        |\")")
			fmt.Println("    fmt.Println(\"| 3. Hapus Data        |\")")
			fmt.Println("    fmt.Println(\"| 4. Keluar            |\")")
			fmt.Println("    fmt.Println(\"------------------------\")")
			fmt.Println("}")
			fmt.Println("\nfunc main() {")
			fmt.Println("    // Memanggil prosedur tampilkanMenu")
			fmt.Println("    tampilkanMenu()")
			fmt.Println("\n    fmt.Println(\"Silakan pilih opsi Anda...\")")
			fmt.Println("}")
			fmt.Println("```")
			fmt.Println("\nDalam contoh di atas:")
			fmt.Println("* Kita mendefinisikan prosedur bernama `tampilkanMenu` yang tidak menerima parameter dan tidak memiliki nilai kembalian.")
			fmt.Println("* Tugas utamanya adalah mencetak serangkaian baris ke konsol untuk membentuk menu.")
			fmt.Println("* Di dalam fungsi `main`, kita **memanggil** `tampilkanMenu()` untuk mengeksekusi kode di dalamnya dan menampilkan menu kepada pengguna.")
			fmt.Println("\nIni adalah contoh dasar bagaimana prosedur membantu kita mengorganisir kode untuk melakukan tindakan tertentu tanpa perlu mengembalikan nilai.")
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
