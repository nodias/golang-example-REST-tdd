package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname, host, port string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", user, password, dbname, host, port)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	log.Println("## APP INITIALIZED")
}

func (a *App) InitializeRoute() {
	r := a.Router.PathPrefix("/").Subrouter()
	r.Path("/products/").HandlerFunc(a.getProducts).Methods("GET")
	r.Path("/product/{id:[0-9]+}").HandlerFunc(a.getProduct).Methods("GET")
	r.Path("/product/{id:[0-9]+}").HandlerFunc(a.updateProduct).Methods("UPDATE")
	r.Path("/product/").HandlerFunc(a.createProduct).Methods("POST")
	r.Path("/product/{id:[0-9]+}").HandlerFunc(a.deleteProduct).Methods("DELETE")
}

func (a *App) Run(addr string) {}

func (a *App) getProducts(w http.ResponseWriter, req *http.Request) {
}

func (a *App) getProduct(w http.ResponseWriter, req *http.Request) {
	var p product
	p.ID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := p.get(a.DB); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(Response{
		Products: []product{p},
		Err:      ResponseError{err},
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func (a *App) updateProduct(w http.ResponseWriter, req *http.Request) {

}

func (a *App) createProduct(w http.ResponseWriter, req *http.Request) {

}

func (a *App) deleteProduct(w http.ResponseWriter, req *http.Request) {

}
