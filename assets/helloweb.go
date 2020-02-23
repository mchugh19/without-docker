package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", BasicServer)
        http.ListenAndServe(":88", nil)
}

func BasicServer(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "Hey, %s!", req.URL.Path[1:])
}
