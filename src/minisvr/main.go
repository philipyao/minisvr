package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func wssEcho(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Print("upgrade:", err)
        return
    }
    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            fmt.Println("read:", err)
            break
        }
        fmt.Printf("recv: %s", message)
        err = c.WriteMessage(mt, message)
        if err != nil {
            fmt.Println("write:", err)
            break
        }
    }
}

func main() {.0000
    fmt.Println("start mini server...")

    http.HandleFunc("/wssEcho", wssEcho)

    //处理用户点餐
    handle_mini()

    //点餐商户端
    handle_admin()

    if err := http.ListenAndServeTLS(":443", "./conf/214316622370391.crt", "./conf/214316622370391.key", nil); err != nil {
        panic("ListenAndServeTLS Error: " + err.Error())
    }
}

