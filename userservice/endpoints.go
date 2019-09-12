package userservice

import (
	"fmt"
	"log"
	"net/http"
	"scratch/mockExample/util"

	"github.com/gorilla/mux"
)

// HandleRequests ...
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/{name}", GetEndpoint).Methods("GET")
	myRouter.HandleFunc("/user/{name}", DeleteEndpoint).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", CreateUpdateEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// NewStore ...
func NewStore() *Store {

	s := Store{}

	return &s
}

// CreateUpdateEndpoint ...
func CreateUpdateEndpoint(w http.ResponseWriter, r *http.Request) {

	myservice, err := NewUserService()
	if err != nil {
		fmt.Println(err)
		return
	}

	mystore := NewStore()

	vars := mux.Vars(r)

	myRequest := util.User{
		Name:  vars["name"],
		Email: vars["email"],
	}

	code, msg, err := myservice.CreateUpdateService(myRequest, mystore)
	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(code)
	w.Write([]byte(msg))

}

// GetEndpoint ...
func GetEndpoint(w http.ResponseWriter, r *http.Request) {

	myservice, err := NewUserService()
	if err != nil {
		fmt.Println(err)
		return
	}

	mystore := NewStore()

	vars := mux.Vars(r)

	myRequest := util.User{
		Name: vars["name"],
	}

	code, msg, err := myservice.GetService(&myRequest, mystore)
	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(code)
	w.Write([]byte(msg))

}

// DeleteEndpoint ...
func DeleteEndpoint(w http.ResponseWriter, r *http.Request) {

	myservice, err := NewUserService()
	if err != nil {
		fmt.Println(err)
		return
	}

	mystore := NewStore()

	vars := mux.Vars(r)

	myRequest := util.User{
		Name: vars["name"],
	}

	code, msg, err := myservice.DeleteService(&myRequest, mystore)
	if err != nil {
		w.WriteHeader(code)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(code)
	w.Write([]byte(msg))

}
