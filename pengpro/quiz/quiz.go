package quiz

import (
	"Project-Alpro/atribut"
	"fmt"
	"strings"
)

const NMAX int = 15

// DataQuiz is the global dataquiz array for your application
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

// FindStudentIndex finds a student's index by ID. It's exported.
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

// PromptKembaliToMain prompts the user to return to the submenu.
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

// HandleNewStudentRegistration mengelola alur pendaftaran siswa baru.
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
