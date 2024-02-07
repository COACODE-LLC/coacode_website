package main

import (
  "fmt"
  "text/template"
  "net/http"
  "log"
)

func main() {
  //server code here
  fmt.Println("Initializing http server...")
  
  //create defualtmux connections
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/about", aboutHandler)
  http.HandleFunc("/contact", contactHandler)
  
  //Create server on port :8080
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func contactHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello Contact")
  renderTemplate(w, "html/contact.html")
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello Home")
  renderTemplate(w, "html/index.html")
}


func aboutHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello About")
  renderTemplate(w, "html/about.html")
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
