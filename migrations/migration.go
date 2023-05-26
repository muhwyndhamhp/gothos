package main

import (
	"fmt"
	"muhwyndham/gothos/hello/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		"localhost",
		"5432",
		"postgres",
		"gothos_db",
		"postgres",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := Migrate(db); err != nil {
		panic(err)
	}
}

func Migrate(db *gorm.DB) error {
	db.Debug()
	err := migrate20230526_084108(db)
	if err != nil {
		return err
	}
	return nil
}
func migrate20230526_084108(db *gorm.DB) error {
	return db.Debug().AutoMigrate(&models.Photo{})
}
