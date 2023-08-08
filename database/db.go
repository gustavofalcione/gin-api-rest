package database

import (
	"log"

	"github.com/gustavofalcione/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectingDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Error when connecting database ❌")
	}

	DB.AutoMigrate(&models.Student{})
}
