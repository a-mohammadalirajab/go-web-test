package lib

import (
	"os"

	"github.com/a-mohammadalirajab/go-web-test/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbaddress := os.Getenv("DBPATH")
	if len(dbaddress) == 0 {
		dbaddress = "./production.db"
	}
	db, err := gorm.Open(sqlite.Open(dbaddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.TaskModel{})
	return db
}
