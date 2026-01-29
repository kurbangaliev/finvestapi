package db

import (
	"finvestapi/internal/models"
	"fmt"
	"log"
)

func LoadAllNews() ([]models.News, error) {
	var items []models.News

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	result := db.Order("message_date DESC").Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return items, nil
}

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
		log.Printf("Error calling stored procedure:", result.Error)
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
		log.Printf("Error calling stored procedure:", result.Error)
	}

	return nil
}

func GetLikesView(id int) (models.CallProcedureResult, error) {
	db, err := DbConnection()
	if err != nil {
		return models.CallProcedureResult{}, err
	}
	// Get the underlying *sql.DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closing of the underlying connection pool
	defer sqlDB.Close()

	var callResult models.CallProcedureResult
	result := db.Raw("CALL public.sp_getlikesview(?)", id).Scan(&callResult)
	if result.Error != nil {
		log.Printf("Error calling stored procedure:", result.Error)
	}

	return callResult, nil
}

func GetNewsAnalytics(id int) (models.News, error) {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	// Get the underlying *sql.DB connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closing of the underlying connection pool
	defer sqlDB.Close()

	var newsItem models.News
	query := fmt.Sprintf("SELECT id, liked, disliked, view_count FROM news WHERE id = %d", id)
	err = db.Raw(query).Scan(&newsItem).Error
	if err != nil {
		return models.News{}, err
	}
	return newsItem, nil
}
