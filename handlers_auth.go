package main

import (
	"github.com/xyproto/permissions"
	"net/http"
)

type AuthHandler struct {
	userstate *permissions.UserState
}

func (ah *AuthHandler) Register(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")

	if ok := ah.userstate.HasUser(username); ok != false {
		http.Redirect(w, req, "/register", http.StatusFound)
	}

	if ok := permissions.ValidUsernamePassword(username, password); ok != nil {

	}

	ah.userstate.AddUser(username, password, email)

	http.Redirect(w, req, "/", http.StatusFound)
}

func (ah *AuthHandler) Login(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	password := req.FormValue("password")

	if ok := ah.userstate.CorrectPassword(username, password); ok != false {

	}

	ah.userstate.Login(w, username)

	http.Redirect(w, req, "/", http.StatusFound)
}

func (ah *AuthHandler) Logout(w http.ResponseWriter, req *http.Request) {
	ah.userstate.Logout(ah.userstate.GetUsername(req))

	http.Redirect(w, req, "/", http.StatusFound)
}
