package quiz

import (
	"Project-Alpro/atribut" // Assuming this path is correct for your project
	"fmt"
	"strings"
)

type Quiz struct {
	Nama                         string
	ID                           string
	TotalScore                   int
	GoLanguageScore              int
	PercabanganScore             int
	KonversiTipeDataScore        int
	OperasiMatematikaLogikaScore int
	PerulanganScore              int
	TipeDataGoScore              int
	VariableConstantScore        int
}

// User represents a user with a username, password, and score.
type User struct {
	Username, Password string
	Score              float64
}

// NMAX defines the maximum number of users the system can hold.
const NMAX int = 45

// Daftarnama is a fixed-size array to store User structs.
type Daftarnama [NMAX]User

// Global variables for user management.
var users Daftarnama
var jumlahPengguna int // 'jumlahPengguna' means 'number of users'

// initUsers initializes the 'users' array with a predefined set of user data.
// This function should ideally be called once at the application start.
var DataQuiz [NMAX]atribut.Quiz = [NMAX]atribut.Quiz{
	{Nama: "Budi Santoso", ID: "QS001", TotalScore: 0},
	{Nama: "Siti Aminah", ID: "QS002", TotalScore: 0},
	{Nama: "Joko Susilo", ID: "QS003", TotalScore: 0},
	{Nama: "Dewi Lestari", ID: "QS004", TotalScore: 0},
	{Nama: "Agus Ramadhan", ID: "QS005", TotalScore: 0},
	{Nama: "Putri Indah", ID: "QS006", TotalScore: 0},
	{Nama: "Rizky Pratama", ID: "QS007", TotalScore: 0},
	{Nama: "Nurul Hidayah", ID: "QS008", TotalScore: 0},
	{Nama: "Faisal Rahman", ID: "QS009", TotalScore: 0},
	{Nama: "Linda Wijaya", ID: "QS010", TotalScore: 0},
}

func initUsers() {
	// Predefined user data
	// Using a temporary array directly for initialization without 'make' or 'copy'
	var data = [42]User{
		{"admin", "admin123", 0.0},
		{"nathasyayuanmaharani", "0001", 85.5},
		{"theodoreelvisestrada", "0006", 92.0},
		{"dyahkusumawardani", "0009", 78.5},
		{"azrielraihaneldovahartoto", "0010", 88.0},
		{"muhammadilhamalifianda", "0022", 95.0},
		{"alyaazizaputeri", "0026", 76.0},
		{"ahmadabdansyakuro", "0029", 81.5},
		{"fathurrahmanalfarizi", "0035", 89.0},
		{"nuswantorosetyomukti", "0040", 72.0},
		{"anggitacahyatihidayat", "0041", 90.0},
		{"wibnuhijrahfranstio", "0048", 83.0},
		{"meyshaprimiandita", "0050", 91.0},
		{"muhamadfiqrihabibi", "0056", 79.5},
		{"fitriacahyani", "0060", 87.0},
		{"triansyahdaniswaraibrahim", "0062", 93.0},
		{"rakhaabdillahalkautsar", "0068", 75.0},
		{"avicenanaufallathif", "0073", 86.0},
		{"naylaassyifa", "0078", 94.0},
		{"williampetervanxnajoan", "0084", 77.0},
		{"rayvanalifarlomahesworo", "0087", 90.5},
		{"zaidansalamrojab", "0088", 82.0},
		{"audreyfredileyhanas", "0093", 89.5},
		{"muhammadnaelfadly", "0096", 74.0},
		{"nairacahayaputridarmawansinaga", "0100", 96.0},
		{"muhamadalwansuryadi", "0104", 80.0},
		{"dhafyahmadzubaidi", "0107", 91.5},
		{"muhammadfarisdhiyaylhaqsarbini", "0117", 78.0},
		{"nursyadira", "0123", 92.5},
		{"rayfitokrisnawijaya", "0124", 84.0},
		{"mochammadrafirisqullah", "0129", 88.5},
		{"iputugedeagastyakrisnawidartha", "0134", 73.0},
		{"rendil", "0137", 85.0},
		{"muhammadariqazzaki", "0138", 90.0},
		{"edmundyuliusgantur", "0155", 77.5},
		{"muhammadsayyidhuwaidi", "0157", 86.5},
		{"muhdzuljalalwaliikramjalil", "0160", 93.5},
		{"ramadhantangguhdefennder", "0003", 81.0},
		{"adzkiyaputrirahmawan", "0025", 89.0},
		{"fathimahradhiyya", "0029", 79.0},
		{"rakanghazianadiwjaya", "0034", 95.5},
		{"jihannabilamubarakah", "0037", 87.5},
	}

	jumlahPengguna = 0
	var i int = 0
	for i < len(data) { // Loop without 'break'
		if i < NMAX {
			users[i] = data[i]
			jumlahPengguna++
		}
		i++
	}
}

