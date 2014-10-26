package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
)

var (
	debugMode = flag.Bool("debug", false, "Debug switch.")
)

func init() {
	flag.Parse()
}

func main() {

	routes := NewRoutes()

	router := mux.NewRouter()

	router.HandleFunc("/", routes.GetIndex).Methods("GET")
	router.HandleFunc("/register", routes.GetRegister).Methods("GET")
	router.HandleFunc("/register", routes.PostRegister).Methods("POST")
	router.HandleFunc("/login", routes.Login).Methods("POST")
	router.HandleFunc("/logout", routes.Logout).Methods("GET")

	router.HandleFunc("/builds", routes.BuildsJson).Methods("GET")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(router)
	n.Run(":8888")

}
