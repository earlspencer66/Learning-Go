### Go and MySQL - Project on a Bookstore Management API
1. Database --- MySQL
2. GORM
```cmd
go get "github.com/jinzhu/gorm/"

go get "github.com/jinzhu/gorm/dialects/mysql"
```
3. Json marshall, unmarshall
4. Project Structure (See flowchart)
5. Gorilla mux

```cmd
go get "github.com/gorilla/mux"
```

#### Steps
1. Start with creating routes in bookstoreroutes.go file --- link with Postman -- contains the routes
2. Go to app.go --- to interact with the DB MySQL
```go
//check documentation for...
import (
    _ "github.com/jihnzu/gorm/dialects/mysql"
 )
```
3. Utils.go file..used for unmarshiling
4. Work on models i.e. the book.go file
5. Add main.go file
6. All functions included in the book.go
7. Update the controllers i.e. book-controller.go
8. Check for bugs and test using postman

