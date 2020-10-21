package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(NewServer(":8080", "cert.pem", "key.pem", 2*time.Second, 5*time.Second))
}

func NewServer(
	address string,
	certPath string,
	keyPath string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
) error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Web Server\n"))
	})

	srv := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return srv.ListenAndServeTLS(certPath, keyPath)
}
