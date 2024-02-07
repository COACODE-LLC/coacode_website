package main

import (
  "fmt"
  "text/template"
  "net/http"
  "log"
  "time"
)

func main() {
  //server code here
  fmt.Println("Initializing http server...")

  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/about", aboutHandler)
  http.HandleFunc("/contact", contactHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func printLogs() {
  time.Sleep(1000 * time.Millisecond)
  fmt.Println("Server online: Localhost:8080")
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
