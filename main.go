package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World! This is %s.\n", name)
    })

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Healthy!")
    })

    http.ListenAndServe(":8888", nil)
}
