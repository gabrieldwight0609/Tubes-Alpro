package main

import "fmt"

const NMAX int = 100

type info struct {
	namaAset             string
	volumeBeli           float64
	hargaBeli, hargaJual float64
	hargaSekarang        float64
	Nama, jenisAsset     string
	n                    int
}

type tabInfo [NMAX]info

func tampilkanMenu() {
	fmt.Println("\n=== Aplikasi Manajemen Investasi ===")
	fmt.Println("1. Kalkulator Take Profit")
	fmt.Println("2. Top Aset")
	fmt.Println("3. Laporan PNL")
	fmt.Println("4. Wallet asset")
	fmt.Println("5. Keluar")
}

func fiturBiodata(A *info, B *tabInfo) {
	var pekerjaan string
	var umur int
	var jawaban string

	fmt.Println("\nSelamat datang di Aplikasi Manajemen Investasi!")
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&A.Nama)
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scanln(&umur)
	fmt.Print("Masukkan pekerjaan Anda: ")
	fmt.Scanln(&pekerjaan)

	fmt.Printf("\nHalo %s, %d tahun, seorang %s. Selamat menggunakan aplikasi!\n", A.Nama, umur, pekerjaan)

	fmt.Println("Apakah anda sudah memiliki aset?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("Jawab dengan kata ya atau tidak")
	fmt.Scanln(&jawaban)

	if jawaban == "Tidak" || jawaban == "tidak" {
		fmt.Println("Anda dapat mencari asset terlebih dahulu, lalu masukkan ke dalam sini!!")
		fmt.Println("Silahkan melanjutkan aplikasinya!")
	}
	if jawaban == "Ya" || jawaban == "ya" {
		fmt.Println("Masukkan banyaknya asset yang kamu punya")
		fmt.Scanln(&A.n)
		fmt.Println("Masukkan nama aset dan juga jenisnya")

		for i := 0; i < A.n; i++ {
			fmt.Printf("\nData aset ke-%d:\n", i+1)
			fmt.Print("Nama aset: ")
			fmt.Scanln(&B[i].namaAset)
			fmt.Print("Jenis aset: ")
			fmt.Scanln(&B[i].jenisAsset)
			fmt.Print("Volume beli: ")
			fmt.Scanln(&B[i].volumeBeli)
			fmt.Print("Harga beli: ")
			fmt.Scanln(&B[i].hargaBeli)
			fmt.Print("Harga jual (jika sudah dijual, jika belum masukkan 0): ")
			fmt.Scanln(&B[i].hargaJual)
		}

		fmt.Println("\n Asset yang kamu punya:")
		for i := 0; i < A.n; i++ {
			fmt.Printf("%d. %s (%s) - Volume: %.2f, Harga Beli: %.2f, Harga Jual: %.2f\n",
				i+1, B[i].namaAset, B[i].jenisAsset, B[i].volumeBeli, B[i].hargaBeli, B[i].hargaJual)
		}
	}
}

// menghitung keuntungan atau kerugian dari target yang diinginkan, serta menampilkan apakah aset tersebut sudah take profit atau belum
func fiturKalkulator(A info) {
	var target, keuntungan float64

	fmt.Println("\n=== Kalkulator Take Profit ===")
	fmt.Print("Nama aset: ")
	fmt.Scanln(&A.namaAset)
	fmt.Print("Harga beli: ")
	fmt.Scanln(&A.hargaBeli)
	fmt.Print("Harga saat ini: ")
	fmt.Scanln(&A.hargaSekarang)
	fmt.Print("Target keuntungan (dalam persen): ")
	fmt.Scanln(&target)

	keuntungan = ((A.hargaSekarang - A.hargaBeli) / A.hargaBeli) * 100
	if keuntungan >= target {
		fmt.Printf("\n Aset %s layak untuk take profit! (Keuntungan: %.2f%%)\n", A.namaAset, keuntungan)
	} else {
		fmt.Printf("\n Aset %s belum layak take profit. (Keuntungan: %.2f%%)\n", A.namaAset, keuntungan)
	}

}

