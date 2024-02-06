package main

import (
  "fmt"
  "text/template"
  "net/http"
  "log"
)

func main() {
  //server code here
  fmt.Println("Hello World")

  h1 := func(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("./html/index.html"))
    tmpl.Execute(w, nil)
  }

  http.HandleFunc("/", h1)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
