package db

import (
	"finvestapi/internal/models"
	"fmt"
	"log"
	"os"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenConnection for database postgres
func OpenConnection() (*gorm.DB, error) {
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

// AutoMigrateObjects automatically migrates all registered model types
func AutoMigrateObjects() error {
	db, err := OpenConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer sqlDB.Close()

	log.Println(db)

	// Migrate all registered models
	for _, model := range models.ModelRegistry {
		modelType := reflect.TypeOf(model).Elem()
		modelName := modelType.Name()

		err = db.AutoMigrate(model)
		if err != nil {
			log.Printf("Failed to migrate model: %s", modelName)
			return &MigrateError{
				Model:   modelName,
				Message: err.Error(),
				Err:     ErrAutoMigrate,
			}
		}
		log.Printf("Successfully migrated model: %s", modelName)
	}

	return nil
}
