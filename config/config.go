package config

import (
	"github.com/joho/godotenv"
	"log"
)

//NewConfig loads environment variable
func NewConfig(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Println(err)
		return
	}
}
