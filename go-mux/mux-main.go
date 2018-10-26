package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type apiRes struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	r.HandleFunc("/", getBook).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/gps", getGPS).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//func getBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	api := apiRes{Status: 200, Data: "This is Home!"}
//	json.NewEncoder(w).Encode(api)
//}
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

type gps struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Time int     `json:"time"`
}

var res []gps

func getGPS(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/loglag_gps")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT latitude, longitude, created_at FROM gps_event_logs WHERE order_vehicle_id = 29 AND driver_id = 395")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var latitude, longitude float64
		var created_at int
		err = result.Scan(&latitude, &longitude, &created_at)
		if err != nil {
			panic(err.Error())
		}
		re := gps{latitude, longitude, created_at}
		res = append(res, re)
	}
	json.NewEncoder(w).Encode(res)
}
