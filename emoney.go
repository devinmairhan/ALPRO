package main

import (
	"fmt"
)

type Akun struct {
	ID        int
	Nama      string
	Saldo     float64
	Disetujui bool
}

type Transaksi struct {
	ID       int
	Dari     int
	Ke       int
	Jumlah   float64
	Jenis    string
}

const maxAkun = 100
const maxTransaksi = 100

var akunList [maxAkun]Akun
var transaksiList [maxTransaksi]Transaksi
var akunCount int
var transaksiCount int

func tambahAkun(akun Akun) {
	if akunCount < maxAkun {
		akunList[akunCount] = akun
		akunCount++
	} else {
		fmt.Println("Maksimum jumlah akun tercapai")
	}
}

func tambahTransaksi(transaksi Transaksi) {
	if transaksiCount < maxTransaksi {
		transaksiList[transaksiCount] = transaksi
		transaksiCount++
	} else {
		fmt.Println("Maksimum jumlah transaksi tercapai")
	}
}

func registrasiAkun(id int,nama string) {
	akun := Akun{ID : id,Nama : nama, Saldo : 0.0 , Disetujui : false}
	tambahAkun(akun)
	fmt.Printf("Akun Terdaftar: ID: %d, Nama: %s, Saldo: %.2f, Disetujui: %v\n", akun.ID, akun.Nama, akun.Saldo, akun.Disetujui)
	// fmt.Println("Akun Terdaftar",akun)
}

func setujuiAkun(id int) {
	for i := 0 ; i < akunCount ; i++ {
		if akunList[i].ID == id {
		akunList[i].Disetujui = true
		fmt.Printf("Akun dengan ID: %d dan Nama: %s telah disetujui.\n", akunList[i].ID, akunList[i].Nama)
		// fmt.Println("Akun Di Setujui:",akunList[i])
		return 
	}
}
fmt.Println("Akun tidak ditemukan")
}	

func menyetorUang(id int, jumlah float64) {
    for i := 0; i < akunCount; i++ {
        if akunList[i].ID == id {
            akunList[i].Saldo += jumlah

            transaksi := Transaksi{
                ID:      transaksiCount + 1,
                Dari:    0, // Setoran Tunai
                Ke:      id,
                Jumlah:  jumlah,
                Jenis:   "Setoran",
            }
            tambahTransaksi(transaksi)
			fmt.Printf("Setoran berhasil: Transaksi ke: %d, Dari ID: %d, Ke ID: %d, Jumlah: %.2f, Jenis: %s\n", transaksi.ID, transaksi.Dari, transaksi.Ke, transaksi.Jumlah, transaksi.Jenis)
            // fmt.Println("Setoran berhasil:", transaksi)
            return
        }
    }
	fmt.Println("Akun tidak ditemukan")
}			

func transferUang(dari, ke int, jumlah float64) {
    var akunDari, akunKe *Akun

    for i := 0; i < akunCount; i++ {
        if akunList[i].ID == dari {
            akunDari = &akunList[i]
        } else if akunList[i].ID == ke {
            akunKe = &akunList[i]
        }
    }

    if akunDari != nil && akunKe != nil && akunDari.Disetujui && akunKe.Disetujui {
        if akunDari.Saldo >= jumlah {
            akunDari.Saldo -= jumlah
            akunKe.Saldo += jumlah

            transaksi := Transaksi{
                ID:      transaksiCount + 1,
                Dari:    dari,
                Ke:      ke,
                Jumlah:  jumlah,
                Jenis:   "Transfer",
            }
            tambahTransaksi(transaksi)
			fmt.Printf("Transfer berhasil: Transaksi ke: %d, Dari ID: %d, Ke ID: %d, Jumlah: %.2f, Jenis: %s\n", transaksi.ID, transaksi.Dari, transaksi.Ke, transaksi.Jumlah, transaksi.Jenis)
            // fmt.Println("Transfer berhasil:", transaksi)
        } else {
            fmt.Println("Saldo tidak mencukupi")
        }
    } else {
        fmt.Println("Transfer gagal. Akun tidak disetujui atau tidak ditemukan.")
    }
}

