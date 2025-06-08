package p_sorting

import (
	selection "Project-Alpro/alpro/praktikum/Selection_Sort"
	insertion "Project-Alpro/alpro/praktikum/insertion_sort"
	"Project-Alpro/atribut"
	"fmt"
	"strings"
	"time"
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
3. Perbandingan Selection Sort dan Insertion Sort
4. Referensi Soal Selection Sort
5. Referensi Soal Insertion Sort
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

		case "2":
			// belajarInsertionSort akan mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := belajarInsertionSort()
			if !shouldStayInMenu {
				isRunning = false // Hentikan loop MainMenu jika belajarInsertionSort mengindikasikan keluar
			}

		case "3":
			shouldStayInMenu := visualisasi()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}

		case "4":
			shouldStayInMenu := selection.SoalSelectionSort()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}

		case "5": // Case 4: Referensi Soal Insertion Sort
			// Asumsi: insertion.SoalInsertionSort() mengembalikan true jika ingin kembali ke MainMenu,
			// dan false jika pengguna ingin keluar dari program.
			shouldStayInMenu := insertion.SoalInsertionSort()
			if !shouldStayInMenu {
				fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
				isRunning = false // Hentikan loop MainMenu jika ingin keluar dari program
			}

		case "0": // Case 0: Keluar
			fmt.Println("Terima kasih telah menggunakan program pembelajaran sorting.")
			isRunning = false // Set isRunning menjadi false untuk menghentikan loop

		default:
			fmt.Println("Pilihan tidak valid. Harap masukkan '1', '2', '3', '4', atau '0'.")
			fmt.Print("Tekan Enter untuk melanjutkan..karena isRunning masih true")
		}
	}
}

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

const nmax int = 100 // Ukuran maksimum array
const nmax2 int = 5

type Tabint [nmax]int // Tipe array integer dengan ukuran nmax

type Tabuser [nmax2]int

var data10 Tabint
var data50 Tabint
var data100 Tabint

func initData10(jData10 *int) {
	data := [10]int{73, 74, 8, 23, 59, 9, 24, 83, 61, 91}
	*jData10 = 0
	i := 0
	for i < nmax {
		if i < len(data) {
			data10[i] = data[i]
			*jData10++
		}
		i++
	}
}

func initData50(jData50 *int) {
	data := [50]int{36, 8, 78, 82, 76, 96, 23, 2, 91, 86, 36, 40, 93, 4, 72, 13, 47, 64, 9, 100, 81, 89, 8, 39, 48, 21, 38, 43, 5, 16, 26, 12, 83, 67, 65, 94, 39, 2, 9, 37, 49, 93, 96, 88, 77, 65, 47, 19, 25, 1}
	*jData50 = 0
	i := 0
	for i < nmax {
		if i < len(data) {
			data50[i] = data[i]
			*jData50++
		}
		i++
	}
}

func initData100(jData100 *int) {
	data := [100]int{31, 81, 60, 60, 22, 19, 24, 11, 86, 77, 65, 91, 22, 49, 64, 7, 22, 5, 31, 22, 3, 21, 89, 32, 21, 8, 4, 98, 41, 25, 70, 65, 42, 97, 7, 56, 15, 76, 63, 49, 81, 32, 79, 66, 62, 70, 13, 2, 74, 28, 40, 13, 79, 22, 83, 79, 7, 100, 3, 2, 93, 26, 61, 75, 75, 86, 100, 88, 36, 55, 40, 60, 54, 40, 77, 75, 45, 83, 92, 100, 53, 95, 90, 99, 72, 67, 39, 42, 34, 63, 62, 28, 84, 30, 71, 67, 8, 22, 67, 17}
	*jData100 = 0
	i := 0
	for i < nmax {
		if i < len(data) {
			data100[i] = data[i]
			*jData100++
		}
		i++
	}
}

