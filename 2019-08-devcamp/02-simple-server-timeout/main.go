package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(NewServer(":8080", 2*time.Second, 5*time.Second))
}

func NewServer(address string, rt, wt time.Duration) error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Web Server\n"))
	})

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"success\":true}"))
	})

	srv := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}
	return srv.ListenAndServe()
}
