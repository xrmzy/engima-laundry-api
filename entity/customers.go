package entity

import "fmt"

type Customer struct {
	Id          string
	Name        string
	Address     string
	PhoneNumber string
	Email       string
}

// generate customer ID
func GenerateCustomerID() string {
	randomNumber := GenerateRandomNumber(1000)
	codePrefix := "CS"
	return fmt.Sprintf("%s%03d", codePrefix, randomNumber)
}
