package mai

import (
	"go-auth/config"
	"go-auth/database"
	"go-auth/server"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := server.NewServer()
	s.Run()
}
