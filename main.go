package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	re redis.Conn
)

func main() {
	http.Handle("/Login/", http.StripPrefix("/Login/", http.FileServer(http.Dir("./LoginFiles"))))
	http.HandleFunc("/Login/handleLogin/", handleLogin)

	http.Handle("/SignUp/", http.StripPrefix("/SignUp/", http.FileServer(http.Dir("./SignUpFiles"))))
	http.HandleFunc("/SignUp/handleSignUp/", handleSignup)

	http.HandleFunc("/", handleRoot)

	http.HandleFunc("/Welcome", handleWelcome)

	http.HandleFunc("/Logout", handleLogout)

	initDB()
	initRedis()
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable password="+os.Getenv("DBPass"))
	if err != nil {
		panic(err)
	}
}

func initRedis() {
	var err error
	re, err = redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
}
