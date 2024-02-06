package main

import (
  "fmt"
  "text/template"
  "net/http"
  "log"
)

type PageData struct {
  Title string
}

func main() {
  //server code here
  fmt.Println("Hello World")

  homeHandler := func(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("./html/index.html"))
    tmpl.Execute(w, nil)
    log.Println(r)
  }

  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/about", aboutHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func aboutHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("Hello About")

  renderTemplate(w, "html/about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
  //Parse HTML template files
  /*tmplFiles := []string {
    "html/index.html",
    "html/about.html",
  }*/

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
