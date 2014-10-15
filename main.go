package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/xyproto/permissions"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

var (
	debugMode = flag.Bool("debug", false, "Debug switch.")
)

func init() {
	flag.Parse()
}

func getUserState() *permissions.UserState {
	perm := permissions.New()

	return perm.UserState()
}

func main() {
	userstate := getUserState()

	authHandler := AuthHandler{userstate: userstate}

	mux := mux.NewRouter()

	r := render.New(render.Options{})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if userstate.IsLoggedIn(userstate.GetUsername(req)) {
			r.HTML(w, http.StatusOK, "dashboard", nil)
		} else {
			r.HTML(w, http.StatusOK, "landing", nil)
		}
	})

	mux.HandleFunc("/register", authHandler.Register).Methods("POST")
	mux.HandleFunc("/login", authHandler.Login).Methods("POST")
	mux.HandleFunc("/logout", authHandler.Logout).Methods("GET")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)
	n.Run(":8080")

}
