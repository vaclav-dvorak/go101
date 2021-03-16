package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    _ "github.com/vaclav-dvorak/go101/api-hello/docs" // This line is necessary for go-swagger to find your docs!
)

type Response struct {
    Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request){
    // swagger:operation GET /hello/{name} hello Hello
    //
    // Returns a simple Hello message
    // ---
    // consumes:
    // - text/plain
    // produces:
    // - application/json
    // parameters:
    // - name: name
    //   in: path
    //   description: Name to be returned.
    //   required: true
    //   type: string
    // responses:
    //   '200':
    //     description: The hello message
    //     type: string

    vars := mux.Vars(r)
    name := vars["name"]

    fmt.Println("Endpoint Hit: hello")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Response{Message: "Hello " + name})
}

func handleRequests() {
    // creates a new instance of a mux router
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/hello/{name}", hello).Methods("GET")

    handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}
