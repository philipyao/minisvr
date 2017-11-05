package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func doWriteJson(w http.ResponseWriter, pkg interface{}) {
	data, err := json.Marshal(pkg)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
