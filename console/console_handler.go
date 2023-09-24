package console

import (
	"enigma-laundry-console-api/config"
	"enigma-laundry-console-api/entity"
	"fmt"
	"os"
)

func ConnectConsole() {
	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		fmt.Println("=== Enigma Laundry ===")
		fmt.Println("1. Tambah Pelanggan")
		fmt.Println("2. Ubah Data Pelanggan")
		fmt.Println("3. Hapus Data Pelanggan")
		fmt.Println("4. Tambah Pesanan")
		fmt.Println("5. Ubah Data Pesanan")
		fmt.Println("6. Hapus Data Pesanan")
		fmt.Println("7. Cari Pelanggan")
		fmt.Println("8. Cari Pesanan dan Status")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih tindakan (1/2/3/4/5/6/7/8/9): ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			AddCustomer(entity.Customer{})
		case "2":
			UpdateCustomer(entity.Customer{})
		case "3":
			DeleteCustomer(entity.Customer{})
		case "4":
			// addOrder(entity.Orders{})
		case "5":
			// updateOrder(entity.Orders{})
		case "6":
			// deleteOrder(entity.Orders{})
		case "7":
		searchMenu:
			for {
				fmt.Println("1. Cari berdasarkan ID")
				fmt.Println("2. Cari semua data")
				fmt.Println("0. Kembali ke menu utama")
				fmt.Print("Pilih tindakan (1/2/0) :")

				var searchChoice string
				fmt.Scanln(&searchChoice)
				switch searchChoice {
				case "1":
					// searchBy(entity.Customer{})
				case "2":
					// searchAll(entity.Customer{})
				case "0":
					break searchMenu // balik ke menu utama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}
		case "8":
		searchMenuTx:
			for {
				fmt.Println("1. Cari berdasarkan ID Transaksi")
				fmt.Println("2. Cari semua data")
				fmt.Println("0. Kembali ke menu utama")
				fmt.Print("Pilih tindakan (1/2/0) :")

				var searchChoice string
				fmt.Scanln(&searchChoice)
				switch searchChoice {
				case "1":
					// searchOrderBy(entity.Orders{})
				case "2":
					// SearchAllOrders(entity.Orders{})
				case "0":
					break searchMenuTx // balik ke menu utama
				default:
					fmt.Println("Pilihan tidak valid.")
				}
			}
		case "9":
			fmt.Println("Anda keluar dari aplikasi !")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

}