// DisplayScoresMenu menampilkan submenu untuk pilihan tampilan skor.
func DisplayScoresMenu() {
	var choice string
	berhenti := true
	for berhenti {
		atribut.ClearScreen()
		fmt.Println("====================================")
		fmt.Println("          MENU TAMPILAN SKOR        ")
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
			displayQuizScores("original") // Tampilan asli
		case "2":
			displayQuizScores("descending") // Urut menurun
		case "3":
			displayQuizScores("ascending") // Urut menaik
		case "4":
			fmt.Println("Kembali ke menu utama...")
			berhenti = false
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Println("Tekan Enter untuk melanjutkan...")
			fmt.Scanln()
		}
	}
}

// displayQuizScores menampilkan semua skor kuis untuk semua siswa.
// Menerima parameter `sortOrder` untuk menentukan urutan tampilan:
// "original", "descending", "ascending".
func displayQuizScores(sortOrder string) {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          SKOR KUIS SISWA           ")

	// 1. Buat array lokal baru untuk menampung data yang akan diurutkan.
	// Inisialisasi dengan semua slot kosong (zero value).
	var studentsToSort [NMAX]atribut.Quiz
	activeCount := 0 // Menghitung berapa banyak siswa aktif yang disalin

	// 2. Salin hanya siswa yang aktif ke array lokal.
	// Ini juga akan "memadatkan" data ke awal array studentsToSort.
	for i := 0; i < NMAX; i++ {
		if DataQuiz[i].Nama != "" {
			studentsToSort[activeCount] = DataQuiz[i]
			activeCount++
		}
	}

	// 3. Lakukan pengurutan pada bagian yang berisi siswa aktif dari studentsToSort
	switch sortOrder {
	case "descending":
		fmt.Println("     (Diurutkan Menurun berdasarkan Total Poin) ")
		// Insertion Sort untuk urut menurun, disesuaikan dengan format while-loop
		// Kamus: i, pass : integer; temp : types.Quiz
		pass := 1                // pass <- 1
		for pass < activeCount { // while pass < activeCount do (setara dengan pass <= n-1 di pseudocode jika n=activeCount)
			// { Pencarian indeks yang tepat untuk elemen }
			i := pass
			temp := studentsToSort[pass] // temp <- A[pass]

			// while i > 0 and A[i-1] < temp.TotalScore do (untuk menurun)
			for i > 0 && studentsToSort[i-1].TotalScore < temp.TotalScore {
				studentsToSort[i] = studentsToSort[i-1] // A[i] <- A[i-1]
				i--                                     // i <- i - 1
			}
			// { Menempatkan elemen pada lokasi tersebut}
			studentsToSort[i] = temp // A[i] <- temp

			pass++ // pass <- pass + 1
		}
	case "ascending":
		fmt.Println("     (Diurutkan Menaik berdasarkan Total Poin) ")
		pass := 0                  // pass <- 0
		for pass < activeCount-1 { // while pass <= n-1 do (simulasi)
			// { 1. Pencarian nilai idx ekstrim (minimum) }
			idx := pass           // idx <- pass
			i := pass + 1         // i <- pass + 1
			for i < activeCount { // while i < n do (simulasi)
				if studentsToSort[i].TotalScore < studentsToSort[idx].TotalScore { // Kondisi untuk mencari MINIMUM
					idx = i
				}
				i++ // i <- i + 1
			}
			// { 2. Pertukaran }
			temp := studentsToSort[idx]                // temp <- A[idx]
			studentsToSort[idx] = studentsToSort[pass] // A[idx] <- A[pass]
			studentsToSort[pass] = temp                // A[pass] <- temp

			pass++ // pass <- pass + 1
		}
	default: // "original" atau jika ada nilai sortOrder lain yang tidak dikenali
		fmt.Println("     (Urutan Asli) ")
	}
	fmt.Println("====================================")

	// 4. Tampilkan data
	fmt.Printf("%-15s %-10s %-8s %-8s %-12s %-8s %-8s %-10s %-8s %s\n",
		"Nama", "ID", "Total", "GoLang", "Percabangan", "KonvTD", "OpML", "Perulangan", "TDGo", "VarConst")
	fmt.Println(strings.Repeat("-", 120))

	if activeCount == 0 { // Cek apakah ada siswa aktif yang ditemukan
		fmt.Println("Belum ada siswa yang terdaftar atau mengambil kuis.")
	} else {
		if sortOrder == "original" {
			// Tampilkan dari array global asli untuk menjaga urutan awal
			for i := 0; i < NMAX; i++ {
				if DataQuiz[i].Nama != "" {
					fmt.Printf("%-15s %-10s %-8d %-8d %-12d %-8d %-8d %-10d %-8d %d\n",
						DataQuiz[i].Nama,
						DataQuiz[i].ID,
						DataQuiz[i].TotalScore,
						DataQuiz[i].GoLanguageScore,
						DataQuiz[i].PercabanganScore,
						DataQuiz[i].KonversiTipeDataScore,
						DataQuiz[i].OperasiMatematikaLogikaScore,
						DataQuiz[i].PerulanganScore,
						DataQuiz[i].TipeDataGoScore,
						DataQuiz[i].VariableConstantScore,
					)
				}
			}
		} else {
			// Tampilkan dari array lokal yang sudah diurutkan (studentsToSort)
			// Hanya loop sebanyak activeCount untuk menghindari slot kosong
			for i := 0; i < activeCount; i++ {
				student := studentsToSort[i] // Ambil data dari array yang diurutkan
				fmt.Printf("%-15s %-10s %-8d %-8d %-12d %-8d %-8d %-10d %-8d %d\n",
					student.Nama,
					student.ID,
					student.TotalScore,
					student.GoLanguageScore,
					student.PercabanganScore,
					student.KonversiTipeDataScore,
					student.OperasiMatematikaLogikaScore,
					student.PerulanganScore,
					student.TipeDataGoScore,
					student.VariableConstantScore,
				)
			}
		}
	}

	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("Tekan Enter untuk kembali ke menu sebelumnya...")
	fmt.Scanln()
}

