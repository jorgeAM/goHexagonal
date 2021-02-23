package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server is running ...")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hola"))
	})

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal("somenthing got wrong with your server")
	}
}
