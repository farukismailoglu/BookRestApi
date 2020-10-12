# Simple Book Rest Api Project with Golang and Mongodb

## Endpoints


### Get All Book
```
Get
```

### Get Single Book
```
GET /books/{isbn}
```

### Create New Book
```
POST /books/
```
### Delete Book
```
DELETE /books/{isbn}
```
### Update Single Book
```
PUT /book/{isbn}
```

## Prerequities
* Go 1.15.2
* Docker


## Run 

step1: Mongodb can be run with **docker-compose up** command.
step2: Run the project with **go run bookrestapi.go** command.

## Test

You can use postmanproject/BookRestApi.postman_collection.json file for test with Postman.

 