// untuk menyortir jenis aset yang paling banyak atau paling sedikit dimiliki oleh pengguna atau user
func fiturJumlahInvestasi(A tabInfo) {
	var jumlahData int

	fmt.Println("\n=== Top Aset ===")
	fmt.Print("Masukkan jumlah aset yang ingin diinput: ")
	fmt.Scanln(&jumlahData)

	i := 0
	for i < jumlahData {
		fmt.Printf("\nNama aset ke-%d: ", i+1)
		fmt.Scanln(&A[i].namaAset)
		fmt.Printf("Volume beli %s: ", A[i].namaAset)
		fmt.Scanln(&A[i].volumeBeli)
		i++
	}

	i = 0
	for i < jumlahData-1 {
		j := 0
		for j < jumlahData-i-1 {
			if A[j].volumeBeli < A[j+1].volumeBeli {
				A[j].namaAset, A[j+1].namaAset = A[j+1].namaAset, A[j].namaAset
				A[j].volumeBeli, A[j+1].volumeBeli = A[j+1].volumeBeli, A[j].volumeBeli
			}
			j++
		}
		i++
	}

	fmt.Println("\n Daftar Aset berdasarkan Top Volume + Analisis:")
	i = 0
	for i < jumlahData {
		fmt.Printf("%d. %s\n", i+1, A[i].namaAset)
		fmt.Printf("   Total Volume: %.2f\n", A[i].volumeBeli)
		i++
	}
}

// untuk melihat keuntungan dan kerugian yang terjadi dalam jangka waktu tertentu
func fiturPNL(B tabInfo, A info) {
	fmt.Println("\n=== Laporan Profit and Loss (PNL) ===")
	if A.n == 0 {
		fmt.Println("Tidak ada aset yang dimiliki.")
		return
	}

	for i := 0; i < A.n; i++ {
		var pnlNominal, pnlPersen float64
		pnlNominal = (B[i].hargaSekarang - B[i].hargaBeli) * B[i].volumeBeli

		if B[i].hargaBeli != 0 {
			pnlPersen = ((B[i].hargaSekarang - B[i].hargaBeli) / B[i].hargaBeli) * 100
		}

		sign := ""
		if pnlNominal >= 0 {
			sign = "+"
		}

		fmt.Printf("%d. %s\n", i+1, B[i].namaAset)
		fmt.Printf("   Volume: %.2f\n", B[i].volumeBeli)
		fmt.Printf("   Harga Beli: %.2f\n", B[i].hargaBeli)
		fmt.Printf("   Harga Sekarang: %.2f\n", B[i].hargaSekarang)
		fmt.Printf("   PNL: %s %.2f %s %.2f\n", sign, pnlNominal, sign, pnlPersen)
	}
}

// Sequential Search berdasarkan nama aset
func seqSearchNama(T tabInfo, n int, X string) int {
	found := -1
	j := 0
	for j < n && found == -1 {
		if T[j].namaAset == X {
			found = j
		}
		j++
	}
	return found
}

// Sequential Search berdasarkan jenis aset
func seqSearchJenis(T tabInfo, n int, X string) int {
	found := -1
	j := 0
	for j < n && found == -1 {
		if T[j].jenisAsset == X {
			found = j
		}
		j++
	}
	return found
}

