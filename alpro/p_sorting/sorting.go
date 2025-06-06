package p_sorting

import (
	selection "Project-Alpro/alpro/praktikum/Selection_Sort"
	insertion "Project-Alpro/alpro/praktikum/insertion_sort"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

// MainMenu adalah fungsi utama untuk navigasi menu pembelajaran Sorting.
func MainMenu() {
	var pilihan string
	// Gunakan variabel boolean untuk mengontrol loop menu utama.
	isRunning := true

	for isRunning { // Loop utama agar menu terus tampil sampai pengguna memilih untuk keluar
		atribut.ClearScreen() // Bersihkan layar setiap kali menu utama ditampilkan
		fmt.Println(`
======================================
Selamat Datang di Pembelajaran Sorting
======================================
1. Apa itu Selection Sort
2. Apa itu Insertion Sort
3. Referensi Soal Selection Sort
4. Referensi Soal Insertion Sort
0. Keluar
		`)
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln() // Membersihkan newline
		atribut.ClearScreen()

		switch pilihan {
		case "1":
			// belajarSelectionSort akan mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := belajarSelectionSort()
			if !shouldStayInMenu {
				isRunning = false // Hentikan loop MainMenu jika belajarSelectionSort mengindikasikan keluar
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "2":
			// belajarInsertionSort akan mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := belajarInsertionSort()
			if !shouldStayInMenu {
				isRunning = false // Hentikan loop MainMenu jika belajarInsertionSort mengindikasikan keluar
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "3": // Case 3: Referensi Soal Selection Sort
			// Asumsi: selection.SoalSelectionSort() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := selection.SoalSelectionSort()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "4": // Case 4: Referensi Soal Insertion Sort
			// Asumsi: insertion.SoalInsertionSort() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := insertion.SoalInsertionSort()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}
			// Jika shouldStayInMenu true, loop MainMenu akan berlanjut secara otomatis.

		case "0": // Case 0: Keluar
			fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', '3', '4', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan...")
			fmt.Scanln() // Tunggu Enter
			// Loop akan berlanjut secara otomatis karena isRunning masih true
		}
	}
}

// belajarSelectionSort mengelola tampilan materi Selection Sort.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program.
func belajarSelectionSort() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var Choice string
	isStudying := true // Kontrol loop utama materi

	for isStudying {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Selection Sort di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan Anda memiliki tumpukan kartu yang tidak berurutan dan ingin mengurutkannya dari yang terkecil ke terbesar. Dengan **Selection Sort**, Anda akan mencari kartu terkecil dari seluruh tumpukan, lalu meletakkannya di posisi paling awal. Kemudian, Anda ulangi proses ini untuk sisa kartu yang belum terurut.")
			fmt.Println("\nDalam pemrograman, **Selection Sort (Pengurutan Seleksi)**, atau sering disebut **Linear Sort**, adalah algoritma pengurutan sederhana yang bekerja dengan cara berulang kali menemukan elemen minimum (atau maksimum) dari bagian daftar yang belum terurut dan menukarnya (swap) ke posisi yang benar di bagian yang terurut.")
			fmt.Println("\n**Konsep Dasar Selection Sort:**")
			fmt.Println("1.  Pindai seluruh daftar untuk menemukan elemen terkecil.")
			fmt.Println("2.  Tukar elemen terkecil tersebut dengan elemen di posisi pertama daftar.")
			fmt.Println("3.  Kemudian, pindai sisa daftar (mulai dari posisi kedua) untuk menemukan elemen terkecil berikutnya.")
			fmt.Println("4.  Tukar elemen terkecil yang ditemukan dengan elemen di posisi kedua.")
			fmt.Println("5.  Lanjutkan proses ini hingga seluruh daftar terurut.")
			fmt.Println("\n---")
			fmt.Println("### Pseudocode Selection Sort:")
			fmt.Println("```")
			fmt.Println("kamus")
			fmt.Println("  i : integer { untuk indeks pencacah pencarian nilai ekstrim }")
			fmt.Println("  idx : integer { untuk indeks nilai ekstrim }")
			fmt.Println("  pass : integer { 1 siklus proses pencarian idx nilai ekstrim dan")
			fmt.Println("             pertukaran }")
			fmt.Println("  temp : integer { untuk menyimpan data sementara, patokan elemen array }")
			fmt.Println("algoritma")
			fmt.Println("{ outer loop }")
			fmt.Println("pass <- 1")
			fmt.Println("while pass <= n-1 do")
			fmt.Println("  { inner loop }")
			fmt.Println("  { 1. Pencarian nilai idx ekstrim (minimum) }")
			fmt.Println("  idx <- pass - 1")
			fmt.Println("  i <- pass")
			fmt.Println("  while i <= n-1 do")
			fmt.Println("    if A[i] < A[idx] then")
			fmt.Println("      idx <- i")
			fmt.Println("    endif")
			fmt.Println("    i <- i + 1")
			fmt.Println("  endwhile")
			fmt.Println("  { 2. Pertukaran}")
			fmt.Println("  temp <- A[idx]")
			fmt.Println("  A[idx] <- A[pass-1]")
			fmt.Println("  A[pass-1] <- temp")
			fmt.Println("  pass <- pass + 1")
			fmt.Println("endwhile")
			fmt.Println("endprocedure")
			fmt.Println("```")
		case 2:
			fmt.Println("Berikut adalah ilustrasi langkah demi langkah Selection Sort:")
			fmt.Println("Misalkan kita punya array: [64, 25, 12, 22, 11]")
			fmt.Println("\n**Iterasi 1:**")
			fmt.Println("- Cari elemen terkecil di [64, 25, 12, 22, 11]. Yaitu `11`.")
			fmt.Println("- Tukar `11` dengan `64` (elemen di posisi pertama).")
			fmt.Println("  Array menjadi: [11, 25, 12, 22, 64]")
			fmt.Println("  Bagian terurut: [11] | Bagian belum terurut: [25, 12, 22, 64]")
			fmt.Println("\n**Iterasi 2:**")
			fmt.Println("- Cari elemen terkecil di bagian belum terurut [25, 12, 22, 64]. Yaitu `12`.")
			fmt.Println("- Tukar `12` dengan `25` (elemen di posisi kedua).")
			fmt.Println("  Array menjadi: [11, 12, 25, 22, 64]")
			fmt.Println("  Bagian terurut: [11, 12] | Bagian belum terurut: [25, 22, 64]")
			fmt.Println("\nDan seterusnya, sampai seluruh array terurut.")
			fmt.Println("\n**Contoh Implementasi Selection Sort di Go:**")
			fmt.Println("```go")
			fmt.Println("func selectionSort(arr []int) {")
			fmt.Println("    n := len(arr)")
			fmt.Println("    for i := 0; i < n-1; i++ {")
			fmt.Println("        // Cari elemen minimum di sisa array yang belum terurut")
			fmt.Println("        minIndex := i")
			fmt.Println("        for j := i + 1; j < n; j++ {")
			fmt.Println("            if arr[j] < arr[minIndex] {")
			fmt.Println("                minIndex = j")
			fmt.Println("            }")
			fmt.Println("        }")
			fmt.Println("        // Tukar elemen minimum yang ditemukan dengan elemen pertama dari bagian belum terurut")
			fmt.Println("        arr[i], arr[minIndex] = arr[minIndex], arr[i]")
			fmt.Println("    }")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggilnya:")
			fmt.Println("// numbers := []int{64, 25, 12, 22, 11}")
			fmt.Println("// selectionSort(numbers)")
			fmt.Println("// fmt.Println(numbers) // Output: [11 12 22 25 64]")
			fmt.Println("```")
		case 3:
			fmt.Println("### Kelebihan dan Kekurangan Selection Sort")
			fmt.Println("\n**Kelebihan:**")
			fmt.Println("- **Jumlah Pertukaran Minimal:** Selection Sort melakukan jumlah pertukaran (swap) elemen yang paling sedikit dibandingkan algoritma pengurutan perbandingan lainnya. Ini bisa menjadi keuntungan jika operasi pertukaran sangat 'mahal' (memakan banyak waktu atau sumber daya).")
			fmt.Println("- **Sederhana:** Mudah dipahami dan diimplementasikan.")
			fmt.Println("\n**Kekurangan:**")
			fmt.Println("- **Tidak Efisien untuk Data Besar:** Baik dalam kasus terbaik, rata-rata, maupun terburuk, **Time Complexity** Selection Sort adalah **O(n^2)**. Ini berarti waktu eksekusi meningkat secara kuadratis dengan ukuran input. Untuk daftar yang besar, ini akan sangat lambat.")
			fmt.Println("- **Tidak Adaptif:** Kinerjanya tidak meningkat bahkan jika daftar sudah sebagian terurut. Ia akan selalu melakukan jumlah perbandingan yang sama.")
			fmt.Println("\nKarena efisiensinya yang O(n^2), Selection Sort jarang digunakan pada daftar besar. Namun, kesederhanaannya membuatnya cocok untuk daftar yang sangat kecil atau sebagai alat pembelajaran fundamental dalam algoritma pengurutan.")
			fmt.Println("\n---")
			fmt.Println("### Latihan Sederhana Selection Sort")
			fmt.Println("Coba terapkan Selection Sort untuk mengurutkan daftar string (nama-nama).")
			fmt.Println("```go")
			fmt.Println("func sortStrings(names []string) {")
			fmt.Println("    // Tulis kodemu di sini")
			fmt.Println("    // Hint: Perbandingan string di Go menggunakan operator < atau >")
			fmt.Println("}")
			fmt.Println("```")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			return true // Kembali ke MainMenu jika ada kesalahan halaman
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln() // Clear newline

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
				return true        // Kembali ke MainMenu
			} else if strings.ToLower(Choice) == "s" && halamanSekarang == totalHalaman {
				// Handle redirection to Selection Sort quiz directly from here if 's' is chosen
				atribut.ClearScreen()
				shouldStayInMenu := selection.SoalSelectionSort() // Panggil soal selection sort
				if !shouldStayInMenu {
					return false // Exit program
				}
				isStudying = false // Stop studying loop to prevent re-displaying this section
				inputValid = true  // Exit input loop
				return true        // Return to MainMenu after quiz
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y', 'n', atau 's' (jika di halaman terakhir).")
			}
		}
	}
	return true // Default return: kembali ke MainMenu setelah selesai belajar
}

