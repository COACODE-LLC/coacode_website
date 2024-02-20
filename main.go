package main

import (
  "fmt"
  "text/template"
  "net/http"
  "log"
  "github.com/savioxavier/termlink"
)

func main() {
  //server code here
  fmt.Println("Initializing http server to...")
  fmt.Println("Connected to:", termlink.Link("localhost:8080", "http://localhost:8080")) 

  //handle html template serving
  http.HandleFunc("/", pageHandler)

  //handle static file serving
  http.HandleFunc("/assets/", assetHandler)
  
  //Create server on port :8080
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func assetHandler(w http.ResponseWriter, r *http.Request) {
  //gets url path from r, searches for file in src/assets/
  http.ServeFile(w, r, "./src/assets/"+r.URL.Path[8:])
}


func pageHandler(w http.ResponseWriter, r *http.Request) {
  pages := map[string]string {
    "/": "index.html",
    "/about": "src/pages/about.html",
    "/contact": "src/pages/contact.html",
  }

  renderTemplate(w, pages[r.URL.Path])
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
