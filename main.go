package main

import (
        "fmt"
        "net/http"
        "os"
)

func main() {
        port := "8080"
        http.HandleFunc("/", handler)
        http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
        hostname, err := os.Hostname()
        if err != nil {
                panic(err)
        }
        fmt.Fprintf(w, "<h1>Hostname : %s</h1>\n", hostname)
        fmt.Fprintf(w, "<h1>GitOps !!!</h1>\n")
}
