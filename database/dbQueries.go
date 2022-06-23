package database

import (
	"database/sql"
	"fmt"
	"github.com/kjasuquo/jumia_task/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteDb struct {
	DB *sql.DB
}

//Init runs first. It Opens the sample Database and create the necessary "number" table
func (sl *SqliteDb) Init() error {

	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		log.Println(err)
		return err
	}

	CreateNumberTable(db)

	sl.DB = db

	return nil
}

// CreateNumberTable is a database function with Query for creating a table called number. It is called in Init func
func CreateNumberTable(db *sql.DB) {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS number (id INTEGER PRIMARY KEY, country VARCHAR(50), valid VARCHAR(50), country_code VARCHAR(50), num VARCHAR(50))")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println(err)
		return
	}
}

//ScanCustomerTable contains the queries for retrieving data from the given customer table in the sample db
func (sl *SqliteDb) ScanCustomerTable() ([]models.Customer, error) {

	statement := "SELECT * FROM customer"
	rows, err := sl.DB.Query(statement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

//ScanNumberTable contains the queries for retrieving data from the given customer table in the sample db
func (sl *SqliteDb) ScanNumberTable() ([]models.Numbers, error) {

	statement := "SELECT * FROM number"
	rows, err := sl.DB.Query(statement)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var numbers []models.Numbers
	for rows.Next() {
		var number models.Numbers
		err := rows.Scan(&number.ID, &number.Country, &number.Valid, &number.CountryCode, &number.Num)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

//InsertToNumbersTable receives the slice of number after validation/compilation and inserts it into the table if table is empty
func (sl *SqliteDb) InsertToNumbersTable(numbers, dbNum []models.Numbers) error {

	if len(dbNum) < 1 {
		for i := 0; i < len(numbers); i++ {

			statement, err := sl.DB.Prepare("INSERT INTO number(country, valid, country_code, num) VALUES (?, ?, ?, ?)")
			if err != nil {
				log.Println(err)
				return err
			}
			_, err = statement.Exec(numbers[i].Country, numbers[i].Valid, numbers[i].CountryCode, numbers[i].Num)
			if err != nil {
				log.Println(err)
				return err
			}

		}

	}

	return nil
}

//Search filters numbers in the database
func (sl *SqliteDb) Search(country, state string) ([]models.Numbers, error) {
	var numbers []models.Numbers
	var rows *sql.Rows
	var err error

	if country == "" && state == "" {
		statement := "SELECT * FROM number"
		rows, err = sl.DB.Query(statement)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else if country == "" && state != "" {

		rows, err = sl.DB.Query("SELECT * FROM number WHERE valid = ?;", state)
		if err != nil {
			log.Println(err)
			return nil, err
		}

	} else if country != "" && state == "" {
		rows, err = sl.DB.Query("SELECT * FROM number WHERE country = ?;", country)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else if country != "" && state != "" {
		stmt, err := sl.DB.Prepare("select * from number where country = ? AND valid = ?")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		rows, err = stmt.Query(country, state)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var number models.Numbers
		err = rows.Scan(&number.ID, &number.Country, &number.Valid, &number.CountryCode, &number.Num)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}
