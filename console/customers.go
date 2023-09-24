package console

import (
	"enigma-laundry-console-api/config"
	"enigma-laundry-console-api/entity"
	"fmt"
	"log"
)

func AddCustomer(customer entity.Customer) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== PENDAFTARAN CUSTOMER ===")
	customer.Id = entity.GenerateCustomerID()
	customer.Name = entity.ReadInput("Masukkan Nama : ")
	customer.Address = entity.ReadInput("Masukkan Alamat : ")
	customer.PhoneNumber = entity.ReadInput("Masukkan Nomor Telepon : ")
	customer.Email = entity.ReadInput("Masukkan Email : ")

	//validasi email
	if entity.IsEmailExists(db, customer.Email) {
		fmt.Println("Email sudah digunakan. Silahkan gunakan email lain!")
		return
	}

	entity.InsertCustomerSQL(db, customer)
	if err != nil {
		fmt.Printf("Failed Insert Data: %v\n", err)
	} else {
		fmt.Println("Succesfully Insert Data!")
	}
}


func UpdateCustomer(customer entity.Customer) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== UBAH DATA CUSTOMER ===")
	customer.Id = entity.ReadInput("Masukkan ID Customer: ")

	if !entity.IsCustomerExists(db, customer.Id){
		fmt.Println("Pelanggan Tidak Ditemukan!")
		return
	}

customer.Name = entity.ReadInput("Masukkan Nama Customer: ")
customer.Address = entity.ReadInput("Masukkan Alamat: ")
customer.PhoneNumber = entity.ReadInput("Masukkan Nomor Telepon: ")
customer.Email = entity.ReadInput("Masukkan Email: ")

	//validasi email
	if entity.IsEmailExists(db, customer.Email) {
		fmt.Println("Email sudah digunakan. Silahkan gunakan email lain!")
		return
	}

	entity.UpdateCustomerSQL(db, customer) 
	if err != nil {
		fmt.Printf("Failed Update Data: %v\n", err)
	} else {
		fmt.Println("Succesfully Update Data!")
		}
}


func DeleteCustomer(customer entity.Customer) {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Failed to Connect to the dabase: %v", err)
	}
	defer db.Close()

	fmt.Println("=== HAPUS DATA CUSTOMER ===")
	customer.Id = entity.ReadInput("Masukkan ID Customer: ")

	if !entity.IsCustomerExists(db, customer.Id) {
		fmt.Println("Pelanggan Tidak Ditemukan!")
		return
	}


	entity.DeleteCustomerSQL(db, customer)
	if err != nil {
		fmt.Printf("Failed Delete Data: %v\n", err)
	} else {
		fmt.Println("Succesfully Delete Data!")
	}
}