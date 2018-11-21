package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "This is a test Description", Content: "The Contents of This Article are generated for testing a restful api in Golang"},
	}

	fmt.Println("<h1>All Articles EndPoint Works!</h1>")
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Post Endpoint for Articles is responding...</h1>")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Serving Plain Text and HTML</h1>")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
