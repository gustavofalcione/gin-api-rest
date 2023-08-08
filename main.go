package main

import (
	"github.com/gustavofalcione/api-go-gin/database"
	"github.com/gustavofalcione/api-go-gin/routes"
)

func main() {
	database.ConnectingDatabase()
	routes.HandleRequests()
}
