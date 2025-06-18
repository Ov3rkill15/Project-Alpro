package quiz

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

type Quiz struct {
	Nama                         string
	ID                           string
	TotalScore                   float64
	GoLanguageScore              int
	PercabanganScore             int
	KonversiTipeDataScore        int
	OperasiMatematikaLogikaScore int
	PerulanganScore              int
	TipeDataGoScore              int
	VariableConstantScore        int
}

// NMAX mendefinisikan jumlah maksimum siswa yang dapat ditampung sistem.
const NMAX int = 50

var studentsData [NMAX]Quiz

func init() {
	InitStudentsData()
}
func InitStudentsData() {
	data := [45]Quiz{
		{Nama: "nathasyayuanmaharani", ID: "0001", TotalScore: 85.5},
		{Nama: "theodoreelvisestrada", ID: "0006", TotalScore: 92.0},
		{Nama: "dyahkusumawardani", ID: "0009", TotalScore: 78.5},
		{Nama: "azrielraihaneldovahartoto", ID: "0010", TotalScore: 88.0},
		{Nama: "muhammadilhamalifianda", ID: "0022", TotalScore: 95.0},
		{Nama: "alyaazizaputeri", ID: "0026", TotalScore: 76.0},
		{Nama: "ahmadabdansyakuro", ID: "0029", TotalScore: 81.5},
		{Nama: "fathurrahmanalfarizi", ID: "0035", TotalScore: 89.0},
		{Nama: "nuswantorosetyomukti", ID: "0040", TotalScore: 72.0},
		{Nama: "anggitacahyatihidayat", ID: "0041", TotalScore: 90.0},
		{Nama: "wibnuhijrahfranstio", ID: "0048", TotalScore: 83.0},
		{Nama: "meyshaprimiandita", ID: "0050", TotalScore: 91.0},
		{Nama: "muhamadfiqrihabibi", ID: "0056", TotalScore: 79.5},
		{Nama: "fitriacahyani", ID: "0060", TotalScore: 87.0},
		{Nama: "triansyahdaniswaraibrahim", ID: "0062", TotalScore: 93.0},
		{Nama: "rakhaabdillahalkautsar", ID: "0068", TotalScore: 75.0},
		{Nama: "avicenanaufallathif", ID: "0073", TotalScore: 86.0},
		{Nama: "naylaassyifa", ID: "0078", TotalScore: 94.0},
		{Nama: "williampetervanxnajoan", ID: "0084", TotalScore: 77.0},
		{Nama: "rayvanalifarlomahesworo", ID: "0087", TotalScore: 90.5},
		{Nama: "zaidansalamrojab", ID: "0088", TotalScore: 82.0},
		{Nama: "audreyfredileyhanas", ID: "0093", TotalScore: 89.5},
		{Nama: "muhammadnaelfadly", ID: "0096", TotalScore: 74.0},
		{Nama: "nairacahayaputridarmawansinaga", ID: "0100", TotalScore: 96.0},
		{Nama: "muhamadalwansuryadi", ID: "0104", TotalScore: 80.0},
		{Nama: "dhafyahmadzubaidi", ID: "0107", TotalScore: 91.5},
		{Nama: "muhammadfarisdhiyaylhaqsarbini", ID: "0117", TotalScore: 78.0},
		{Nama: "nursyadira", ID: "0123", TotalScore: 92.5},
		{Nama: "rayfitokrisnawijaya", ID: "0124", TotalScore: 84.0},
		{Nama: "mochammadrafirisqullah", ID: "0129", TotalScore: 88.5},
		{Nama: "iputugedeagastyakrisnawidartha", ID: "0134", TotalScore: 73.0},
		{Nama: "rendil", ID: "0137", TotalScore: 85.0},
		{Nama: "muhammadariqazzaki", ID: "0138", TotalScore: 90.0},
		{Nama: "edmundyuliusgantur", ID: "0155", TotalScore: 77.5},
		{Nama: "muhammadsayyidhuwaidi", ID: "0157", TotalScore: 86.5},
		{Nama: "muhdzuljalalwaliikramjalil", ID: "0160", TotalScore: 93.5},
		{Nama: "ramadhantangguhdefennder", ID: "0003", TotalScore: 81.0},
		{Nama: "adzkiyaputrirahmawan", ID: "0025", TotalScore: 89.0},
		{Nama: "fathimahradhiyya", ID: "0029", TotalScore: 79.0},
		{Nama: "rakanghazianadiwjaya", ID: "0034", TotalScore: 95.5},
		{Nama: "jihannabilamubarakah", ID: "0037", TotalScore: 87.5},
		{Nama: "admin", ID: "admin123", TotalScore: 0.0},
	}
	i := 0
	for i < len(data) && i < NMAX {
		studentsData[i] = data[i]
		i++
	}
}

