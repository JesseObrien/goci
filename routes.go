package main

import (
	"github.com/mholt/binding"
	"github.com/xyproto/permissions"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"net/url"
)

type Routes struct {
	userstate permissions.UserState
	renderer  render.Render
}

func (r *Routes) LayoutRender(w http.ResponseWriter, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {

	htmlOpt = append(htmlOpt, render.HTMLOptions{Layout: "layout"})

	r.renderer.HTML(w, status, name, binding, htmlOpt...)
}

func (r *Routes) Render(w http.ResponseWriter, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {
	r.renderer.HTML(w, status, name, binding, htmlOpt...)
}

func NewRoutes() *Routes {

	r := render.New(render.Options{})
	perm := permissions.New()
	us := perm.UserState()

	return &Routes{userstate: *us, renderer: *r}
}

func (r *Routes) GetIndex(w http.ResponseWriter, req *http.Request) {
	r.LayoutRender(w, http.StatusOK, "layout", nil)
}

func (r *Routes) GetRegister(w http.ResponseWriter, req *http.Request) {
	r.LayoutRender(w, http.StatusOK, "registration", nil)
}

func (r *Routes) PostRegister(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")

	if ok := r.userstate.HasUser(username); ok != false {
		http.Redirect(w, req, "/register", http.StatusFound)
	}

	if ok := permissions.ValidUsernamePassword(username, password); ok != nil {

	}

	r.userstate.AddUser(username, password, email)

	http.Redirect(w, req, "/", http.StatusFound)
}

func (r *Routes) Login(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	password := req.FormValue("password")

	if ok := r.userstate.CorrectPassword(username, password); ok != false {

	}

	// r.userstate.Login(w, username)

	http.Redirect(w, req, "/builds", http.StatusFound)
}

func (r *Routes) Logout(w http.ResponseWriter, req *http.Request) {
	r.userstate.Logout(r.userstate.GetUsername(req))

	http.Redirect(w, req, "/", http.StatusFound)
}

type Build struct {
	Id          int             `json: "id"`
	Name        string          `json: "name"`
	Description string          `json: "description"`
	LastTen     map[string]bool `json: "LastTen"`
	VcType      string          `json: "vc_type"`
	VcURL       url.URL
}

func (b *Build) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&b.Name:        "name",
		&b.Description: "description",
		&b.VcType:      "vc_type",
		&b.VcURL:       "vc_url",
	}
}

func (r *Routes) ListBuilds(w http.ResponseWriter, req *http.Request) {

	b := []Build{
		Build{Id: 1, Name: "Beta", LastTen: map[string]bool{"1": true, "2": false, "3": true, "4": true, "5": true, "6": false, "7": true, "8": false, "9": true, "10": true}},
		Build{Id: 2, Name: "Alpha", LastTen: map[string]bool{"1": true, "2": false, "3": true, "4": true, "5": true, "6": false, "7": true, "8": false, "9": true, "10": true}},
	}

	r.renderer.JSON(w, http.StatusOK, b)
}

func (r *Routes) CreateBuild(w http.ResponseWriter, req *http.Request) {
	// @TODO figure out why this isn't getting vars from angular
	newBuild := new(Build)
	if errs := binding.Bind(req, newBuild); errs != nil {
		panic(errs)
	}

	r.renderer.JSON(w, http.StatusOK, newBuild)
}
