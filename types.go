package main

type credentials struct {
	pswd string `db:"password"`
}

type getUser struct {
	uname string
}
