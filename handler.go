package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	uname, pswd := r.URL.Query().Get("u"), r.URL.Query().Get("p")

	result := db.QueryRow("select password from users where username=$1", uname)
	c := &credentials{}

	err := result.Scan(&c.pswd)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintln(w, "Incorrect login details")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(c.pswd), []byte(pswd)); err != nil {
		fmt.Fprintln(w, "Incorrect login details")
		return
	}

	fmt.Fprintln(w, "Login Worked")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	uname, pswd := r.URL.Query().Get("u"), r.URL.Query().Get("p")

	pswdHash, err := bcrypt.GenerateFromPassword([]byte(pswd), 8)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.Query("insert into users values ($1, $2)", uname, string(pswdHash))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Signup Done")

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./index")).ServeHTTP(w, r)
}
