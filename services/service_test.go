package services

import (
	"github.com/kjasuquo/jumia_task/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllNumbers(t *testing.T) {

	customers := []models.Customer{
		{
			ID:    0,
			Name:  "Walid Hammadi",
			Phone: "(212) 6007989253",
		},

		{
			ID:    1,
			Name:  "Yosaf Karrouch",
			Phone: "(212) 698054317",
		},

		{
			ID:    2,
			Name:  "Edunildo Gomes Alberto",
			Phone: "(258) 847651504",
		},

		{
			ID:    3,
			Name:  "Solo Dolo",
			Phone: "(258) 042423566",
		},

		{
			ID:    4,
			Name:  "VINEET SETH",
			Phone: "(256) 704244430",
		},

		{
			ID:    5,
			Name:  "shop23 sales",
			Phone: "(251) 9773199405",
		},

		{
			ID:    6,
			Name:  "EMILE CHRISTIAN KOUKOU DIKANDA HONORE",
			Phone: "(237) 697151594",
		},

		{
			ID:    7,
			Name:  "Kiwanuka Budallah",
			Phone: "(256) 7503O6263",
		},

		{
			ID:    8,
			Name:  "Filimon Embaye",
			Phone: "(251) 914701723",
		},

		{
			ID:    9,
			Name:  "ARREYMANYOR ROLAND TABOT",
			Phone: "(237) 6A0311634",
		},
	}

	numbers := []models.Numbers{
		{
			ID:          0,
			Country:     "Morocco",
			Valid:       "NOK",
			CountryCode: "+212",
			Num:         "6007989253",
		},

		{
			ID:          0,
			Country:     "Morocco",
			Valid:       "OK",
			CountryCode: "+212",
			Num:         "698054317",
		},

		{
			ID:          0,
			Country:     "Mozambique",
			Valid:       "OK",
			CountryCode: "+258",
			Num:         "847651504",
		},

		{
			ID:          0,
			Country:     "Mozambique",
			Valid:       "NOK",
			CountryCode: "+258",
			Num:         "042423566",
		},

		{
			ID:          0,
			Country:     "Uganda",
			Valid:       "OK",
			CountryCode: "+256",
			Num:         "704244430",
		},

		{
			ID:          0,
			Country:     "Ethiopia",
			Valid:       "NOK",
			CountryCode: "+251",
			Num:         "9773199405",
		},

		{
			ID:          0,
			Country:     "Cameroon",
			Valid:       "OK",
			CountryCode: "+237",
			Num:         "697151594",
		},

		{
			ID:          0,
			Country:     "Uganda",
			Valid:       "NOK",
			CountryCode: "+256",
			Num:         "7503O6263",
		},

		{
			ID:          0,
			Country:     "Ethiopia",
			Valid:       "OK",
			CountryCode: "+251",
			Num:         "914701723",
		},

		{
			ID:          0,
			Country:     "Cameroon",
			Valid:       "NOK",
			CountryCode: "+237",
			Num:         "6A0311634",
		},
	}

	got := GetAllValidatedNumbers(customers)
	want := numbers

	assert.Equal(t, want, got)

}
