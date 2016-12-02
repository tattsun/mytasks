package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

func NewDB(dbPath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	logrus.Info("migrating db")
	db.AutoMigrate(&Task{}, &Project{}, &Section{})

	return db
}
