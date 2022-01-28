package main

import (
	// "database/sql"

	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// var db *gorm.DB

type Books struct {
	gorm.Model  `json:"-"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Publishyear int     `json:"publishyear"`
}

// func sqlconn() {
// 	dbURL := "postgres://booksuser:p@ssw0rd123@localhost:5432/books_db?sslmode=disable"
// 	db2, err := sql.Open("postgres", dbURL)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	}
// 	rows, _ := db2.Query("select title from books")
// 	var name string
// 	for rows.Next() {
// 		rows.Scan(&name)
// 		fmt.Println(name)

// 	}
// }

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(&Books{})
}

func DBConnect() *gorm.DB {
	psqlurl := "postgres://booksuser:p@ssw0rd123@localhost:5432/books?sslmode=disable"

	db, err := gorm.Open("postgres", psqlurl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetBooks(db *gorm.DB) []Books {
	var books []Books
	db.Find(&books)
	return books
}

func GetBooksWithFilters(db *gorm.DB) []Books {
	var totalRows int64
	var books []Books

	db.Table("books").Count(&totalRows)
	fmt.Println("count: ", totalRows)
	per_page := 3
	page_no := 2
	offset := per_page * (page_no - 1)

	db.Table("books").Where("price > ?", 6).Offset(offset).Limit(per_page).Find(&books)

	fmt.Println("books: ", books)
	fmt.Println("books: ")

	return []Books{}
}

func AddBook(db *gorm.DB) {
	book := Books{Title: "Alice in the Wonderland", Author: "Lewis Carroll", Price: 5.5, Publishyear: 2006}
	db.Table("books").Where("id = ?", 6).Find(&book)
	data, _ := json.Marshal(book)
	fmt.Println("book data: ", string(data))
	// json.NewDecoder("okok").Decode(&book)
	db.Create(&book)

}

func main() {
	db := DBConnect()
	InitialMigration(db)
	// fmt.Println("db connection: ", db)
	// books := GetBooks(db)
	// for _, book := range books {
	// 	fmt.Println(book)
	// }
	// data, err := json.MarshalIndent(books, "", "\t")
	// fmt.Println("data: ", string(data), "error: ", err)

	// GetBooksWithFilters(db)
	// AddBook(db)
}