// Binary Search ascending berdasarkan nama aset
func binarySearchNama(T tabInfo, n int, X string) int {
	found := -1
	kr := 0
	kn := n - 1
	var med int
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X < T[med].namaAset {
			kn = med - 1
		} else if X > T[med].namaAset {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

// Selection Sort ascending berdasarkan nama aset
func selectionSortNama(T *tabInfo, n int) {
	i := 1
	for i <= n-1 {
		idx_min := i - 1
		j := i
		for j < n {
			if T[j].namaAset < T[idx_min].namaAset {
				idx_min = j
			}
			j++
		}
		t := T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i++
	}
}

// Insertion Sort descending berdasarkan volume beli
func insertionSortVolume(T *tabInfo, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		for j > 0 && temp.volumeBeli > T[j-1].volumeBeli {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

// menambahkan jumlah asset, mengedit jumlah asset dan menghapus data dari daftar asset tersebut (sesuai apa yang mau di lakukan oleh pengguna/user)
func walletAsset(A *info, B *tabInfo) {
	var pilihan int
	fmt.Println("\n== Wallet Asset ==")

	fmt.Println("Selamat datang di walet asset! Silahkan pilih opsi")

	fmt.Println("\nPilih opsi:")
	fmt.Println("1. Tambahkan data")
	fmt.Println("2. Hapus data")
	fmt.Println("3. Edit data")
	fmt.Println("4. Cari data")
	fmt.Println("5. Urutkan data")
	fmt.Println("6. Aset sekarang")
	fmt.Println("7. Keluar")

	fmt.Print("Pilihan Anda (1-6): ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		if A.n >= NMAX {
			fmt.Println("Data aset sudah penuh, tidak bisa menambah lagi.")
			return
		}
		var asetBaru info
		fmt.Print("Masukkan nama aset: ")
		fmt.Scanln(&asetBaru.namaAset)
		fmt.Print("Masukkan jenis aset: ")
		fmt.Scanln(&asetBaru.jenisAsset)
		fmt.Print("Masukkan volume beli: ")
		fmt.Scanln(&asetBaru.volumeBeli)
		fmt.Print("Masukkan harga beli: ")
		fmt.Scanln(&asetBaru.hargaBeli)
		fmt.Print("Masukkan harga jual: ")
		fmt.Scanln(&asetBaru.hargaJual)
		fmt.Print("Masukkan harga sekarang: ")
		fmt.Scanln(&asetBaru.hargaSekarang)

		B[A.n] = asetBaru
		A.n = A.n + 1
		fmt.Println("Data aset berhasil ditambahkan.")
	} else if pilihan == 2 {
		var namaCari string
		fmt.Print("Masukkan nama aset yang ingin dihapus: ")
		fmt.Scanln(&namaCari)
		idx := seqSearchNama(*B, A.n, namaCari)
		if idx == -1 {
			fmt.Println("Aset tidak ditemukan.")
			return
		}
		for i := idx; i < A.n-1; i++ {
			B[i] = B[i+1]
		}
		A.n = A.n - 1
		fmt.Println("Data aset berhasil dihapus.")

	} else if pilihan == 3 {
		var namaEdit string
		fmt.Print("Masukkan nama aset yang ingin diedit: ")
		fmt.Scanln(&namaEdit)

		idx := seqSearchNama(*B, A.n, namaEdit)

		if idx == -1 {
			fmt.Println("Aset tidak ditemukan.")
			return
		}

		fmt.Println("Data aset ditemukan. Silakan masukkan data baru:")
		fmt.Print("Masukkan nama aset baru: ")
		fmt.Scanln(&B[idx].namaAset)
		fmt.Print("Masukkan jenis aset baru: ")
		fmt.Scanln(&B[idx].jenisAsset)
		fmt.Print("Masukkan volume beli baru: ")
		fmt.Scanln(&B[idx].volumeBeli)
		fmt.Print("Masukkan harga beli baru: ")
		fmt.Scanln(&B[idx].hargaBeli)
		fmt.Print("Masukkan harga jual baru: ")
		fmt.Scanln(&B[idx].hargaJual)
		fmt.Print("Masukkan harga sekarang baru: ")
		fmt.Scanln(&B[idx].hargaSekarang)
		fmt.Println("Data aset berhasil diedit.")
	} else if pilihan == 4 {
		var metodeCari int
		fmt.Println("Pilih metode pencarian:")
		fmt.Println("1. Sequential Search berdasarkan nama aset")
		fmt.Println("2. Sequential Search berdasarkan jenis aset")
		fmt.Println("3. Binary Search berdasarkan nama aset (data harus sudah terurut ascending berdasarkan nama aset)")
		fmt.Print("Pilihan Anda (1-3): ")
		fmt.Scanln(&metodeCari)

		if metodeCari == 1 {
			var namaCari string
			fmt.Print("Masukkan nama aset yang dicari: ")
			fmt.Scanln(&namaCari)
			idx := seqSearchNama(*B, A.n, namaCari)
			if idx == -1 {
				fmt.Println("Aset tidak ditemukan.")
			} else {
				fmt.Printf("Aset ditemukan di indeks ke-%d: %s (%s)\n", idx+1, B[idx].namaAset, B[idx].jenisAsset)
			}
		} else if metodeCari == 2 {
			var jenisCari string
			fmt.Print("Masukkan jenis aset yang dicari: ")
			fmt.Scanln(&jenisCari)
			idx := seqSearchJenis(*B, A.n, jenisCari)
			if idx == -1 {
				fmt.Println("Aset tidak ditemukan.")
			} else {
				fmt.Printf("Aset ditemukan di indeks ke-%d: %s (%s)\n", idx+1, B[idx].namaAset, B[idx].jenisAsset)
			}
		} else if metodeCari == 3 {
			var namaCari string
			fmt.Print("Masukkan nama aset yang dicari: ")
			fmt.Scanln(&namaCari)

			selectionSortNama(B, A.n)
			idx := binarySearchNama(*B, A.n, namaCari)
			if idx == -1 {
				fmt.Println("Aset tidak ditemukan.")
			} else {
				fmt.Printf("Aset ditemukan di indeks ke-%d: %s (%s)\n", idx+1, B[idx].namaAset, B[idx].jenisAsset)
			}
		} else {
			fmt.Println("Pilihan metode pencarian tidak valid.")
		}
	} else if pilihan == 5 {
		var metodeUrut int
		fmt.Println("Pilih metode pengurutan:")
		fmt.Println("1. Selection Sort ascending berdasarkan nama aset")
		fmt.Println("2. Insertion Sort descending berdasarkan volume beli")
		fmt.Print("Pilihan Anda (1-2): ")
		fmt.Scanln(&metodeUrut)

		if metodeUrut == 1 {
			selectionSortNama(B, A.n)
			fmt.Println("Data aset sudah diurutkan ascending berdasarkan nama aset.")
		} else if metodeUrut == 2 {
			insertionSortVolume(B, A.n)
			fmt.Println("Data aset sudah diurutkan descending berdasarkan volume beli.")
		} else {
			fmt.Println("Pilihan metode pengurutan tidak valid.")
		}
	} else if pilihan == 6 {
		if A.n == 0 {
			fmt.Println("Anda belum memiliki aset, silahkan berinvestasi terlebih dahulu agar anda memiliki aset")
			fmt.Println("Anda dapat menggunakan fitur kami untuk membantu anda dalam mengambil keputusan untuk berinvestasi")
			return
		}
		fmt.Println("Daftar aset yang anda miliki")
		for i := 0; i < A.n; i++ {
			fmt.Printf("%d. %s (%s) - Volume: %.2f\n", i+1, B[i].namaAset, B[i].jenisAsset, B[i].volumeBeli)
		}
	} else if pilihan == 7 {
		fmt.Println("Terima kasih sudah menggunakan aplikasi")

	} else {
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

func main() {
	var a tabInfo
	var b info

	fiturBiodata(&b, &a)

	var pilihan int
	for pilihan != 5 {
		tampilkanMenu()
		fmt.Print("\nPilih menu (1-5): ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fiturKalkulator(b)
		}
		if pilihan == 2 {
			fiturJumlahInvestasi(a)
		}
		if pilihan == 3 {
			fiturPNL(a, b)
		}
		if pilihan == 4 {
			walletAsset(&b, &a)
		}
		if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		}
	}
}
