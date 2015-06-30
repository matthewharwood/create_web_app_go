package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println(r)
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8000", nil)
}
