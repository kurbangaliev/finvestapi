package db

import (
	"log"
)

func SelectAll[T comparable]() ([]T, error) {
	var items []T

	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}

	result := db.Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return items, nil
}

func SaveObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}
	result := db.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func DeleteObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	result := db.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func UpdateObject[T comparable](item T) error {
	db, err := DbConnection()
	if err != nil {
		return err
	}

	result := db.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
