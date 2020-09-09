package main

import (
	"fmt"
	"net/http"
	"os"
)

func version(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Version is: " + os.Getenv("VERSION") + "\n")
}

func main() {
    http.HandleFunc("/version", version)
    http.ListenAndServe(":3000", nil)
}
