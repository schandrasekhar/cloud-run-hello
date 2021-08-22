package main

import (
    "os"
    "fmt"
    "net/http"
    "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
       port = "8080"
    }
    http.HandleFunc("/", HelloServer)
    // http.ListenAndServe(":8080", nil)
    log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
