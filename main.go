package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

var noteStore = make(map[string]Note)

var id int = 0

func main() {
	r := mux.NewRouter().StrictSlash(false)
	//fs := http.FileServer(http.Dir("public"))
	//r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Handler: r,
		Addr:    ":9000",
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
