package main

import (
	"log"

	"github.com/fadhilradh/simple-auth/db"
	"github.com/fadhilradh/simple-auth/domains/user"
	"github.com/fadhilradh/simple-auth/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize databse connection: %s", err)
	} else {
		log.Print("DB connection succesful")
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	router.InitRouter(userHandler)
	router.Start("localhost:8080")
}