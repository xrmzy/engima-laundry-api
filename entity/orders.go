package entity

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Orders struct {
	OrderId      string
	CustomerId   string
	CustomerName string
	Service      string
	Unit         string
	OutletName   string
	OrderDate    string
	Status       string
}

func GenerateOrderID() string {
	randomNumber := GenerateRandomNumber(1000)
	codePrefix := "TX"
	return fmt.Sprintf("%s%03d", codePrefix, randomNumber)
}

func ValidateIdTX(db *sql.DB, order *Orders) (string, error) {
	var err error
	row := db.QueryRow("SELECT cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders WHERE order_id = $1;", order.OrderId)
	err = row.Scan(&order.CustomerId, &order.CustomerName, &order.Service, &order.Unit, &order.OutletName, &order.OrderDate, &order.Status)
	if err == sql.ErrNoRows {
		fmt.Println("Transaksi tidak ditemukan!")
		return "", nil
	} else if err != nil {
		return "", err
	}
	return order.CustomerId, nil
}

func GetUserMenuChoice(order *Orders) string {
	fmt.Println("Pilih Paket:")
	fmt.Println("1. PAKET A/BERSIH AMAN")
	fmt.Println("2. PAKET B/BERSIH TENANG")
	fmt.Println("3. PAKET C/LENGKAP LUAR BIASA")
	fmt.Print("Pilih paket (1/2/3): ")
	var menuChoice string
	fmt.Scanln(&menuChoice)

	switch menuChoice {
	case "1":
		return "PAKET A/BERSIH AMAN"
	case "2":
		return "PAKET B/BERSIH TENANG"
	case "3":
		return "PAKET C/LENGKAP LUAR BIASA"
	default:
		fmt.Println("Pilihan Anda Tidak Ada")
		return ""
	}
}

func GetUnitChoice(order *Orders) string {
	fmt.Print("Masukkan Jumlah Unit (per Kg) :")
	_, err := fmt.Scanln(&order.Unit)
	if err != nil {
		fmt.Println("Input tidak Valid")
		return ""
	}
	return order.Unit
}

func GetOutletChoice(order *Orders) string {
	fmt.Println("Pilih Outlet:")
	fmt.Println("1. LAUNDRY SENANG")
	fmt.Println("2. LAUNDRY BAHAGIA")
	fmt.Print("Pilih Outlet (1/2/): ")
	var outletChoice string
	fmt.Scanln(&outletChoice)

	switch outletChoice {
	case "1":
		return "LAUNDRY SENANG"
	case "2":
		return "LAUNDRY BAHAGIA"
	default:
		fmt.Println("Pilihan tidak valid")
	}
	return order.OutletName
}

func GetOrderDate(order *Orders) string {
	fmt.Print("Masukkan Tanggal Order (YYYY-MM-DD): ")

	var orderDateStr string
	_, err := fmt.Scanln(&orderDateStr)
	if err != nil {
		fmt.Println("Input Tanggal Tidak Valid")
		return ""
	}

	// Validasi format tanggal
	layout := "2006-01-02"
	parseDate, err := time.Parse(layout, orderDateStr)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Harap Masukkan dengan format yang sudah ditentukan")
		return ""
	}
	// Set order.OrderDate ke hasil parsing yang valid
	order.OrderDate = parseDate.Format(layout)
	return order.OrderDate
}

func GetStatusChoice(order *Orders) string {
	fmt.Println("Pilih Status Order :")
	fmt.Println("1. Dalam Proses")
	fmt.Println("2. Selesai")
	fmt.Println("3. Dibatalkan")
	fmt.Print("Pilih status (1/2/3): ")
	var statusChoice string
	fmt.Scanln(&statusChoice)

	switch statusChoice {
	case "1":
		return "Proses"
	case "2":
		return "Done"
	case "3":
		return "Cancel"
	default:
		fmt.Println("Pilihan tidak ada")
	}
	return order.Status
}

func IsOrderIDValid(db *sql.DB, orderID string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM orders WHERE order_id = $1);", orderID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func IsOrderExist(db *sql.DB, orderID string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM orders WHERE order_id = $1)", orderID)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}
