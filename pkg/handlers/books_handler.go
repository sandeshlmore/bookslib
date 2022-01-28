package handlers

import (
	"encoding/json"
	"strconv"

	"fmt"
	"log"
	"net/http"

	"github.com/sandeshlmore/bookslib/pkg/models"
	"gopkg.in/go-playground/validator.v9"
)

type Error struct {
	Url     string
	Code    int
	Message string
}

func (eati Error) Error() string {
	return fmt.Sprintf(
		"invalid access token - url: '%s', response code: %d, response: '%s'",
		eati.Url,
		eati.Code,
		eati.Message,
	)
}

func parseFilters(r *http.Request) map[string]interface{} {
	id := r.URL.Query().Get("id")
	title := r.URL.Query().Get("title")
	year := r.URL.Query().Get("year")
	author := r.URL.Query().Get("author")

	filters := make(map[string]interface{})

	if len(id) != 0 {
		filters["id"] = id
	}
	if len(title) != 0 {
		filters["title"] = title
	}
	if len(year) != 0 {
		filters["publishyear"] = year
	}
	if len(author) != 0 {
		filters["author"] = author
	}

	return filters
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	filters := parseFilters(r)
	query_params := r.URL.Query()

	var perpage string
	var pageno string

	if query_params.Has("per_page") {
		perpage = query_params.Get("per_page")
	} else {
		perpage = "10"
	}

	if query_params.Has("page_no") {
		pageno = query_params.Get("page_no")
	} else {
		pageno = "1"
	}

	per_page, err := strconv.Atoi(perpage)
	if err != nil {
		log.Println("Error: ", err.Error())
		makeError(w, http.StatusBadRequest, err.Error())
		return
	}

	page_no, err := strconv.Atoi(pageno)
	if err != nil {
		log.Println("Error: ", err)
		makeError(w, http.StatusBadRequest, err.Error())
		return
	}

	books, count := models.FetchBooks(per_page, page_no, filters)

	response := map[string]interface{}{"books": books, "totalbooks": count}

	data, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		log.Println("json validation error for books: ", response)
		makeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write(data)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Books

	json.NewDecoder(r.Body).Decode(&book)

	v := validator.New()
	validation_err := v.Struct(&book)
	if validation_err != nil {
		log.Println(validation_err.Error())
		makeError(w, http.StatusBadRequest, validation_err.Error())
		return
	}
	log.Println("Info: ", "Adding book: ", book)

	models.InsertNewBook(&book)

	json.NewEncoder(w).Encode(book)

}
