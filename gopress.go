package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func rssHandler(w http.ResponseWriter, r *http.Request) {

}

func postHandler(w http.ResponseWriter, r *http.Request) {

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
