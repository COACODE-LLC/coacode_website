package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"func/mail"

	"github.com/gorilla/mux"
	"github.com/savioxavier/termlink"
)

var pageTemplates = map[string]string{
	"/":        "index.html",
	"/about":   "src/pages/about.html",
	"/contact": "src/pages/contact.html",
}

func main() {
	//server code here
	log.Println("Initializing http server to...")
	log.Println("Connected to:", termlink.Link("localhost:80", "http://localhost:80"))

	//create Gorilla router
	router := mux.NewRouter()

	//handle html template serving
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/{page}", pageHandler)
	router.HandleFunc("/submitform", formHandler).Methods("POST")

	//handle asset serving
	router.HandleFunc("/assets/{file}", assetHandler)

	//Create server on port :80
	log.Fatal(http.ListenAndServe(":80", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("home", r.URL.Path)
	renderTemplate(w, "index.html")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("page", r.URL.Path)
	path := r.URL.Path
	if path == "/favicon.ico" {
		http.ServeFile(w, r, "favicon.ico")
		return
	}

	if path == "/submitform" && r.Method == "POST" {
		formHandler(w, r)
		return
	}

	tmpl, ok := pageTemplates[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, tmpl)
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	path := "src/assets/" + r.URL.Path[len("/assets/"):]
	contentType := getContentType(path)
	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, path)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	if name == "" || email == "" || message == "" {
		//send error message:
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<script>alert("Please fill in all fields.");window.history.back();</script>`)
		return
	}

	if err := mail.Mail(name, email, subject, message); err != nil {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<script>alert("Please use a valid email address.");window.history.back();</script>`)
		return
	}

	// Send response back to client
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<script>alert("Message sent successfully!");document.location.href="/";</script>`)
}

func getContentType(path string) string {
	switch filepath.Ext(path) {
	case ".css":
		return "text/css"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".ico":
		return "image/x-icon"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	//Parse template file
	t, err := template.ParseFiles(tmpl)
	check(err, w)
	//Execute Template
	err = t.Execute(w, nil)
	check(err, w)
}

func check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
