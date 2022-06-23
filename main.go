package main

import (
	"fmt"
	"github.com/kjasuquo/jumia_task/config"
	"github.com/kjasuquo/jumia_task/server"
	"log"
)

func main() {
	fmt.Println("Starting... Server")

	//Gets environment variables if .env file is available
	config.NewConfig(".env")

	//Start server
	err := server.Start()
	if err != nil {
		log.Println("Error starting server", err)
		return
	}
}
