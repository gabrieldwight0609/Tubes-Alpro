package main

import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

type info struct {
	namaAset             string
	volumeBeli           float64
	hargaBeli, hargaJual float64
	hargaSekarang        float64
}

type tabInfo [NMAX]info

func tampilkanMenu() {
	fmt.Println("\n=== Aplikasi Manajemen Investasi ===")
	fmt.Println("1. Kalkulator Take Profit")
	fmt.Println("2. Top Aset")
	fmt.Println("3. Laporan PNL")
	fmt.Println("4. Wallet asset.", "wallet asset adalah fitur untuk melihat laporan portofolio")
	fmt.Println("5. Keluar")
}

func fiturBiodata() {
	var nama, pekerjaan string
	var umur int

	fmt.Println("\nSelamat datang di Aplikasi Manajemen Investasi!")
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scanln(&umur)
	fmt.Print("Masukkan pekerjaan Anda: ")
	fmt.Scanln(&pekerjaan)

	fmt.Printf("\nHalo %s, %d tahun, seorang %s. Selamat menggunakan aplikasi!\n", nama, umur, pekerjaan)
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
		fmt.Printf("   Total Volume: Beli: %d\n", A[i].volumeBeli)
		i++
	}
}

// untuk melihat keuntungan dan kerugian yang terjadi dalam jangka waktu tertentu
func fiturPNL() {
	fmt.Print("masih bingung bang")
}

// menambahkan jumlah asset, mengedit jumlah asset dan menghapus data dari daftar asset tersebut (sesuai apa yang mau di lakukan oleh pengguna/user)
func walletAsset() {
	fmt.Print("masih bingung bang")

}

// untuk mencari asset berdasarkan nama maupun jenis dari assetnya, tetapi lebih baik untuk fokus dalam satu asset.
func daftarAsset() {
	fmt.Print("masih bingung bang")
}

func main() {
	var a tabInfo
	var b info

	fiturBiodata()

	var pilihan int
	for pilihan != 5 {
		tampilkanMenu()
		fmt.Print("\nPilih menu (1-4): ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fiturKalkulator(b)
		}
		if pilihan == 2 {
			fiturJumlahInvestasi(a)
		}
		if pilihan == 3 {
			fiturPNL()
		}
		if pilihan == 4 {
			walletAsset()
		}
		if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		}
	}

}
