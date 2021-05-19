package main

import (
	"go-auth/models"
	"log"
)

func main() {
	// config.Init()

	// database.StartDatabase()
	// s := server.NewServer()
	// s.Run()

	va := models.NewUser("yuri", "yuri.snp@hotmail.com", "123456")
	log.Print(va.String())

}
