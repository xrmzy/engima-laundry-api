package entity

import (
	"bufio"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//genarete random ID
func GenerateRandomNumber(max int) int {
	return rand.Intn(max)
}


//validasi email
func IsEmailExists(db *sql.DB, email string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customers WHERE email = $1)", email)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

//generate customer ID 
func GenerateCustomerID() string {
	randomNumber := GenerateRandomNumber(1000)
	codePrefix := "CS"
	return fmt.Sprintf("%s%03d", codePrefix, randomNumber)
}

//read input
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}


func IsCustomerExists(db *sql.DB, customerID string) bool {
	var exists bool
	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customers WHERE cust_id = $1)", customerID)
	err := row.Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