// belajarInsertionSort mengelola tampilan materi Insertion Sort.
// Mengembalikan true jika pengguna ingin kembali ke MainMenu,
// atau false jika pengguna ingin keluar dari program.
func belajarInsertionSort() bool {
	halamanSekarang := 1
	totalHalaman := 3
	var Choice string
	isStudying := true // Kontrol loop utama materi

	for isStudying {
		atribut.ClearScreen()
		fmt.Printf("=== Apa itu Insertion Sort di Go? (Halaman %d/%d) ===\n", halamanSekarang, totalHalaman)

		switch halamanSekarang {
		case 1:
			fmt.Println("Bayangkan Anda sedang mengurutkan kartu di tangan Anda. Anda mengambil satu per satu kartu dari tumpukan yang belum diurutkan, lalu menyisipkannya ke posisi yang benar di antara kartu-kartu yang sudah terurut di tangan Anda. Proses ini terus berlanjut sampai semua kartu tersisipkan dengan benar.")
			fmt.Println("\nDalam pemrograman, **Insertion Sort (Pengurutan Sisipan)** adalah algoritma pengurutan sederhana yang bekerja dengan membangun array (atau daftar) terurut satu elemen pada satu waktu. Ini mengulangi seluruh elemen input dan memindahkannya ke tempat yang sesuai dalam array terurut.")
			fmt.Println("\n**Konsep Dasar Insertion Sort:**")
			fmt.Println("1.  Anggap elemen pertama sudah terurut (sub-array terurut dimulai dari elemen pertama).")
			fmt.Println("2.  Ambil elemen berikutnya dari sub-array yang belum terurut.")
			fmt.Println("3.  Bandingkan elemen ini dengan elemen-elemen di sub-array terurut, bergerak mundur dari kanan ke kiri.")
			fmt.Println("4.  Geser elemen-elemen yang lebih besar dari elemen yang sedang dipertimbangkan satu posisi ke kanan.")
			fmt.Println("5.  Sisipkan elemen yang diambil ke posisi yang benar.")
			fmt.Println("6.  Ulangi proses ini sampai semua elemen tersisipkan.")
			fmt.Println("\n---")
			fmt.Println("### Pseudocode Insertion Sort:")
			fmt.Println("```")
			fmt.Println("procedure insertionSort(in/out A : tabInt, in n : integer)")
			fmt.Println("{ IS: Array A sebanyak n elemen terdefinisi terurut acak")
			fmt.Println("  FS: Array A sebanyak n elemen terurut menaik (ascending)")
			fmt.Println("      menggunakan algoritma insertion sort }")
			fmt.Println("")
			fmt.Println("kamus")
			fmt.Println("  i, pass : integer")
			fmt.Println("  temp : integer")
			fmt.Println("algoritma")
			fmt.Println("pass <- 1")
			fmt.Println("while pass <= n - 1 do")
			fmt.Println("  { Pencarian indeks yang tepat untuk elemen }")
			fmt.Println("  i <- pass")
			fmt.Println("  temp <- A[pass]")
			fmt.Println("  while i > 0 and A[i-1] > temp do")
			fmt.Println("    A[i] <- A[i-1]")
			fmt.Println("    i <- i - 1")
			fmt.Println("  endwhile")
			fmt.Println("  { Menempatkan elemen pada lokasi tersebut}")
			fmt.Println("  A[i] <- temp")
			fmt.Println("  pass <- pass + 1")
			fmt.Println("endwhile")
			fmt.Println("endprocedure")
			fmt.Println("```")
		case 2:
			fmt.Println("Berikut adalah ilustrasi langkah demi langkah Insertion Sort:")
			fmt.Println("Misalkan kita punya array: [12, 11, 13, 5, 6]")
			fmt.Println("\n**Awal:** Array terurut dianggap [12] | Belum terurut: [11, 13, 5, 6]")
			fmt.Println("\n**Iterasi 1 (Ambil 11):**")
			fmt.Println("- Bandingkan 11 dengan 12. Karena 11 < 12, geser 12 ke kanan.")
			fmt.Println("- Sisipkan 11 di awal.")
			fmt.Println("  Array menjadi: [11, 12, 13, 5, 6]")
			fmt.Println("  Bagian terurut: [11, 12] | Bagian belum terurut: [13, 5, 6]")
			fmt.Println("\n**Iterasi 2 (Ambil 13):**")
			fmt.Println("- Bandingkan 13 dengan 12. Karena 13 > 12, tidak ada yang digeser. Sisipkan 13.")
			fmt.Println("  Array menjadi: [11, 12, 13, 5, 6]")
			fmt.Println("  Bagian terurut: [11, 12, 13] | Bagian belum terurut: [5, 6]")
			fmt.Println("\nDan seterusnya, sampai seluruh array terurut.")
			fmt.Println("\n**Contoh Implementasi Insertion Sort di Go:**")
			fmt.Println("```go")
			fmt.Println("func insertionSort(arr []int) {")
			fmt.Println("    n := len(arr)")
			fmt.Println("    for i := 1; i < n; i++ {")
			fmt.Println("        key := arr[i] // Elemen yang akan disisipkan")
			fmt.Println("        j := i - 1")
			fmt.Println("\n        // Pindahkan elemen arr[0..i-1] yang lebih besar dari key")
			fmt.Println("        // ke satu posisi di depan posisi saat ini")
			fmt.Println("        for j >= 0 && arr[j] > key {")
			fmt.Println("            arr[j+1] = arr[j]")
			fmt.Println("            j = j - 1")
			fmt.Println("        }")
			fmt.Println("        arr[j+1] = key // Sisipkan key pada posisi yang benar")
			fmt.Println("    }")
			fmt.Println("}")
			fmt.Println("\n// Cara memanggilnya:")
			fmt.Println("// numbers := []int{12, 11, 13, 5, 6}")
			fmt.Println("// insertionSort(numbers)")
			fmt.Println("// fmt.Println(numbers) // Output: [5 6 11 12 13]")
			fmt.Println("```")
		case 3:
			fmt.Println("### Kelebihan dan Kekurangan Insertion Sort")
			fmt.Println("\n**Kelebihan:**")
			fmt.Println("- **Sederhana:** Mudah dipahami dan diimplementasikan.")
			fmt.Println("- **Efisien untuk Data Kecil atau Hampir Terurut:** Kinerjanya sangat baik untuk array dengan jumlah elemen kecil atau array yang sudah hampir terurut.")
			fmt.Println("- **Stabil:** Tidak mengubah urutan relatif dari elemen-elemen yang memiliki nilai yang sama.")
			fmt.Println("- **In-place:** Tidak memerlukan memori tambahan yang signifikan.")
			fmt.Println("\n**Kekurangan:**")
			fmt.Println("- **Tidak Efisien untuk Data Besar:** Sama seperti Selection Sort, **Time Complexity** Insertion Sort adalah **O(n^2)** dalam kasus terburuk dan rata-rata. Ini membuatnya tidak cocok untuk pengurutan daftar yang sangat besar.")
			fmt.Println("\nInsertion Sort sering digunakan dalam situasi di mana daftar masukan relatif kecil atau sudah terurut sebagian. Ini juga menjadi bagian dari algoritma pengurutan hibrida (seperti Timsort dan Introsort) untuk menangani bagian-bagian kecil dari daftar yang lebih besar.")
			fmt.Println("\n---")
			fmt.Println("### Latihan Sederhana Insertion Sort")
			fmt.Println("Coba terapkan Insertion Sort untuk mengurutkan daftar angka secara menurun (dari terbesar ke terkecil).")
			fmt.Println("```go")
			fmt.Println("func sortDescending(numbers []float64) {")
			fmt.Println("    // Tulis kodemu di sini")
			fmt.Println("    // Hint: Ubah kondisi perbandingan agar mengurutkan secara menurun")
			fmt.Println("}")
			fmt.Println("```")
		default:
			fmt.Println("Halaman tidak ditemukan.")
			return true // Kembali ke MainMenu jika ada kesalahan halaman
		}

		fmt.Println("\n------------------------------------")

		// Loop untuk mendapatkan input navigasi halaman
		inputValid := false
		for !inputValid {
			if halamanSekarang < totalHalaman {
				fmt.Print("Lanjut ke halaman berikutnya (y) atau kembali ke menu utama (n)? ")
			} else { // Jika sudah di halaman terakhir
				fmt.Print("Materi selesai! Mau lanjut ke Soal (s) atau kembali ke menu utama (n)? ")
			}
			fmt.Scan(&Choice)
			fmt.Scanln() // Clear newline

			if strings.ToLower(Choice) == "y" && halamanSekarang < totalHalaman {
				halamanSekarang++
				inputValid = true // Input valid, keluar dari loop input dan lanjut ke halaman berikutnya
			} else if strings.ToLower(Choice) == "n" {
				isStudying = false // Mengindikasikan untuk keluar dari loop materi
				inputValid = true  // Input valid, keluar dari loop input
				return true        // Kembali ke MainMenu
			} else if strings.ToLower(Choice) == "s" && halamanSekarang == totalHalaman {
				// Handle redirection to Insertion Sort quiz directly from here if 's' is chosen
				atribut.ClearScreen()
				shouldStayInMenu := insertion.SoalInsertionSort() // Panggil soal insertion sort
				if !shouldStayInMenu {
					return false // Exit program
				}
				isStudying = false // Stop studying loop to prevent re-displaying this section
				inputValid = true  // Exit input loop
				return true        // Return to MainMenu after quiz
			} else {
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y', 'n', atau 's' (jika di halaman terakhir).")
			}
		}
	}
	return true // Default return: kembali ke MainMenu setelah selesai belajar
}
