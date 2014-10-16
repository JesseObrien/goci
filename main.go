package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
)

var (
	debugMode = flag.Bool("debug", false, "Debug switch.")
	port      = flag.String("port", ":8080", "Server port")
)

func init() {
	flag.Parse()
}

func main() {

	routes := NewRoutes()

	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.GetIndex).Methods("GET")
	mux.HandleFunc("/register", routes.GetRegister).Methods("GET")
	mux.HandleFunc("/register", routes.PostRegister).Methods("POST")
	mux.HandleFunc("/login", routes.Login).Methods("POST")
	mux.HandleFunc("/logout", routes.Logout).Methods("GET")
	mux.HandleFunc("/builds", routes.Builds).Methods("GET")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)
	n.Run(":8888")

}
