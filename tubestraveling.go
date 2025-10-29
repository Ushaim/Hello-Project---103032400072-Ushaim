/* 
Apa itu Aplikasi Budget Traveling?
Aplikasi ini adalah program komputer yang membantu Anda menghitung dan mengelola biaya perjalanan ke berbagai tujuan. Program ini bisa menentukan kelas perjalanan berdasarkan biaya yang Anda miliki dan menghitung total biaya termasuk biaya tambahan seperti asuransi dan administrasi.

Bagian-bagian Utama Aplikasi
1. Struktur Data Biaya (Biaya)
Ini adalah format untuk menyimpan semua informasi tentang biaya perjalanan:
- Tujuan: Kota tujuan (Bali, Bekasi, Bogor, Jakarta, Solo)
- Kelas: Kelas perjalanan (Eksekutif, Premium, Bisnis, Ekonomi)
- Awal: Biaya awal perjalanan
- Tambahan: Biaya tambahan (5% dari biaya awal)
- Asuransi: Biaya asuransi (2% dari biaya awal + Rp4.500)
- Admin: Biaya administrasi (Rp2.500)
- Total: Total semua biaya

2. Fungsi-fungsi Utama
a. Menghitung Biaya (HitungBiaya)
Fungsi ini menghitung semua komponen biaya berdasarkan biaya awal yang dimasukkan:

- Biaya tambahan = 5% dari biaya awal
- Asuransi = 2% dari biaya awal + Rp4.500
- Biaya admin = Rp2.500
- Total = Semua biaya dijumlahkan

b. Menentukan Kelas (TentukanKelas)
Fungsi ini menentukan kelas perjalanan berdasarkan:
- Tujuan yang dipilih
- Biaya awal yang dimasukkan

Setiap tujuan memiliki standar biaya berbeda untuk setiap kelasnya.

c. Menambah Data (TambahData)
Fungsi untuk menambahkan data perjalanan baru:
- Memasukkan biaya awal
- Memilih tujuan
- Sistem akan menghitung semua biaya dan menentukan kelas
- Jika biaya cukup, data akan disimpan

d. Menampilkan Data
- TampilkanData: Menampilkan semua data yang tersimpan
- TampilkanSatuData: Menampilkan detail satu data perjalanan

3. Pencarian Data
- Sequential Search: Mencari data berdasarkan tujuan (pencarian satu per satu)
- Binary Search: Mencari data berdasarkan biaya awal (pencarian cepat dengan syarat data sudah terurut)

4. Pengurutan Data
- Selection Sort: Mengurutkan data berdasarkan total biaya dari terkecil ke terbesar
- Insertion Sort: Mengurutkan data berdasarkan kelas dari tertinggi (Eksekutif) ke terendah (Ekonomi)

5. Menu Utama
Aplikasi memiliki menu dengan pilihan:
- 1. Tambah Data Perjalanan
- 2. Lihat Semua Data
- 3. Cari Data Berdasarkan Tujuan
- 4. Cari Data Berdasarkan Biaya Awal
- 5. Urutkan Data Berdasarkan Total Biaya
- 6. Urutkan Data Berdasarkan Kelas
- 7. Keluar dari Aplikasi
*/

package main
import "fmt"

const NMAX = 100

type Biaya struct {
	Tujuan   string
	Kelas    string
	Awal     float64
	Tambahan float64
	Asuransi float64
	Admin    float64
	Total    float64
}

type daftarBiaya [NMAX]Biaya

// untuk menghitung rincian biaya
func HitungBiaya(biayaAwal float64) Biaya {
	var b Biaya
	b.Awal = biayaAwal
	b.Tambahan = biayaAwal * 0.05
	b.Asuransi = biayaAwal*0.02 + 4500
	b.Admin = 2500
	b.Total = b.Awal + b.Tambahan + b.Asuransi + b.Admin
	return b
}

// untuk menentukan kelas berdasarkan tujuan dan biaya awal
func TentukanKelas(b *Biaya) {
	if b.Tujuan == "bali" {
		if b.Awal >= 3000000 {
			b.Kelas = "eksekutif"
		} else if b.Awal >= 2500000 {
			b.Kelas = "premium"
		} else if b.Awal >= 2000000 {
			b.Kelas = "bisnis"
		} else if b.Awal >= 1500000 {
			b.Kelas = "ekonomi"
		} else {
			b.Kelas = "tidak cukup"
		}
	} else if b.Tujuan == "bekasi" {
		if b.Awal >= 800000 {
			b.Kelas = "eksekutif"
		} else if b.Awal >= 600000 {
			b.Kelas = "premium"
		} else if b.Awal >= 400000 {
			b.Kelas = "bisnis"
		} else if b.Awal >= 200000 {
			b.Kelas = "ekonomi"
		} else {
			b.Kelas = "tidak cukup"
		}
	} else if b.Tujuan == "bogor" {
		if b.Awal >= 900000 {
			b.Kelas = "eksekutif"
		} else if b.Awal >= 700000 {
			b.Kelas = "premium"
		} else if b.Awal >= 500000 {
			b.Kelas = "bisnis"
		} else if b.Awal >= 300000 {
			b.Kelas = "ekonomi"
		} else {
			b.Kelas = "tidak cukup"
		}
	} else if b.Tujuan == "jakarta" {
		if b.Awal >= 850000 {
			b.Kelas = "eksekutif"
		} else if b.Awal >= 650000 {
			b.Kelas = "premium"
		} else if b.Awal >= 450000 {
			b.Kelas = "bisnis"
		} else if b.Awal >= 250000 {
			b.Kelas = "ekonomi"
		} else {
			b.Kelas = "tidak cukup"
		}
	} else if b.Tujuan == "solo" {
		if b.Awal >= 2000000 {
			b.Kelas = "eksekutif"
		} else if b.Awal >= 1500000 {
			b.Kelas = "premium"
		} else if b.Awal >= 1000000 {
			b.Kelas = "bisnis"
		} else if b.Awal >= 700000 {
			b.Kelas = "ekonomi"
		} else {
			b.Kelas = "tidak cukup"
		}
	} else {
		b.Kelas = "tidak cukup"
	}
}

