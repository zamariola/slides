package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println(NewServer(":8080"))
}

func NewServer(address string) error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Web Server\n"))
	})

	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}
	return srv.ListenAndServe()
}
