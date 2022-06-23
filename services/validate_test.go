package services

import (
	"errors"
	"github.com/kjasuquo/jumia_task/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePhoneNumberError(t *testing.T) {

	_, err := ValidatePhoneNumber("customers[9].Phone")
	want := errors.New("error in ValidatePhoneNumber")

	assert.Equal(t, want, err)

}

func TestValidatePhoneNumber(t *testing.T) {

	customer := models.Customer{
		ID:    0,
		Name:  "Walid Hammadi",
		Phone: "(212) 6007989253",
	}

	numbers := models.Numbers{
		ID:          0,
		Country:     "Morocco",
		Valid:       "NOK",
		CountryCode: "+212",
		Num:         "6007989253",
	}

	got, _ := ValidatePhoneNumber(customer.Phone)
	want := numbers

	assert.Equal(t, want, got)

}
