package db

import (
	"finvestapi/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable", dbHost, dbName, dbUser, dbPassword, dbPort)
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbName == "" {
		connectionString = "host=localhost dbname=postgres user=postgres password=postgres port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate() error {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(db)

	err = db.AutoMigrate(&models.News{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "News",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}

	err = db.AutoMigrate(&models.NewsLike{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "NewsLike",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}

	err = db.AutoMigrate(&models.NewsViewing{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "NewsViewing",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}

	err = db.AutoMigrate(&models.NewsAnalytics{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "NewsAnalytics",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return &MigrateError{
			Model:   "User",
			Message: err.Error(),
			Err:     ErrAutoMigrate, // Wrap the sentinel error
		}
	}

	return nil
}
