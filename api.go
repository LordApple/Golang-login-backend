package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUser Makes a GET call to the redis db to retrive the current user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	u := getUser{}

	u.uname = fmt.Sprintf("\"username\": \"%s\"", res)
	json.Marshal(u)

	fmt.Fprint(w, u)
}
