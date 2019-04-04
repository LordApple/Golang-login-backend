package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

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

	Token, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := re.Do("SETEX", Token.String(), "31557600", uname); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   Token.String(),
		Expires: time.Now().Add(8766 * time.Hour),
		Path:    "/",
	})

	http.Redirect(w, r, "/Welcome", http.StatusSeeOther)
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

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Token := cookie.Value
	res, err := re.Do("GET", Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if res == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s", res)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = re.Do("DEL", cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./index")).ServeHTTP(w, r)
}
