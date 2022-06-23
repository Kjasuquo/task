package services

import (
	"errors"
	"github.com/kjasuquo/jumia_task/models"
	"regexp"
	"strings"
)

type CountryValidation struct {
	Country string
	Code    string
	Regex   string
}

var (
	validations = map[string]CountryValidation{
		"(237)": CountryValidation{
			Country: "Cameroon",
			Code:    "+237",
			Regex:   `\(237\)\ ?[2368]\d{7,8}$`,
		},

		"(251)": CountryValidation{
			Country: "Ethiopia",
			Code:    "+251",
			Regex:   `\(251\)\ ?[1-59]\d{8}$`,
		},

		"(212)": CountryValidation{
			Country: "Morocco",
			Code:    "+212",
			Regex:   `\(212\)\ ?[5-9]\d{8}$`,
		},

		"(258)": CountryValidation{
			Country: "Mozambique",
			Code:    "+258",
			Regex:   `\(258\)\ ?[28]\d{7,8}$`,
		},

		"(256)": CountryValidation{
			Country: "Uganda",
			Code:    "+256",
			Regex:   `\(256\)\ ?\d{9}$`,
		},
	}
)

func (c *CountryValidation) Validate(phoneNumber string) bool {
	var re = regexp.MustCompile(c.Regex)
	return re.MatchString(phoneNumber)
}

func ValidatePhoneNumber(phoneNumber string) (models.Numbers, error) {
	var err = errors.New("error in ValidatePhoneNumber")
	var empty models.Numbers
	number := strings.Split(phoneNumber, " ")
	validation, ok := validations[number[0]]
	if ok {
		number := models.Numbers{
			Country:     validation.Country,
			CountryCode: validation.Code,
			Num:         number[1],
		}
		if validation.Validate(phoneNumber) {
			number.Valid = "OK"
			return number, nil
		} else if !validation.Validate(phoneNumber) {
			number.Valid = "NOK"
			return number, nil
		}
	}
	return empty, err
}
