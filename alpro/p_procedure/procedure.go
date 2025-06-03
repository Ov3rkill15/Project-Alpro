package p_procedure

import (
	soal "Project-Alpro/alpro/praktikum/Prosedur"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

func MainMenu() {
	var pilihan string
	var pilihan2 string

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
	atribut.ClearScreen()

	switch pilihan {
	case "1":
		atribut.ClearScreen()
		if belajarProcedure() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			berhenti2 := false
			for !berhenti2 {
				fmt.Scan(&pilihan2)
				fmt.Scanln()
				if strings.ToLower(pilihan2) == "y" {
					berhenti2 = true
					MainMenu()
					return
				} else if strings.ToLower(pilihan2) == "n" {
					berhenti2 = true
					return
				} else {
					fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
					fmt.Print("Mau pilih materi lain atau kembali ke menu utama?(y/n): ")
				}
			}
		} else {
			return
		}
	case "2":
		atribut.ClearScreen()
		if soal.Soal1prosedur() {
			fmt.Println("Mau pilih materi lain atau kembali ke menu utama?(y/n)")
			berhenti2 := false
			for !berhenti2 {
				fmt.Scan(&pilihan2)
				fmt.Scanln()
				if strings.ToLower(pilihan2) == "y" {
					berhenti2 = true
					MainMenu()
					return
				} else if strings.ToLower(pilihan2) == "n" {
					berhenti2 = true
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
		fmt.Println("Terima kasih telah menggunakan program pembelajaran procedure.")
		return
	default:
		fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', atau '0'.")
		fmt.Print("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
		MainMenu()
	}
}

func belajarProcedure() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var berhenti bool = true
	var Choice string

	for berhenti {
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
			fmt.Println("    fmt.Printf(\"Halo, %s! Selamat datang di aplikasi kami.\\n\", nama)") // Perbaikan: %s untuk nama
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
			fmt.Println("    fmt.Println(\" | 2. Lihat Data        |\")")
			fmt.Println("    fmt.Println(\" | 3. Hapus Data        |\")")
			fmt.Println("    fmt.Println(\" | 4. Keluar            |\")")
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
				fmt.Print("kembali ke menu utama (n) ")
				fmt.Scan(&Choice)
				fmt.Scanln()
				if strings.ToLower(Choice) == "n" {
					berhenti = false
					berhenti2 = true
					MainMenu()
					atribut.ClearScreen()
					return false
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		}
	}
	return true
}
