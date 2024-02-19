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

  
  //create defualtmux connections
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/about", aboutHandler)
  http.HandleFunc("/contact", contactHandler)
  
  //Create server on port :8080
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello Home")
  renderTemplate(w, "index.html")
}


func aboutHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello About")
  renderTemplate(w, "src/pages/about.html")
}


func contactHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello Contact")
  renderTemplate(w, "src/pages/contact.html")
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
