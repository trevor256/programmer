package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Listening on http://localhost:2566/index.html")
    err := http.ListenAndServe(":2566", nil)
    if err != nil {
        log.Fatal(err)
    }
}
