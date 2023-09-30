package console

import (
	"enigma-laundry-console-api/config"
	"enigma-laundry-console-api/entity"
	"fmt"
	"log"
)

func AddOrder(order entity.Orders) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== TAMBAH TRANSAKSI ===")
	order.OrderId = entity.GenerateOrderID()
	order.CustomerName = entity.ReadInput("Masukkan Nama: ")

	err = entity.FindCustomerID(db, &order)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	order.Service = entity.GetUserMenuChoice(&order)
	if order.Service == "" {
		return
	}
	order.Unit = entity.GetUnitChoice(&order)
	if order.Unit == "" {
		return
	}
	order.OutletName = entity.GetOutletChoice(&order)
	if order.OutletName == "" {
		return
	}
	order.OrderDate = entity.GetOrderDate(&order)
	if order.OrderDate == "" {
		return
	}
	order.Status = entity.GetStatusChoice(&order)
	if order.Status == "" {
		return
	}

	//Tx Begin
	tx, err := config.BeginTrx(db)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v\n", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	entity.AddOrderSQL(tx, &order)

}

func UpdateOrder(order entity.Orders) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== UBAH DATA TRANSAKSI ===")
	order.OrderId = entity.ReadInput("Masukkan ID Transaksi: ")

	customerID, err := entity.ValidateIdTX(db, &order)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Customer ID: %s\n", customerID)
	}

	order.CustomerName = entity.ReadInput("Masukkan Nama: ")
	err = entity.FindCustomerID(db, &order)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	order.Service = entity.GetUserMenuChoice(&order)
	if order.Service == "" {
		return
	}
	order.Unit = entity.GetUnitChoice(&order)
	if order.Unit == "" {
		return
	}
	order.OutletName = entity.GetOutletChoice(&order)
	if order.OutletName == "" {
		return
	}
	order.OrderDate = entity.GetOrderDate(&order)
	if order.OrderDate == "" {
		return
	}
	order.Status = entity.GetStatusChoice(&order)
	if order.Status == "" {
		return
	}

	//Tx Begin
	tx, err := config.BeginTrx(db)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v\n", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	entity.UpdateOrderSQL(tx, &order)

}

func DeleteOrder(order entity.Orders) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== HAPUS DATA TRANSAKSI ===")
	order.OrderId = entity.ReadInput("Masukkan ID Transaksi: ")

	// Validasi ID Transaksi
	isValid, err := entity.IsOrderIDValid(db, order.OrderId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if !isValid {
		fmt.Println("ID Transaksi tidak ditemukan!")
		return
	}

	//Tx Begin
	tx, err := config.BeginTrx(db)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v\n", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	entity.DeleteOrderSQL(tx, &order)
	if err != nil {
		fmt.Printf("Failed to Delete Order: %v\n", err)
		return
	}

	fmt.Println("Successfully Delete Order!")
}

// search Order by ID
func SearchOrderBy(order entity.Orders) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== CARI TRANSAKSI BERDASARKAN ID ===")
	order.OrderId = entity.ReadInput("Masukkan ID Transaksi: ")

	if !entity.IsOrderExist(db, order.OrderId) {
		fmt.Println("Transaksi tidak ditemukan!")
		return
	}

	if err := entity.SearchOrderBySQL(db, &order); err != nil {
		panic(err)
	}
}

// search All Order
func SearchAllOrders(order entity.Orders) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== CARI SEMUA PELANGGAN ===")
	rows, err := db.Query("SELECT order_id, cust_id, cust_name, service, unit, outlet_name, order_date, status FROM orders;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&order.OrderId, &order.CustomerId, &order.CustomerName, &order.Service, &order.Unit, &order.OutletName, &order.OrderDate, &order.Status)
		if err != nil {
			fmt.Println("Tidak ada data:", err)
			continue
		} else {
			fmt.Println("Order ID :", order.OrderId)
			fmt.Println("Customer ID :", order.CustomerId)
			fmt.Println("Nama Customer :", order.CustomerName)
			fmt.Println("Service :", order.Service)
			fmt.Println("Unit :", order.Unit)
			fmt.Println("Outlet :", order.OutletName)
			fmt.Println("Tanggal Order:", order.OrderDate)
			fmt.Println("Status :", order.Status)
			fmt.Println("============================")
		}
	}
}
