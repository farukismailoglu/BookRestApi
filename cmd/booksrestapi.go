package main

import (
	"./repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

func getBooks(w http.ResponseWriter, r *http.Request, title string) {
	log.Println("getBooks function is running")
	w.Header().Set("Content-Type", "application/json")
	books := repository.GetAllBook()
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request, title string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	book := repository.FindBook(params["isbn"])
	json.NewEncoder(w).Encode(book)

}

func createBook(w http.ResponseWriter, r *http.Request, title string) {
	w.Header().Set("Content-Type", "application/json")
	var book repository.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	repository.InsertSingleBook(book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request, title string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book repository.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	repository.UpdateBook(params["isbn"], book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request, title string) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	repository.DeleteBook(params["isbn"])
	w.Write([]byte("Deleted book isbn:" + params["isbn"]))
}

var validPath = regexp.MustCompile("^/(books)(/([a-zA-Z0-9]+)$)?")

func doWorkForHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func InitHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", doWorkForHandler(getBooks)).Methods("GET")
	r.HandleFunc("/books/{isbn}", doWorkForHandler(getBook)).Methods("GET")
	r.HandleFunc("/books", doWorkForHandler(createBook)).Methods("POST")
	r.HandleFunc("/books/{isbn}", doWorkForHandler(updateBook)).Methods("PUT")
	r.HandleFunc("/books/{isbn}", doWorkForHandler(deleteBook)).Methods("DELETE")

	return r
}
func StartServer(r *mux.Router) {
	port := ":5000"
	log.Printf("Starting server listen on %v port", port)
	err := http.ListenAndServe(port, r)
	log.Fatal(err)
}

func main() {
	router := InitHandlers()
	StartServer(router)

}
