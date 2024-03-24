GoLang API - Gin framework


### Run application 
`go run main.go`    

### Get books
`curl localhost:8080/books`


### Add book
`curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"`