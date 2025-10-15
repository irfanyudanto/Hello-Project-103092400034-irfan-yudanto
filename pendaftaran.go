// Ditambahkan oleh Irfan Yudanto - 103092400034
//Hanya test aja sih ini


package main

import (
	"fmt"
	"strings"
)

const NMAX int = 100

type Jurusan struct {
	Kode string
	Nama string
}

type Mahasiswa struct {
	Nama      string
	Nim       string
	Jurusan   string
	NilaiTest float64
	Status    string
}

var daftarJurusan [NMAX]Jurusan
var jumlahJurusan int

var daftarMahasiswa [NMAX]Mahasiswa
var jumlahMahasiswa int

func tambahJurusan(kode, nama string) {
	if jumlahJurusan < NMAX {
		daftarJurusan[jumlahJurusan] = Jurusan{kode, nama}
		jumlahJurusan++
	}
}

func hapusJurusan(kode string) {
	for i := 0; i < jumlahJurusan; i++ {
		if daftarJurusan[i].Kode == kode {
			for j := i; j < jumlahJurusan-1; j++ {
				daftarJurusan[j] = daftarJurusan[j+1]
			}
			jumlahJurusan--
		}
	}
}

func tambahMahasiswa(nim, nama, jurusan string) {
	if jumlahMahasiswa < NMAX {
		daftarMahasiswa[jumlahMahasiswa] = Mahasiswa{
			Nama:      nama,
			Nim:       nim,
			Jurusan:   jurusan,
			NilaiTest: 0,
			Status:    "Belum Diuji",
		}
		jumlahMahasiswa++
	}
}

func hapusMahasiswa(nim string) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].Nim == nim {
			for j := i; j < jumlahMahasiswa-1; j++ {
				daftarMahasiswa[j] = daftarMahasiswa[j+1]
			}
			jumlahMahasiswa--
		}
	}
}

func isiNilaiTes(nim string, nilai float64) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].Nim == nim {
			daftarMahasiswa[i].NilaiTest = nilai
			if nilai >= 60 {
				daftarMahasiswa[i].Status = "Diterima"
			} else {
				daftarMahasiswa[i].Status = "Ditolak"
			}
		}
	}
}

func tampilkanMahasiswaByStatus(status string) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].Status == status {
			fmt.Print(daftarMahasiswa[i].Nim, " - ", daftarMahasiswa[i].Nama, " - ", daftarMahasiswa[i].Jurusan, " - ", daftarMahasiswa[i].NilaiTest, "\n")
		}
	}
}

func selectionSortNama(ascending bool) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		idx := i
		for j := i + 1; j < jumlahMahasiswa; j++ {
			perbandingan := strings.Compare(daftarMahasiswa[j].Nama, daftarMahasiswa[idx].Nama)
			if (ascending && perbandingan < 0) || (!ascending && perbandingan > 0) {
				idx = j
			}
		}
		temp := daftarMahasiswa[i]
		daftarMahasiswa[i] = daftarMahasiswa[idx]
		daftarMahasiswa[idx] = temp
	}
}

func binarySearchNIM(nim string, ascending bool) int {
	low := 0
	high := jumlahMahasiswa - 1
	for low <= high {
		mid := (low + high) / 2
		if daftarMahasiswa[mid].Nim == nim {
			return mid
		}
		if (ascending && daftarMahasiswa[mid].Nim < nim) || (!ascending && daftarMahasiswa[mid].Nim > nim) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func keluarProgram(pilihan *int) {
	fmt.Println("Keluar dari program.")
	*pilihan = 99
}

func main() {
	var pilihan int

	for pilihan != 99 {
		fmt.Println()
		fmt.Println("=== MENU ===")
		fmt.Println("1. Tambah Jurusan")
		fmt.Println("2. Tambah Mahasiswa")
		fmt.Println("3. Tampilkan Semua Mahasiswa")
		fmt.Println("4. Isi Nilai Tes")
		fmt.Println("5. Hapus Mahasiswa")
		fmt.Println("6. Hapus Jurusan")
		fmt.Println("7. Tampilkan Mahasiswa Diterima")
		fmt.Println("8. Urutkan Mahasiswa (Nama Ascending)")
		fmt.Println("9. Cari Mahasiswa (Binary Search Nim)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			var kode, nama string
			fmt.Print("Kode Jurusan: ")
			fmt.Scan(&kode)
			fmt.Print("Nama Jurusan: ")
			fmt.Scan(&nama)
			tambahJurusan(kode, nama)
		}

		if pilihan == 2 {
			var nim, nama, jurusan string
			fmt.Print("NIM: ")
			fmt.Scan(&nim)
			fmt.Print("Nama: ")
			fmt.Scan(&nama)
			fmt.Print("Kode Jurusan: ")
			fmt.Scan(&jurusan)
			tambahMahasiswa(nim, nama, jurusan)
		}

		if pilihan == 3 {
			fmt.Println("=== Daftar Mahasiswa ===")
			for i := 0; i < jumlahMahasiswa; i++ {
				fmt.Print(daftarMahasiswa[i].Nim, " - ", daftarMahasiswa[i].Nama, " - ", daftarMahasiswa[i].Jurusan, " - ", daftarMahasiswa[i].NilaiTest, "\n")
			}
		}

		if pilihan == 4 {
			var nim string
			var nilai float64
			fmt.Print("NIM Mahasiswa: ")
			fmt.Scan(&nim)
			fmt.Print("Nilai Tes: ")
			fmt.Scan(&nilai)
			isiNilaiTes(nim, nilai)
		}

		if pilihan == 5 {
			var nim string
			fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
			fmt.Scan(&nim)
			hapusMahasiswa(nim)
		}

		if pilihan == 6 {
			var kode string
			fmt.Print("Masukkan kode jurusan yang ingin dihapus: ")
			fmt.Scan(&kode)
			hapusJurusan(kode)
		}

		if pilihan == 7 {
			fmt.Println("=== Mahasiswa Diterima ===")
			tampilkanMahasiswaByStatus("Diterima")
		}

		if pilihan == 8 {
			selectionSortNama(true)
			fmt.Println("=== Mahasiswa Diurutkan Berdasarkan Nama (A-Z) ===")
			for i := 0; i < jumlahMahasiswa; i++ {
				m := daftarMahasiswa[i]
				fmt.Print(m.Nim, " ", m.Nama, " ", m.Jurusan, " ", m.NilaiTest, " ", m.Status, "\n")
			}
		}

		if pilihan == 9 {
			var nim string
			fmt.Print("Masukkan NIM yang ingin dicari: ")
			fmt.Scan(&nim)
			selectionSortNama(true)
			idx := binarySearchNIM(nim, true)
			if idx != -1 {
				m := daftarMahasiswa[idx]
				fmt.Print("Ditemukan: ", m.Nim, " ", m.Nama, " ", m.Jurusan, " ", m.NilaiTest, " ", m.Status, "\n")
			} else {
				fmt.Println("Mahasiswa tidak ditemukan.")
			}
		}

		if pilihan == 0 {
			keluarProgram(&pilihan)
		}
	}
}