// untuk menambahkan data
func TambahData(b *daftarBiaya, n *int) {
	var tujuan string
	if *n < NMAX {
		fmt.Print("Masukkan biaya awal: Rp ")
		fmt.Scanln(&b[*n].Awal)

		fmt.Print("Masukkan tujuan (bali/bekasi/bogor/jakarta/solo): ")
		fmt.Scanln(&tujuan)

		var hasil Biaya
		hasil = HitungBiaya(b[*n].Awal)
		hasil.Tujuan = tujuan
		TentukanKelas(&hasil)

		if hasil.Kelas == "tidak cukup" {
			fmt.Println("Biaya tidak mencukupi untuk melakukan perjalanan ke tujuan tersebut.")
			return
		}

		b[*n] = hasil
		*n = *n + 1
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Data sudah penuh.")
	}
}

// untuk menampilkan seluruh data
func TampilkanData(data []Biaya, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data yang tersimpan.")
		return
	}
	for i := 0; i < jumlah; i++ {
		fmt.Printf("\nData ke-%d:\n", i+1)
		fmt.Printf("Tujuan           : %s\n", data[i].Tujuan)
		fmt.Printf("Kelas            : %s\n", data[i].Kelas)
		fmt.Printf("Biaya Awal       : Rp %.2f\n", data[i].Awal)
		fmt.Printf("Biaya Tambahan   : Rp %.2f\n", data[i].Tambahan)
		fmt.Printf("Biaya Asuransi   : Rp %.2f\n", data[i].Asuransi)
		fmt.Printf("Biaya Admin      : Rp %.2f\n", data[i].Admin)
		fmt.Printf("Total Biaya      : Rp %.2f\n", data[i].Total)
	}
}

// untuk menampilkan 1 data
func TampilkanSatuData(b Biaya) {
	fmt.Printf("\nTujuan           : %s\n", b.Tujuan)
	fmt.Printf("Kelas            : %s\n", b.Kelas)
	fmt.Printf("Biaya Awal       : Rp %.2f\n", b.Awal)
	fmt.Printf("Biaya Tambahan   : Rp %.2f\n", b.Tambahan)
	fmt.Printf("Biaya Asuransi   : Rp %.2f\n", b.Asuransi)
	fmt.Printf("Biaya Admin      : Rp %.2f\n", b.Admin)
	fmt.Printf("Total Biaya      : Rp %.2f\n", b.Total)
}

// Sequential Search berdasarkan tujuan
func SequentialSearch(data daftarBiaya, jumlah int, tujuan string) int {
	var i int
	for i = 0; i < jumlah; i++ {
		if data[i].Tujuan == tujuan {
			return i
		}
	}
	return -1
}

// Binary Search berdasarkan biaya awal
func BinarySearch(data daftarBiaya, jumlah int, target float64) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = jumlah - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].Awal == target {
			return tengah
		} else if target < data[tengah].Awal {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

// Sorting total biaya ascending dengan selection sort
func SelectionSortTotal(b *daftarBiaya, n int) {
	var i, j, min int
	var temp Biaya
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if b[j].Total < b[min].Total {
				min = j
			}
		}
		temp = b[i]
		b[i] = b[min]
		b[min] = temp
	}
}

// Sorting class descending dengan insertion sort
func InsertionSortKelas(b *daftarBiaya, n int) {
	var i, j int
	var key Biaya
	for i = 1; i < n; i++ {
		key = b[i]
		j = i - 1
		for j >= 0 && Prioritas(b[j].Kelas) < Prioritas(key.Kelas) {
			b[j+1] = b[j]
			j = j - 1
		}
		b[j+1] = key
	}
}

func Prioritas(kelas string) int {
	if kelas == "eksekutif" {
		return 4
	} else if kelas == "premium" {
		return 3
	} else if kelas == "bisnis" {
		return 2
	} else if kelas == "ekonomi" {
		return 1
	}
	return 0
}


func main() {
	var data daftarBiaya
	var jumlah int
	var pilih int
	for {
		fmt.Println("\n=== MENU APLIKASI BUDGET TRAVELING ===")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Cari Data (Tujuan)")
		fmt.Println("4. Cari Data (Biaya Awal)")
		fmt.Println("5. Urutkan Data (Total Biaya)")
		fmt.Println("6. Urutkan Data (Kelas)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			TambahData(&data, &jumlah)
		case 2:
			TampilkanData(data[:], jumlah)
		case 3:
			var tujuan string
			fmt.Print("Masukkan tujuan yang dicari (bali/bekasi/bogor/jakarta/solo): ")
			fmt.Scanln(&tujuan)
			var idx int
			idx = SequentialSearch(data, jumlah, tujuan)
			if idx != -1 {
				fmt.Println("Data ditemukan:")
				TampilkanSatuData(data[idx])
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 4:
			var awal float64
			fmt.Print("Masukkan biaya awal yang dicari: Rp ")
			fmt.Scanln(&awal)
			SelectionSortTotal(&data, jumlah)
			var idx int
			idx = BinarySearch(data, jumlah, awal)
			if idx != -1 {
				fmt.Println("Data ditemukan:")
				TampilkanSatuData(data[idx])
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 5:
			SelectionSortTotal(&data, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan total biaya.")
		case 6:
			InsertionSortKelas(&data, jumlah)
			fmt.Println("Data berhasil diurutkan berdasarkan kelas.")
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
