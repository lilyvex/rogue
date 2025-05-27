package main

import (
    "fmt"
    "net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "It's working!")
}

func main() {
    http.HandleFunc("/", root)

    http.ListenAndServe(":8080", nil)
}
