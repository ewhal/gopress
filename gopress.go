package main

import (
	"database/sql"
	"html"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Page stuct {
	Title string
	Post string
	Date string
	Author string
	Email string

}

// checkErr logger
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)
	defer db.Close()

	query, err := db.Query("select * from post")
	checkErr(err)

	b := []Page{}
	for query.Next() {
		p := Page{}
		var input string
		query.Scan(&p.Title, &input, &p.Date, &p.Author, &p.Email)
		unsafe := blackfriday.MarkdownCommon(input)
		p.Post = bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		p = append(b, p)


	}
	err = templates.ExecuteTemplate(w, "index.html", &b)
	checkErr(err)

}

func rssHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)
	defer db.Close()

	query, err := db.Query("select * from post")
	checkErr(err)

	b := []Page{}
	for query.Next() {
		p := Page{}
		var input string
		query.Scan(&p.Title, &input, &p.Date, &p.Author, &p.Email)
		unsafe := blackfriday.MarkdownCommon(input)
		p.Post = bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		p = append(b, p)


	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post := vars["id"]

	db, err := sql.Open("mysql", DATABASE)
	checkErr(err)
	defer db.Close()

	query, err := db.Query("select * from post where name=?", html.EscapeString(post))
	checkErr(err)

	p := Page{}
	for query.Next() {
		var input string
		query.Scan(&p.Title, &input, &p.Date, &p.Author, &p.Email)
		unsafe := blackfriday.MarkdownCommon(input)
		p.Post = bluemonday.UGCPolicy().SanitizeBytes(unsafe)


	}
	err = templates.ExecuteTemplate(w, "post.html", &p)
	checkErr(err)

}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func registerHandler(w http.ResponseWriter, r *http.Request) {

}

func newHandler(w http.ResponseWriter, r *http.Request) {

}

func delHandler(w http.ResponseWriter, r *http.Request) {

}

func modifyHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/rss", rssHandler)
	router.HandleFunc("/post/{id}", postHandler)
	router.HandleFunc("/404", notfoundHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/register", registerHandler)
	router.HandleFunc("/post/new", newHandler)
	router.HandleFunc("/post/del/{id}", delHandler)
	router.HandleFunc("/post/modify/{id}", modifyHandler)

	err := http.ListenAndServe(PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
