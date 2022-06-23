package service

import (
	"github.com/kjasuquo/jumia_task/models"
	"github.com/kjasuquo/jumia_task/utils"
	"strings"
)

//GetAllNumbers compiles all the validated numbers to be inserted in the "number" table
func GetAllNumbers(customers []models.Customer) []models.Numbers {
	var data []models.Numbers

	for i := 0; i < len(customers); i++ {
		number := strings.Split(customers[i].Phone, " ")
		if number[0] == "(237)" {
			num := utils.CameroonValidation(customers[i].Phone, "+237", number[1])
			data = append(data, num)
		} else if number[0] == "(251)" {
			num := utils.EthiopiaValidation(customers[i].Phone, "+251", number[1])
			data = append(data, num)
		} else if number[0] == "(212)" {
			num := utils.MoroccoValidation(customers[i].Phone, "+212", number[1])
			data = append(data, num)
		} else if number[0] == "(258)" {
			num := utils.MozambiqueValidation(customers[i].Phone, "+258", number[1])
			data = append(data, num)
		} else if number[0] == "(256)" {
			num := utils.UgandaValidation(customers[i].Phone, "+256", number[1])
			data = append(data, num)
		}
	}
	return data
}
