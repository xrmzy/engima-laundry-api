package entity

import (
	"database/sql"
	"fmt"
	"log"
)



func InsertCustomerSQL(db *sql.DB, customer Customer) {
	sqlStatement := "INSERT INTO customers (cust_id, cust_name, phone_number, address, email) VALUES ($1, $2, $3, $4, $5);"
	_, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	if err != nil {
		log.Fatalf("Failed to Insert Customer data: %v", err)
	}
}


func UpdateCustomerSQL(db *sql.DB, customer Customer) {
	sqlStatement := "UPDATE customers SET cust_name = $2, phone_number = $3, address = $4, email = $5 WHERE cust_id = $1;"
	_, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.PhoneNumber, customer.Address, customer.Email)
	if err != nil {
		log.Fatalf("Failed to Update Customer Data: %v", err)
	} 
}

func DeleteCustomerSQL(db *sql.DB, customer Customer) {
	sqlStatement := "DELETE FROM customers WHERE cust_id = $1;"
	_, err := db.Exec(sqlStatement, customer.Id)
	if err != nil {
		log.Fatalf("Failed to Delete Customer Data: %v", err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}