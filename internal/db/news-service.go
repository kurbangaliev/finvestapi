package db

import (
	"finvestapi/internal/models"
	"log"
)

func LikeNews(item models.NewsLike) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	// Get the underlying *sql.DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closing of the underlying connection pool
	defer sqlDB.Close()

	result := db.Raw("CALL public.sp_likenews(?, ?, ?)", item.NewsId, item.UserId, item.Type).Scan(&struct{}{})
	if result.Error != nil {
		log.Fatal("Error calling stored procedure:", result.Error)
	}

	return nil
}

func ViewNews(item models.NewsViewing) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	// Get the underlying *sql.DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closing of the underlying connection pool
	defer sqlDB.Close()

	result := db.Raw("CALL public.sp_viewnews(?, ?)", item.NewsId, item.UserId).Scan(&struct{}{})
	if result.Error != nil {
		log.Fatal("Error calling stored procedure:", result.Error)
	}

	return nil
}