func lakukanPembayaran(dari int, layanan string, jumlah float64) {
    var akunDari *Akun

    for i := 0; i < akunCount; i++ {
        if akunList[i].ID == dari {
            akunDari = &akunList[i]
        }
    }

    if akunDari != nil && akunDari.Disetujui {
        if akunDari.Saldo >= jumlah {
            akunDari.Saldo -= jumlah

            transaksi := Transaksi{
                ID:      transaksiCount + 1,
                Dari:    dari,
                Ke:      0,
                Jumlah:  jumlah,
                Jenis:   "Pembayaran" + layanan, //tambahan layanan buat identidikasi pembayaran apa
            }
            tambahTransaksi(transaksi)
			fmt.Printf("Pembayaran %s berhasil: Transaksi ke : %d, Dari ID: %d, Ke ID: %d, Jumlah: %.2f\n", layanan, transaksi.ID, transaksi.Dari, transaksi.Ke, transaksi.Jumlah)
            // fmt.Println("Pembayaran berhasil:", transaksi)
        } else {
            fmt.Println("Saldo tidak mencukupi")
        }
    } else {
        fmt.Println("Pembayaran gagal. Akun tidak disetujui atau tidak ditemukan.")
    }
}

func cetakAkun() {
	fmt.Println("Daftar Akun:")
	for i := 0; i < akunCount; i++ {
		status := "Belum Disetujui"
		if akunList[i].Disetujui {
			status = "Disetujui"
		}
		fmt.Printf("ID: %d, Nama: %s, Saldo: %.2f, Status: %s\n", akunList[i].ID, akunList[i].Nama, akunList[i].Saldo, status)
	}
}

func cetakTransaksi() {
	fmt.Println("Daftar Transaksi:")
	for i := 0; i < transaksiCount; i++ {
		fmt.Printf("Transaksi: %d, Dari ID: %d, Ke ID: %d, Jumlah: %.2f, Jenis: %s\n", transaksiList[i].ID, transaksiList[i].Dari, transaksiList[i].Ke, transaksiList[i].Jumlah, transaksiList[i].Jenis)
	}
}

func main() {
	for {
		var pilihan int
		fmt.Println("\n========================")
		fmt.Println("    Aplikasi E-Money    ")
		fmt.Println("========================")
		fmt.Println("1. Registrasi Akun")
		fmt.Println("2. Setujui Akun")
		fmt.Println("3. Menyetor Uang")
		fmt.Println("4. Transfer Uang")
		fmt.Println("5. Lakukan Pembayaran")
		fmt.Println("6. Cetak Akun")
		fmt.Println("7. Cetak Transaksi")
		fmt.Println("8. Keluar")
		fmt.Println("========================")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var id int
			var nama string
			fmt.Print("Masukkan ID Akun: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan Nama Akun: ")
			fmt.Scanln(&nama)
			registrasiAkun(id, nama)

		case 2:
			var id int
			fmt.Print("Masukkan ID Akun untuk Disetujui: ")
			fmt.Scanln(&id)
			setujuiAkun(id)

		case 3:
			var id int
			var jumlah float64
			fmt.Print("Masukkan ID Akun Anda: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan Jumlah untuk Disetor: ")
			fmt.Scanln(&jumlah)
			menyetorUang(id, jumlah)

		case 4:
			var dari, ke int
			var jumlah float64
			fmt.Print("Masukkan ID Akun Anda: ")
			fmt.Scanln(&dari)
			fmt.Print("Masukkan ID Akun Tujuan: ")
			fmt.Scanln(&ke)
			fmt.Print("Masukkan Jumlah untuk Ditransfer: ")
			fmt.Scanln(&jumlah)
			transferUang(dari, ke, jumlah)

		case 5:
			var dari int
			var layanan string
			var jumlah float64
			fmt.Print("Masukkan ID Akun Anda: ")
			fmt.Scanln(&dari)
			fmt.Print("Masukkan Nama Layanan: ")
			fmt.Scanln(&layanan)
			fmt.Print("Masukkan Jumlah untuk Pembayaran: ")
			fmt.Scanln(&jumlah)
			lakukanPembayaran(dari, layanan, jumlah)

		case 6:
			cetakAkun()

		case 7:
			cetakTransaksi()

		case 8:
			fmt.Println("Aplikasi telah log out")
			return


		}
	}
}