func visualisasi() bool {
	var pilihan string
	var efisiensi int // Variabel untuk menyimpan jumlah efisiensi. Dideklarasikan sekali saja.
	var const10 int
	var const50 int
	var const100 int
	//===================== 10 DATA ===========================
	initData10(&const10)
	fmt.Printf("================== %d DATA ==================\n", const10)
	var data10sel Tabint = data10
	var data10ins Tabint = data10 // Data terpisah untuk Insertion Sort

	fmt.Println("Data asli (10 data):")
	cetakData(data10sel, const10)

	// --- Selection Sort ---
	SelectionSort(&data10sel, const10, &efisiensi)
	fmt.Println("Terurut descending oleh Selection Sort:")
	cetakData(data10sel, const10)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	efisiensi = 0
	// --- Insertion Sort ---
	insertionSort(&data10ins, const10, &efisiensi) // Gunakan data10ins
	fmt.Println("Terurut descending oleh Insertion Sort:")
	cetakData(data10ins, const10)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	efisiensi = 0

	//===================== 50 DATA ===========================
	initData50(&const50)
	fmt.Printf("================== %d DATA ==================\n", const50)
	var data50sel Tabint = data50
	var data50ins Tabint = data50
	// --- Selection Sort ---
	SelectionSort(&data50sel, const50, &efisiensi)
	fmt.Printf("Terurut descending oleh Selection Sort(%d data tanpa di print):\n", const50)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	efisiensi = 0
	// --- Insertion Sort ---
	insertionSort(&data50ins, const50, &efisiensi)
	fmt.Printf("Terurut descending oleh Insertion Sort(%d data tanpa di print):\n", const50)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	efisiensi = 0

	//===================== 100 DATA ===========================
	initData100(&const100)
	fmt.Printf("================== %d DATA ==================\n", const100)
	var data100sel Tabint = data100
	var data100ins Tabint = data100 // Ukuran data untuk set 100

	// --- Selection Sort ---
	SelectionSort(&data100sel, const100, &efisiensi)
	fmt.Printf("Terurut descending oleh Selection Sort(%d data tanpa di print):\n", const100)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	efisiensi = 0
	// --- Insertion Sort ---
	insertionSort(&data100ins, const100, &efisiensi)
	fmt.Printf("Terurut descending oleh Insertion Sort(%d data tanpa di print):\n", const100)
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)

	fmt.Println("\nProgram Selesai.")
	efisiensi = 0
	if nmax < 100 {
		fmt.Println("============================================")
		fmt.Println("Hanya sebagian data yang masuk (NMAX < 100).")
		fmt.Println("============================================")
	}

	//===================== MASUKAN DATA ===========================
	fmt.Println("\n================== MASUKAN DATA ==================")
	var dataMasuk Tabint
	var nDataMasuk int
	var dataTemp Tabint
	bacadata(&dataMasuk, &nDataMasuk)
	fmt.Println("Data masuk:")
	cetakData(dataMasuk, nDataMasuk)
	dataTemp = dataMasuk
	// --- Selection Sort ---
	PenjelasanSelectionSort(&dataTemp, nDataMasuk, &efisiensi)
	fmt.Println("Terurut descending oleh Selection Sort:")
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	fmt.Println("=============================================")
	cetakData(dataMasuk, nDataMasuk)
	efisiensi = 0
	// --- Insertion Sort ---
	PenjelasanInsertionSort(&dataMasuk, nDataMasuk, &efisiensi)
	fmt.Println("Terurut descending oleh Insertion Sort:")
	fmt.Println("Dengan proses efisiensi sebanyak:", efisiensi)
	fmt.Println("=============================================")
	inputValid := false
	for !inputValid {
		fmt.Println("\n------------------------------------")
		fmt.Print("Ingin kembali ke menu utama (n)? ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		if strings.ToLower(pilihan) == "n" {
			inputValid = true // Input valid, keluar dari loop input
			// Mengembalikan true berarti kembali ke MainMenu
			return true // Kembali ke MainMenu
		} else {
			fmt.Println("Pilihan tidak valid. Harap masukkan 'n'.")
			// Loop akan terus berjalan sampai input 'n' diterima
		}
	}
	// Jika loop berakhir karena input valid, berarti 'n' sudah dimasukkan,
	return true
}

// reset mengatur ulang elemen array menjadi 0 dan mereset jumlah data aktif.

