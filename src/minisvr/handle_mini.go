package main

import (
    "fmt"
    "net/http"
)


func mini_login(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello mini")
}

func mini_listshop(w http.ResponseWriter, r *http.Request) {
}

func mini_shopinfo(w http.ResponseWriter, r *http.Request) {
}

func handle_mini() {
    http.HandleFunc("/api/login", mini_login)

    http.HandleFunc("/api/listshop", mini_listshop)

    http.HandleFunc("/api/shopinfo", mini_shopinfo)
}
