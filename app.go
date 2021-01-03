
package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"os"
)

func sayHello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(res, `<h1>Hello You!</h1>`)
}

func main() {

	checkIfEnvExist := godotenv.Load()
	if checkIfEnvExist != nil {
    log.Fatal("Error loading .env file")
	}
	

	router := mux.NewRouter()

	port := os.Getenv("PORT")

	log.Println("server started on port", port)

	router.HandleFunc("/", sayHello).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}