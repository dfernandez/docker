package main

import (
    "fmt"
    "net/http"
    redigo "github.com/garyburd/redigo/redis"
)

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/redis", redis)
    http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Go")
}

func redis(w http.ResponseWriter, r *http.Request) {
    c, err := redigo.Dial("tcp", "redis:6379")

    if err != nil {
        fmt.Fprint(w, err)
        return
    }
    defer c.Close()

    _, err = c.Do("PING")
    if err != nil {
        fmt.Fprint(w, err)
        return
    }

    fmt.Fprint(w, "PONG")
}