func HandleNewStudentRegistration() {
	atribut.ClearScreen()
	fmt.Println("====================================")
	fmt.Println("          DAFTAR SISWA BARU         ")
	fmt.Println("====================================")

	var newNama, newID string
	fmt.Print("Masukkan Nama Lengkap Anda: ")
	fmt.Scanln(&newNama)

	registrationAttempted := false  // Flag untuk menunjukkan apakah sudah mencoba mendaftar
	registrationSuccessful := false // Flag untuk menunjukkan apakah pendaftaran berhasil

	// Loop akan terus berjalan sampai pendaftaran berhasil atau tidak mungkin lagi
	for !registrationSuccessful && !registrationAttempted { // Loop terus selama belum berhasil dan belum ada upaya
		// Jika sudah ada upaya, berarti loop ini dari percobaan sebelumnya yang gagal
		// dan kita perlu meminta ID lagi kecuali sudah dinyatakan gagal total
		if registrationAttempted {
			fmt.Print("Masukkan ID Anda (contoh: QS011): ") // Minta ID lagi setelah kegagalan
			fmt.Scanln(&newID)
		} else {
			// Pertama kali masuk loop, minta ID tanpa pesan error sebelumnya
			fmt.Print("Masukkan ID Anda (contoh: QS011): ")
			fmt.Scanln(&newID)
			registrationAttempted = true // Tandai bahwa sudah ada upaya pertama
		}

		isValidInput := true // Flag untuk validasi input ID
		if strings.TrimSpace(newID) == "" {
			fmt.Println("ID tidak boleh kosong. Silakan masukkan ID yang valid.")
			isValidInput = false
		} else if strings.Contains(newID, " ") {
			fmt.Println("ID tidak boleh mengandung spasi. Silakan masukkan ID yang valid.")
			isValidInput = false
		}

		if isValidInput {
			newIndex := RegisterNewStudent(newNama, newID)

			if newIndex != -1 {
				// Pendaftaran berhasil
				fmt.Printf("\nSelamat datang, %s! Data Anda (ID: %s) telah berhasil didaftarkan.\n", newNama, newID)
				fmt.Println("Sekarang Anda bisa mengambil kuis menggunakan ID Anda.")
				registrationSuccessful = true // Atur flag untuk keluar dari loop
			} else {
				// Pendaftaran gagal, tentukan alasannya
				if FindStudentIndex(newID) != -1 { // Cek apakah ID sudah ada
					fmt.Println("ID tersebut sudah digunakan. Silakan gunakan ID lain.")
					// Loop akan terus berjalan untuk meminta ID baru
				} else { // Jika bukan karena ID ada, berarti kapasitas penuh
					fmt.Println("Maaf, kapasitas pendaftar sudah penuh (" + fmt.Sprintf("%d", NMAX) + " orang).")
					registrationSuccessful = true // Atur flag untuk keluar dari loop karena sudah penuh
				}
			}
		}
		// Jika isValidInput false, loop akan berulang dan meminta ID lagi
	}
	PromptKembaliToMain()
}

