package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
	"github.com/savioxavier/termlink"
)

func main() {
  //server code here
  fmt.Println("Initializing http server to...")
  fmt.Println("Connected to:", termlink.Link("localhost:8080", "http://localhost:8080")) 
  
  //create Gorilla router
  router := mux.NewRouter()

  //handle html template serving
  router.HandleFunc("/", homeHandler)
  router.HandleFunc("/{path}", pageHandler)
  
  //Create server on port :8080
  log.Fatal(http.ListenAndServe(":8080", router))
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path)
  renderTemplate(w, "index.html")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path)
  tmpl := "src/pages" + r.URL.Path + ".html"

  renderTemplate(w, tmpl)
}


func renderTemplate(w http.ResponseWriter, tmpl string) {
  //Parse template file
  t, err := template.ParseFiles(tmpl);
  check(err, w)

  //Execute Template
  err = t.Execute(w, nil);
  check(err, w)
}


func check(err error, w http.ResponseWriter) {
  if (err != nil) {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
