package main

import (
	editor "../editor"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../../ui/")))
	http.HandleFunc("/edit", EditPageHandler)
	http.HandleFunc("/save", SavePageHandler)
	http.HandleFunc("/delete", DeletePageHandler)
	http.ListenAndServe(":8080", nil)
}

func EditPageHandler(rw http.ResponseWriter, r *http.Request) {
	//rw.Write([]byte(editor.Render()))
	http.ListenAndServe(":8080", editor.Render())
}

func SavePageHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("The save handler"))
}

func DeletePageHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("The delete handler"))
}
