package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:id`
	Isbn string `json:isbn`
	Title string `json:title`
	Author *Author `json:author`
}

type Author struct {
	Name string `json:name`
	Fname string `json:fname`
}

// init a book var - a slice of the Book
var books []Book

// Get All Books
func firstPage(w http.ResponseWriter, r *http.Request){

}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get One Specific Books
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}


// Create a new or some new Books
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
// Update One Specific Book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}


// Delete One Specific Book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main(){
	// building a new router
	r := mux.NewRouter()

	// Making an initial data for the book
	books = append(books, Book{
		ID:     "1",
		Isbn:   "43218",
		Title:  "Sunset Oasis",
		Author: &Author{
			Name:  "Bahaa",
			Fname: "Taher",
		},
	},
	Book{
		ID:     "2",
		Isbn:   "52136",
		Title:  "Kafka by the sea",
		Author: &Author{
			Name:  "HaruKi",
			Fname: "Morakami",
		},
	})

	// Router Handler / Endpoints
	r.HandleFunc("/", firstPage).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}