// HandleQuizStart mengelola alur untuk memulai kuis.
func HandleQuizStart(quizSpecificFunction func(quizDataArray *[NMAX]atribut.Quiz, index int), materialName string) {
	studentIndex := GetStudentIndexPrompt()
	if studentIndex != -1 {
		quizSpecificFunction(&DataQuiz, studentIndex)
	} else {
		fmt.Println("ID siswa tidak ditemukan atau tidak valid. Kuis " + materialName + " tidak dapat dimulai.")
		fmt.Println("Silakan daftar terlebih dahulu (pilih 'Daftar / Masukkan Data Baru') atau gunakan ID yang terdaftar.")
		PromptKembaliToMain()
	}
}

func FindStudentIndex(studentID string) int {
	for i := 0; i < NMAX; i++ {
		if strings.ToLower(DataQuiz[i].ID) == strings.ToLower(studentID) {
			return i
		}
	}
	return -1
}

// RegisterNewStudent adds a new student's data to the DataQuiz array. It's exported.
func RegisterNewStudent(nama string, id string) int {
	if FindStudentIndex(id) != -1 {
		return -1
	}

	emptyIndex := -1
	foundEmptySlot := false
	for i := 0; i < NMAX && !foundEmptySlot; i++ {
		if DataQuiz[i].ID == "" {
			emptyIndex = i
			foundEmptySlot = true
		}
	}

	if !foundEmptySlot {
		return -1
	}

	DataQuiz[emptyIndex].Nama = nama
	DataQuiz[emptyIndex].ID = id
	DataQuiz[emptyIndex].TotalScore = 0

	return emptyIndex
}

func PromptKembaliToMain() {
	var choice string
	berhenti3 := false
	for !berhenti3 {
		fmt.Print("\nTekan 'm' untuk kembali ke submenu: ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "m" {
			berhenti3 = true
		} else {
			fmt.Println("Pilihan tidak valid. Silakan masukkan 'm'.")
		}
	}
}

// GetStudentIndexPrompt meminta ID siswa dari pengguna dan mencari indeksnya.
func GetStudentIndexPrompt() int {
	var studentID string
	fmt.Print("Masukkan ID siswa Anda (contoh: QS001): ")
	fmt.Scanln(&studentID)

	return FindStudentIndex(studentID)
}
