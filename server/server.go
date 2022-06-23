package server

import (
	"fmt"
	"github.com/kjasuquo/jumia_task/database"
	"github.com/kjasuquo/jumia_task/handler"
	"github.com/kjasuquo/jumia_task/router"
	"log"
)

//Start starts our router
func Start() error {

	//Setting up the Sqlite Database
	var sdb = new(database.SqliteDb)

	h := &handler.Handler{DB: sdb}
	err := sdb.Init()
	if err != nil {
		log.Println("Error trying to Init", err)
		return err
	}

	route, port := router.SetupRouter(h)
	fmt.Println("connected on port ", port)
	err = route.Run(port)
	if err != nil {
		log.Printf("Error from SetupRouter :%v", err)
		return err
	}

	return nil
}