func bacadata(A *Tabint, n *int) {
	var i int = 0
	var x int
	var stop bool = true

	fmt.Printf("MAKS DATA: %d\n", nmax2)
	fmt.Println("Masukkan data (0 untuk berhenti):")
	for i < nmax2 && stop {
		fmt.Scan(&x)
		if x == 0 {
			stop = false
		} else {
			A[i] = x
			i++
		}
	}
	*n = i
}

// cetakData mencetak elemen array dari indeks 0 hingga n-1.
func cetakData(A Tabint, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
		if n == 0 {
			fmt.Println("[Array kosong]")
		}
	}
	fmt.Println()
}

func SelectionSort(A *Tabint, n int, eff *int) {
	var i, idx, pass int
	var temp int
	pass = 1
	for pass < n {
		i = pass
		idx = pass - 1
		for i < n {
			*eff++
			if A[i] > A[idx] {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}
func insertionSort(A *Tabint, n int, eff *int) {
	var i, pass int
	var temp int
	pass = 1

	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp > A[i-1] {
			*eff++
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++

	}
}

func PenjelasanSelectionSort(A *Tabint, n int, eff *int) {
	var i, idx, pass int
	var temp int
	*eff = 0
	pass = 1
	for pass < n {
		i = pass
		idx = pass - 1
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Memulai Pass ke-", pass, "---")
		fmt.Printf("Mencari elemen terbesar untuk ditempatkan di posisi A[%d] (yang saat ini bernilai %d).\n", pass-1, A[pass-1])
		idx = pass - 1
		for i < n {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("   Membandingkan A[%d]: %d dengan elemen terbesar sementara (A[%d]): %d\n", i, A[i], idx, A[idx])
			*eff++
			if A[i] > A[idx] {
				fmt.Printf("   -> %d LEBIH BESAR dari %d. Index terbesar sementara berubah dari %d ke %d.\n", A[i], A[idx], idx, i)
				idx = i
			} else {
				fmt.Printf("   -> %d TIDAK lebih besar dari %d. Index terbesar sementara tetap %d.\n", A[i], A[idx], idx)
			}
			i++
		}
		time.Sleep(1 * time.Second)
		if idx != (pass - 1) {
			fmt.Printf("Ditemukan elemen terbesar %d (di indeks %d) untuk Pass ke-%d.\n", A[idx], idx, pass)
			fmt.Printf("Menukar A[%d] (%d) dengan A[%d] (%d).\n", pass-1, A[pass-1], idx, A[idx])
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			fmt.Println("Array setelah pertukaran:", A[:n])
			pass++
		} else {
			fmt.Printf("Elemen di posisi A[%d] (%d) sudah merupakan yang terbesar yang ditemukan di sisa array. Tidak perlu pertukaran.\n", pass-1, A[pass-1])
			pass++
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n--- PROSES SELESAI ---")
	fmt.Println("Array akhir terurut:", A[:n])
}
func PenjelasanInsertionSort(A *Tabint, n int, eff *int) {
	var i, pass int
	var temp int
	*eff = 0
	pass = 1
	for pass < n {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Memulai Pass ke-", pass, "---")
		fmt.Printf("Mengambil elemen A[%d]: %d untuk disisipkan ke bagian yang sudah terurut.\n", pass, A[pass])
		fmt.Println("Bagian yang sudah terurut:", A[0:pass])

		i = pass
		temp = A[pass]

		time.Sleep(1 * time.Second)
		fmt.Println("Memulai proses pergeseran dan perbandingan...")

		for i > 0 && temp > A[i-1] {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("   A[%d]: %d lebih kecil dari A[%d]: %d. Menggeser %d ke kanan.\n", i-1, A[i-1], i, temp, A[i-1])
			*eff++
			A[i] = A[i-1]
			i--
			fmt.Println("   Array sementara:", A[:pass+1])
		}

		time.Sleep(1 * time.Second)
		if i != pass {
			fmt.Printf("Elemen %d telah digeser ke posisi A[%d].\n", temp, i)
		} else {
			fmt.Printf("Elemen %d sudah berada di posisi yang benar (A[%d]). Tidak ada pergeseran.\n", temp, i)
		}
		A[i] = temp
		pass++
		fmt.Println("Array setelah pass:", A[:n])
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n--- PROSES SELESAI ---")
	fmt.Println("Array akhir terurut:", A[:n])
}
