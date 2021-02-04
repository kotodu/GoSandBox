package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func webPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my site!")
	fmt.Println("home page endpoint")
}

func handleRequests() {
	http.HandleFunc("/", webPage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{}
	for i := 0; i < 10; i++ {
		title := "Hello_%d"
		articles = append(articles,
			Article{
				Title:   fmt.Sprintf(title, i),
				Desc:    "Article Description",
				Content: "Article Content"})
	}
	fmt.Println("hit:AllArticles")
	json.NewEncoder(w).Encode(articles)
}
