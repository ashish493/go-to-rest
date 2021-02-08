package app

import (

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


type app struct {

	Router *mux.Router
	DB *gorm.DB
}

func (a *app) setRouters(){

	//for routing
}

func (a *app) Get(path string, f func(w http.ResponseWriter, r *http.Request)){
	a.Router.Handlefunc(path,f).Methods("GET")
}

func (a *app) Post(path string, f func(w http.ResponseWriter, r *http.Request)){
	a.Router.Handlefunc(path,f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}