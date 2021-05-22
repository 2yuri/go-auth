package main

import (
	"go-auth/database"
	"go-auth/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()

	// u := models.NewUser("yuri", "yuri.miguel@ecocentauro.com.br", "123456")
	// log.Print(u.Password)
}
