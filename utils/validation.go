package utils

import (
	"github.com/kjasuquo/jumia_task/models"
	"regexp"
)

//CameroonValidation validates all phone numbers with Cameroon area code using regex
func CameroonValidation(number, code, phoneNum string) models.Numbers {
	var num models.Numbers

	num.CountryCode = code
	num.Country = "Cameroon"
	num.Num = phoneNum

	var re = regexp.MustCompile(`\(237\)\ ?[2368]\d{7,8}$`)
	if re.MatchString(number) {
		num.Valid = "OK"
	} else {
		num.Valid = "NOK"
	}

	return num
}

//EthiopiaValidation validates all phone numbers with Ethiopia area code using regex
func EthiopiaValidation(number, code, phoneNum string) models.Numbers {
	var num models.Numbers

	num.CountryCode = code
	num.Country = "Ethiopia"
	num.Num = phoneNum

	var re = regexp.MustCompile(`\(251\)\ ?[1-59]\d{8}$`)
	if re.MatchString(number) {
		num.Valid = "OK"
	} else {
		num.Valid = "NOK"
	}

	return num
}

//MoroccoValidation validates all phone numbers with Morocco area code using regex
func MoroccoValidation(number, code, phoneNum string) models.Numbers {
	var num models.Numbers

	num.CountryCode = code
	num.Country = "Morocco"
	num.Num = phoneNum

	var re = regexp.MustCompile(`\(212\)\ ?[5-9]\d{8}$`)
	if re.MatchString(number) {
		num.Valid = "OK"
	} else {
		num.Valid = "NOK"
	}

	return num
}

//MozambiqueValidation validates all phone numbers with Mozambique area code using regex
func MozambiqueValidation(number, code, phoneNum string) models.Numbers {
	var num models.Numbers

	num.CountryCode = code
	num.Country = "Mozambique"
	num.Num = phoneNum

	var re = regexp.MustCompile(`\(258\)\ ?[28]\d{7,8}$`)
	if re.MatchString(number) {
		num.Valid = "OK"
	} else {
		num.Valid = "NOK"
	}

	return num
}

//UgandaValidation validates all phone numbers with Uganda area code using regex
func UgandaValidation(number, code, phoneNum string) models.Numbers {
	var num models.Numbers

	num.CountryCode = code
	num.Country = "Uganda"
	num.Num = phoneNum

	var re = regexp.MustCompile(`\(256\)\ ?\d{9}$`)
	if re.MatchString(number) {
		num.Valid = "OK"
	} else {
		num.Valid = "NOK"
	}

	return num
}
