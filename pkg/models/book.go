package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/sandeshlmore/bookslib/pkg/database"
)

type Books struct {
	gorm.Model
	Title       string  `json:"title" validate:"required"`
	Author      string  `json:"author" validate:"required"`
	Price       float64 `json:"price"`
	Publishyear int     `json:"publishyear"`
}

var dbconn *gorm.DB = database.GetDBInstance()

func FetchBooks(per_page int, page_no int, filters map[string]interface{}) ([]Books, int64) {
	var totalRows int64

	var books []Books

	offset := per_page * (page_no - 1)
	log.Println("offset: ", offset, " per_page: ", per_page, " page_no", page_no)

	dbconn.Table("books").Count(&totalRows)

	log.Println("count: ", totalRows)
	if filters != nil {
		log.Println("Info: filters: ", filters)
		q := dbconn.Table("books").Where(filters).Offset(offset).Limit(per_page).Find(&books)
		log.Println("Info: Query: ", q.QueryExpr())

	} else {
		dbconn.Table("books").Offset(offset).Limit(per_page).Find(&books)
	}
	return books, totalRows
}

func InsertNewBook(book *Books) {

	dbconn.Create(&book)
}
