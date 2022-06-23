package services

import (
	"github.com/kjasuquo/jumia_task/models"
	"log"
)

//GetAllValidatedNumbers compiles all the validated numbers to be inserted in the "number" table
func GetAllValidatedNumbers(customers []models.Customer) []models.Numbers {
	var data []models.Numbers
	for i := 0; i < len(customers); i++ {
		number, err := ValidatePhoneNumber(customers[i].Phone)
		if err != nil {
			log.Println(err)
		}
		data = append(data, number)
	}

	return data
}
