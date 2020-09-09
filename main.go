package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"os"
)

func version(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Version is: " + os.Getenv("VERSION"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/version", version)
	log.Fatal(http.ListenAndServe(":3000", router))
}
