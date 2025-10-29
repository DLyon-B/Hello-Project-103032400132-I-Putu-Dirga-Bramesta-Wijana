package main

import (
	"fmt"
	"time"
)

const NMAX = 100

type BahanMakanan struct {
	Nama        string
	Jumlah      int
	Kedaluwarsa time.Time
}

type ArrayBahan struct {
	Data [NMAX]BahanMakanan
	N    int
}

// Menampilkan semua stok bahan
func tampilkanStok(A ArrayBahan) {
	var i int
	
	fmt.Println("Stok Bahan Makanan:")
	for i = 0; i < A.N; i++ {
		fmt.Printf("%d. %s - %d (Kedaluwarsa: %s)\n", i+1, A.Data[i].Nama, A.Data[i].Jumlah, A.Data[i].Kedaluwarsa.Format("2006-01-02"))
	}
}

// Fungsi pembanding string tanpa strings.EqualFold
func samaTanpaCase(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		pa, pb := a[i], b[i]
		if pa >= 'A' && pa <= 'Z' {
			pa += 32
		}
		if pb >= 'A' && pb <= 'Z' {
			pb += 32
		}
		if pa != pb {
			return false
		}
	}
	return true
}

// Sequential Search
func cariSequential(A ArrayBahan, nama string) int {
	var i int
	
	for i = 0; i < A.N; i++ {
		if samaTanpaCase(A.Data[i].Nama, nama) {
			return i
		}
	}
	return -1
}

// Binary Search (harus terurut berdasarkan nama ascending)
func binarySearch(A ArrayBahan, nama string) int {
	var low, high, mid int
	
	low, high = 0, A.N-1
	for low <= high {
		mid = (low + high) / 2
		if samaTanpaCase(A.Data[mid].Nama, nama) {
			return mid
		} else if A.Data[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Selection Sort berdasarkan jumlah
func selectionSortJumlah(A *ArrayBahan, ascending bool) {
	var i, j, idx int
	
	for i = 0; i < A.N-1; i++ {
		idx = i
		for j = i + 1; j < A.N; j++ {
			if (ascending && A.Data[j].Jumlah < A.Data[idx].Jumlah) || (!ascending && A.Data[j].Jumlah > A.Data[idx].Jumlah) {
				idx = j
			}
		}
		A.Data[i], A.Data[idx] = A.Data[idx], A.Data[i]
	}
}

// Insertion Sort berdasarkan nama
func insertionSortNama(A *ArrayBahan, ascending bool) {
	var i, j int
	var key BahanMakanan
	
	for i = 1; i < A.N; i++ {
		key = A.Data[i]
		j = i - 1
		for j >= 0 && ((ascending && A.Data[j].Nama > key.Nama) || (!ascending && A.Data[j].Nama < key.Nama)) {
			A.Data[j+1] = A.Data[j]
			j--
		}
		A.Data[j+1] = key
	}
}

// Tambah bahan
func tambahBahan(A *ArrayBahan, nama string, jumlah int, kadaluwarsa string) {
	var tgl time.Time
	var err error
	
	if A.N >= NMAX {
		fmt.Println("Stok penuh!")
		return
	}
	tgl, err = time.Parse("2006-01-02", kadaluwarsa)
	if err != nil {
		fmt.Println("Format tanggal salah.")
		return
	}
	A.Data[A.N] = BahanMakanan{nama, jumlah, tgl}
	A.N++
}

// Ubah data (sequential search)
func ubahBahan(A *ArrayBahan, nama string, jumlahBaru int) {
	var idx int
	
	idx = cariSequential(*A, nama)
	if idx != -1 {
		A.Data[idx].Jumlah = jumlahBaru
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Bahan tidak ditemukan.")
	}
}

// Hapus data (binary search setelah diurutkan)
func hapusBahan(A *ArrayBahan, nama string) {
	var i, idx int
	
	insertionSortNama(A, true)
	idx = binarySearch(*A, nama)
	if idx != -1 {
		for i = idx; i < A.N-1; i++ {
			A.Data[i] = A.Data[i+1]
		}
		A.N--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Main Program
func main() {
	var stok ArrayBahan

	// Data awal
	tambahBahan(&stok, "Beras", 10, "2025-12-01")
	tambahBahan(&stok, "Gula", 5, "2025-11-01")
	tambahBahan(&stok, "Minyak", 3, "2025-10-15")
	tambahBahan(&stok, "Sayur", 8, "2025-06-15")

	var pilihan int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan Stok")
		fmt.Println("2. Tambah Bahan Makanan")
		fmt.Println("3. Ubah Jumlah ")
		fmt.Println("4. Hapus Bahan ")
		fmt.Println("5. Sort Jumlah ")
		fmt.Println("6. Sort Nama ")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanStok(stok)
		case 2:
			var nama, tgl string
			var jumlah int
			fmt.Print("Nama: ")
			fmt.Scan(&nama)
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)
			fmt.Print("Tanggal Kadaluwarsa (YYYY-MM-DD): ")
			fmt.Scan(&tgl)
			tambahBahan(&stok, nama, jumlah, tgl)
		case 3:
			var nama string
			var jumlah int
			fmt.Print("Nama bahan: ")
			fmt.Scan(&nama)
			fmt.Print("Jumlah baru: ")
			fmt.Scan(&jumlah)
			ubahBahan(&stok, nama, jumlah)
		case 4:
			var nama string
			fmt.Print("Nama bahan yang dihapus: ")
			fmt.Scan(&nama)
			hapusBahan(&stok, nama)
		case 5:
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			selectionSortJumlah(&stok, asc == 1)
			fmt.Println("Stok diurutkan berdasarkan jumlah.")
		case 6:
			var asc int
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			insertionSortNama(&stok, asc == 1)
			fmt.Println("Stok diurutkan berdasarkan nama.")
		case 7:
			fmt.Println("Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
