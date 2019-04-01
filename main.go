package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	http.Handle("/Login/", http.StripPrefix("/Login/", http.FileServer(http.Dir("./LoginFiles"))))
	http.HandleFunc("/Login/handleLogin/", handleLogin)

	http.Handle("/SignUp/", http.StripPrefix("/SignUp/", http.FileServer(http.Dir("./SignUpFiles"))))
	http.HandleFunc("/SignUp/handleSignUp/", handleSignup)

	http.HandleFunc("/", handleRoot)

	initDB()
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable password="+os.Getenv("DBPass"))
	if err != nil {
		panic(err)
	}
}
