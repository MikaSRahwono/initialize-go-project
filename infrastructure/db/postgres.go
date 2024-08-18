package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler(connectString string, dbName string) (DBHandler, error) {
	dbHandler := DBHandler{}
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}
	dbHandler.DB = db
	return dbHandler, nil
}

// func (dbHandler DBHandler) FindAllBooks() ([]*domain.Book, error) {
// 	var books []*domain.Book
// 	if err := dbHandler.DB.Find(&books).Error; err != nil {
// 		return nil, err
// 	}
// 	return books, nil
// }

// func (dbHandler DBHandler) SaveBook(book domain.Book) error {
// 	if err := dbHandler.DB.Create(&book).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (dbHandler DBHandler) SaveAuthor(author domain.Author) error {
// 	if err := dbHandler.DB.Create(&author).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
