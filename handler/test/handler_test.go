package test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	mockDatabase "github.com/kjasuquo/jumia_task/database/mocks"
	"github.com/kjasuquo/jumia_task/handler"
	"github.com/kjasuquo/jumia_task/models"
	"github.com/kjasuquo/jumia_task/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchNumberHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockDatabase.NewMockDB(ctrl)

	h := &handler.Handler{DB: mockDB}

	route, _ := router.SetupRouter(h)

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

	customersJSON, err := json.Marshal(customers)
	if err != nil {
		t.Fail()
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

	numbersJSON, err := json.Marshal(numbers)
	if err != nil {
		t.Fail()
	}

	nums := []models.Numbers{
		{
			ID:          1,
			Country:     "Morocco",
			Valid:       "NOK",
			CountryCode: "+212",
			Num:         "6007989253",
		},

		{
			ID:          2,
			Country:     "Morocco",
			Valid:       "OK",
			CountryCode: "+212",
			Num:         "698054317",
		},

		{
			ID:          3,
			Country:     "Mozambique",
			Valid:       "OK",
			CountryCode: "+258",
			Num:         "847651504",
		},

		{
			ID:          4,
			Country:     "Mozambique",
			Valid:       "NOK",
			CountryCode: "+258",
			Num:         "042423566",
		},

		{
			ID:          5,
			Country:     "Uganda",
			Valid:       "OK",
			CountryCode: "+256",
			Num:         "704244430",
		},

		{
			ID:          6,
			Country:     "Ethiopia",
			Valid:       "NOK",
			CountryCode: "+251",
			Num:         "9773199405",
		},

		{
			ID:          7,
			Country:     "Cameroon",
			Valid:       "OK",
			CountryCode: "+237",
			Num:         "697151594",
		},

		{
			ID:          8,
			Country:     "Uganda",
			Valid:       "NOK",
			CountryCode: "+256",
			Num:         "7503O6263",
		},

		{
			ID:          9,
			Country:     "Ethiopia",
			Valid:       "OK",
			CountryCode: "+251",
			Num:         "914701723",
		},

		{
			ID:          10,
			Country:     "Cameroon",
			Valid:       "NOK",
			CountryCode: "+237",
			Num:         "6A0311634",
		},
	}

	numsJSON, err := json.Marshal(nums)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for successful queries in SearchHandle", func(t *testing.T) {
		mockDB.EXPECT().ScanCustomerTable().Return(customers, nil)
		mockDB.EXPECT().ScanNumberTable().Return(nums, nil)
		mockDB.EXPECT().InsertToNumbersTable(numbers, nums).Return(nil)
		mockDB.EXPECT().Search("", "").Return(numbers, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", strings.NewReader(string(numbersJSON)))
		route.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), string(numbersJSON))
	})

	t.Run("Testing for errors in ScanCustomerTable func", func(t *testing.T) {
		mockDB.EXPECT().ScanCustomerTable().Return(nil, errors.New("error from ScanCustomer"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", strings.NewReader(string(customersJSON)))
		route.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "error from ScanCustomerTable")
	})

	t.Run("Testing for errors in ScanNumberTable func", func(t *testing.T) {
		mockDB.EXPECT().ScanCustomerTable().Return(customers, nil)
		mockDB.EXPECT().ScanNumberTable().Return(nil, errors.New("error from ScanNumberTable"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", strings.NewReader(string(numsJSON)))
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "error from ScanNumberTable")
	})

	t.Run("Testing for errors in InsertToNumberTable func", func(t *testing.T) {
		mockDB.EXPECT().ScanCustomerTable().Return(customers, nil)
		mockDB.EXPECT().ScanNumberTable().Return(nums, nil)
		mockDB.EXPECT().InsertToNumbersTable(numbers, nums).Return(errors.New("error Inserting numbers to table"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", strings.NewReader(string(numbersJSON)))
		route.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "error Inserting number into Table")
	})

	t.Run("Testing for errors in Search func", func(t *testing.T) {
		mockDB.EXPECT().ScanCustomerTable().Return(customers, nil)
		mockDB.EXPECT().ScanNumberTable().Return(nums, nil)
		mockDB.EXPECT().InsertToNumbersTable(numbers, nums).Return(nil)
		mockDB.EXPECT().Search("", "").Return(nil, errors.New("error searching database"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/", strings.NewReader(string(numbersJSON)))
		route.ServeHTTP(rw, req)

		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "error searching database")
	})

}
