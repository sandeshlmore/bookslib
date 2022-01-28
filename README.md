# Books REST API with Golang
A RESTful API books library with Go using **gorilla/mux** and **gorm** (An ORM for Go)



### Installation & Run
```bash
# Download this project
go get github.com/sandeshlmore/bookslib

# start application
docker-compose -f docker-compose-local.yaml up --build

# database setup
find scripts at local_development/sqlqueries

```

### Books Schema:

    type Books struct {
        ID          int
        Title       string 
        Author      string  
        Price       float64 
        Publishyear int     
    }


```
### API Endpoint : http://localhost:8080
```

## API

### /books
* `GET` : Get Books
* `POST` : Add New Book


## API Usage:

### Get Books:
    Books can be retrieved by id, title, author, year
    per_page defaults to 10
    page_no dfaults to 1

    Request1: 
    curl --location --request GET 'http://localhost:8080/books?id=1' --header 'Content-Type: application/json'

    Sample Response:
        
        {
            "books": [
                {
                    "ID": 1,
                    "title": "Emma",
                    "author": "Jayne Austen",
                    "price": 9.44,
                    "publishyear": 2000,
                    "CreatedAt": "2022-01-28T06:34:57.743349016Z",
                    "UpdatedAt": "2022-01-28T06:34:57.743349016Z",
                    "DeletedAt": null,
                }
            ],
            "totalbooks": 36
        }
        status code : 200


    Request2: 
    curl --location --request GET 'http://localhost:8080/books?per_page=3&page_no=1'

    Sample Response:

        {
            "books": [
                {
                    "title": "Harry Potter",
                    "author": "J. K. Rowling",
                    "price": 16.99,
                    "publishyear": 2003,
                    "CreatedAt": "2022-01-28T06:34:57.743349016Z",
                    "UpdatedAt": "2022-01-28T06:34:57.743349016Z",
                    "DeletedAt": null,
                },
                {
                    "title": "Adventures of Sherlock Holmes",
                    "author": "Sir Arthur Conan Doyle",
                    "price": 10.99,
                    "publishyear": 2005,
                    "CreatedAt": "2022-01-28T06:34:57.743349016Z",
                    "UpdatedAt": "2022-01-28T06:34:57.743349016Z",
                    "DeletedAt": null,
                },
                {
                    "title": "Alice in the Wonderland",
                    "author": "Lewis Carroll",
                    "price": 5.5,
                    "publishyear": 2006,
                    "CreatedAt": "2022-01-28T06:34:57.743349016Z",
                    "UpdatedAt": "2022-01-28T06:34:57.743349016Z",
                    "DeletedAt": null,
                }
            ],
            "totalbooks": 38
        }
    
        

### Add New Book:

    Request:
    curl --location --request POST 'http://localhost:8080/books' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title": "TestBook",
            "author": "John Doe",
            "price": 101,
            "publishyear": 2001
    }'

    Response:
    {
        "ID": 387,
        "CreatedAt": "2022-01-28T04:54:07.227616939Z",
        "UpdatedAt": "2022-01-28T04:54:07.227616939Z",
        "DeletedAt": null,
        "title": "TestBook",
        "author": "John Doe",
        "price": 101,
        "publishyear": 2001
    }   
    status code : 200
