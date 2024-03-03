module github.com/COACODE-LLC/coacode_website

go 1.18

require (
	func/mail v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/savioxavier/termlink v1.3.0
)

require github.com/joho/godotenv v1.5.1 // indirect

replace func/mail => ./src/scripts/mail/