func DisplayScoresMenu() {
	var choice string
	isDisplayingMenu := true // Flag untuk mengontrol loop menu tampilan skor

	for isDisplayingMenu {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          MENU TAMPILAN SKOR        ")
		fmt.Println("====================================")
		fmt.Println("Pilih opsi tampilan skor:")
		fmt.Println("1. Daftar Nilai (Urutan Asli)")
		fmt.Println("2. Daftar Nilai (Urut Menurun berdasarkan Total Poin)")
		fmt.Println("3. Daftar Nilai (Urut Menaik berdasarkan Total Poin)")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			ShowQuizScores("original") // Menampilkan skor dalam urutan asli
		case "2":
			ShowQuizScores("descending") // Menampilkan skor urut menurun
		case "3":
			ShowQuizScores("ascending") // Menampilkan skor urut menaik
		case "4":
			fmt.Println("Kembali ke menu utama...")
			isDisplayingMenu = false // Menghentikan loop menu skor
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Println("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

func ShowQuizScores(sortOrder string) {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          SKOR KUIS SISWA           ")

	// studentsToProcess adalah array lokal sementara untuk menampung data yang akan diurutkan.
	var studentsToProcess [NMAX]Quiz
	activeCount := 0 // Menghitung berapa banyak siswa aktif yang disalin
	i := 0
	for i < NMAX {
		if studentsData[i].ID != "" {
			studentsToProcess[activeCount] = studentsData[i]
			activeCount++
		}
		i++
	}

	fmt.Println("====================================")
	switch sortOrder {
	case "descending":
		fmt.Println("     (Diurutkan Menurun berdasarkan Total Poin) ")
		i = 1
		for i < activeCount { // Outer loop
			pass := studentsToProcess[i]
			j := i - 1
			canShift := true // Flag untuk mengontrol loop geser
			for canShift && j >= 0 {
				if studentsToProcess[j].TotalScore < pass.TotalScore {
					studentsToProcess[j+1] = studentsToProcess[j]
					j--
				} else {
					canShift = false // Hentikan geser jika kondisi tidak terpenuhi
				}
			}
			studentsToProcess[j+1] = pass
			i++
		}
	case "ascending":
		fmt.Println("     (Diurutkan Menaik berdasarkan Total Poin) ")
		i = 0
		for i < activeCount-1 { // Outer loop
			minIndex := i
			j := i + 1
			for j < activeCount { // Inner loop untuk mencari minIndex
				if studentsToProcess[j].TotalScore < studentsToProcess[minIndex].TotalScore {
					minIndex = j
				}
				j++
			}
			// Tukar elemen minimum yang ditemukan dengan elemen saat ini (studentsToProcess[i])
			temp := studentsToProcess[i]
			studentsToProcess[i] = studentsToProcess[minIndex]
			studentsToProcess[minIndex] = temp
			i++
		}
	default: // "original"
		fmt.Println("     (Urutan Asli) ")
	}

	fmt.Println(strings.Repeat("=", 60)) // Menyesuaikan lebar untuk tampilan konsol yang lebih baik

	// Menampilkan header kolom
	fmt.Printf("%-25s %-10s %-8s \n", "Nama", "ID", "Total")
	fmt.Println(strings.Repeat("-", 120))

	if activeCount == 0 { // Cek apakah ada siswa aktif yang ditemukan
		fmt.Println("Belum ada siswa yang terdaftar atau mengambil kuis.")
	} else {
		// Menampilkan data siswa dari array lokal studentsToProcess yang sudah diurutkan (atau tidak diurutkan jika "original")
		i = 0
		for i < activeCount {
			student := studentsToProcess[i]
			fmt.Printf("%-25s %-10s %-8.1f \n", // Menggunakan .1f untuk float dengan satu desimal
				student.Nama,
				student.ID,
				student.TotalScore,
			)
			i++
		}
	}

	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("Opsi Pencarian:")
	fmt.Print("Ingin mencari siswa berdasarkan nama? (y/n): ")
	var searchChoice string
	fmt.Scanln(&searchChoice)

	if strings.ToLower(searchChoice) == "y" {
		fmt.Print("Masukkan nama siswa yang dicari: ")
		var searchName string
		fmt.Scanln(&searchName)

		foundStudentIndex := -1 // Default: tidak ditemukan
		isSearching := true     // Flag untuk mengontrol loop pencarian

		fmt.Println("\n--- Melakukan Pencarian Sekuensial Berdasarkan Nama ---")
		i = 0
		for i < activeCount && isSearching { // Loop selama masih dalam batas dan belum ditemukan
			if strings.EqualFold(strings.ToLower(studentsToProcess[i].Nama), strings.ToLower(searchName)) { // Perbandingan tidak case-sensitive
				foundStudentIndex = i
				isSearching = false // Set isSearching ke false untuk menghentikan loop pada iterasi berikutnya
			}
			i++
		}

		fmt.Println(strings.Repeat("-", 40))
		if foundStudentIndex != -1 {
			foundStudent := studentsToProcess[foundStudentIndex]
			fmt.Println("Siswa ditemukan:")
			fmt.Printf("Nama: %s\n", foundStudent.Nama)
			fmt.Printf("ID: %s\n", foundStudent.ID)
			fmt.Printf("Total Score: %.1f\n", foundStudent.TotalScore)
		} else {
			fmt.Println("Siswa dengan nama '" + searchName + "' tidak ditemukan.")
		}
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("Tekan Enter untuk melanjutkan...")
		fmt.Scanln()
	}

	fmt.Println("Tekan Enter untuk kembali ke menu sebelumnya...")
	fmt.Scanln()
}

// RegisterNewStudentFlow mengelola alur untuk pendaftaran siswa baru.
func RegisterNewStudentFlow() {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          DAFTAR SISWA BARU         ")
	fmt.Println("====================================")

	var newNama, newID string
	fmt.Print("Masukkan Nama Lengkap Anda: ")
	fmt.Scanln(&newNama)

	registrationSuccessful := false // Flag untuk menunjukkan apakah pendaftaran berhasil

	for !registrationSuccessful { // Loop akan terus berjalan sampai pendaftaran berhasil
		isValidInputRound := true // Flag untuk validasi input ID di putaran saat ini

		fmt.Print("Masukkan ID Anda (contoh: 011, QS011, dll.): ")
		fmt.Scanln(&newID)

		// Validasi input ID
		if strings.TrimSpace(newID) == "" {
			fmt.Println("ID tidak boleh kosong. Silakan masukkan ID yang valid.")
			isValidInputRound = false
		}
		if isValidInputRound && strings.Contains(newID, " ") { // Cek spasi hanya jika input sudah dianggap valid sejauh ini
			fmt.Println("ID tidak boleh mengandung spasi. Silakan masukkan ID yang valid.")
			isValidInputRound = false
		}
		if isValidInputRound && GetStudentIndexByID(newID) != -1 { // Cek ID sudah ada hanya jika input sudah dianggap valid sejauh ini
			fmt.Println("ID tersebut sudah digunakan. Silakan gunakan ID lain.")
			isValidInputRound = false
		}

		if isValidInputRound { // Jika semua validasi lolos
			indexAdded := AddNewStudent(newNama, newID) // Mencoba mendaftarkan siswa baru

			if indexAdded != -1 {
				fmt.Printf("\nSelamat datang, %s! Data Anda (ID: %s) telah berhasil didaftarkan.\n", newNama, newID)
				fmt.Println("Sekarang Anda bisa mengambil kuis menggunakan ID Anda.")
				registrationSuccessful = true // Pendaftaran berhasil, keluar dari loop
			} else {
				fmt.Println("Maaf, kapasitas pendaftar sudah penuh (" + fmt.Sprintf("%d", NMAX) + " orang).")
				registrationSuccessful = true // Kapasitas penuh, keluar dari loop
			}
		}
	}
	PromptReturnToPreviousMenu()
}

func StartQuizSession(quizSpecificFunction func(nilaiTemp *float64) bool, materialName string) {
	studentIndex := GetStudentIndexPrompt() // Mendapatkan indeks siswa berdasarkan ID
	if studentIndex != -1 {
		quizSpecificFunction(&studentsData[studentIndex].TotalScore)
	} else {
		fmt.Println("ID siswa tidak ditemukan atau tidak valid. Kuis " + materialName + " tidak dapat dimulai.")
		fmt.Println("Silakan daftar terlebih dahulu (pilih 'Daftar / Masukkan Data Baru') atau gunakan ID yang terdaftar.")
		PromptReturnToPreviousMenu() // Kembali ke menu sebelumnya setelah pesan kesalahan
	}
}

func GetStudentIndexByID(studentID string) int {
	i := 0
	foundIndex := -1 // Default: tidak ditemukan
	isFound := false // Flag untuk mengontrol loop dan menunjukkan apakah ditemukan

	for i < NMAX && !isFound { // Loop selama masih dalam batas array dan belum ditemukan
		if strings.ToLower(studentsData[i].ID) == strings.ToLower(studentID) {
			foundIndex = i // Ditemukan, simpan indeks
			isFound = true // Set flag ke true untuk menghentikan loop pada iterasi berikutnya
		}
		i++
	}
	return foundIndex // Mengembalikan indeks yang ditemukan atau -1
}

// AddNewStudent menambahkan data siswa baru ke array studentsData.
// Mengembalikan indeks siswa yang baru ditambahkan, atau -1 jika pendaftaran gagal (kapasitas penuh).
func AddNewStudent(nama string, id string) int {
	emptyIndex := -1
	i := 0
	foundEmptySlot := false // Flag untuk mengontrol loop agar berhenti jika slot kosong ditemukan

	for i < NMAX && !foundEmptySlot { // Loop selama masih dalam batas array dan belum menemukan slot kosong
		if studentsData[i].ID == "" { // Cari slot kosong berdasarkan ID yang kosong
			emptyIndex = i
			foundEmptySlot = true // Slot kosong ditemukan, set flag ke true
		}
		i++
	}

	if emptyIndex == -1 { // Tidak ada slot kosong (kapasitas penuh)
		return -1
	}

	// Inisialisasi siswa baru di slot yang ditemukan
	studentsData[emptyIndex].Nama = nama
	studentsData[emptyIndex].ID = id
	studentsData[emptyIndex].TotalScore = 0.0 // Nilai default 0.0 saat daftar
	return emptyIndex
}

// PromptReturnToPreviousMenu meminta pengguna untuk menekan 'm' untuk kembali ke menu sebelumnya.
func PromptReturnToPreviousMenu() {
	var choice string
	isInputHandled := false // Flag untuk mengontrol loop input

	for !isInputHandled { // Loop selama input belum ditangani
		fmt.Print("\nTekan 'm' untuk kembali ke menu sebelumnya: ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "m" {
			isInputHandled = true // Input valid, keluar dari loop
		} else {
			fmt.Println("Pilihan tidak valid. Silakan masukkan 'm'.")
		}
	}
}

// GetStudentIndexPrompt meminta ID siswa dari pengguna dan mencari indeksnya.
// Mengembalikan indeks siswa jika ditemukan, atau -1 jika tidak.
func GetStudentIndexPrompt() int {
	var studentID string
	fmt.Print("Masukkan ID siswa Anda (contoh: 0001): ")
	fmt.Scanln(&studentID)

	return GetStudentIndexByID(studentID) // Menggunakan fungsi pencarian yang disesuaikan
}
