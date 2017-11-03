package main

import (
    "net/http"
)


func admin_login(w http.ResponseWriter, r *http.Request) {
}

func admin_listshop(w http.ResponseWriter, r *http.Request) {
}

func admin_addshop(w http.ResponseWriter, r *http.Request) {
}

func handle_admin() {
    http.HandleFunc("/admin/login", admin_login)

    http.HandleFunc("/admin/listshop", admin_listshop)

    http.HandleFunc("/admin/addshop", admin_addshop)
}
