package database

import "github.com/kjasuquo/jumia_task/models"

type DB interface {
	ScanCustomerTable() ([]models.Customer, error)
	ScanNumberTable() ([]models.Numbers, error)
	InsertToNumbersTable(numbers, dbNum []models.Numbers) error
	Search(country, state string) ([]models.Numbers, error)
}
