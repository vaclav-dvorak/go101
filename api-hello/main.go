package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func returnHelloName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Println("Endpoint Hit: returnHelloName")
	json.NewEncoder(w).Encode(Response{Message: "Hello " + name})
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/hello/{name}", returnHelloName)
    log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
