package main

import (
    "net/http"
	"fmt"
)

var (
	initShops []*ShopProfile
)

func init() {
	initShops = []*ShopProfile{
		&ShopProfile{1, "三只汪外卖", "北京", "西城区xxx号"},
		&ShopProfile{2, "三只汪外卖", "上海", "西城区xxx号"},
	}
}

func admin_login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello admin")
}

func admin_listshop(w http.ResponseWriter, r *http.Request) {
	var rspmsg AdminListShopRsp
	rspmsg.Shops = initShops
	doWriteJson(w, rspmsg)
}

func admin_addshop(w http.ResponseWriter, r *http.Request) {
}

func handle_admin() {
    http.HandleFunc("/admin/login", admin_login)

    http.HandleFunc("/admin/listshop", admin_listshop)

    http.HandleFunc("/admin/addshop", admin_addshop)
}
