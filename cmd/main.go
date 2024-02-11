package main

import (
	"log"
	"net/http"

	"github.com/LigeronAhill/hello-api/handlers/rest"
)

func main() {
	addr := ":8000"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rest.TranslateHandler)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
