package main

import (
	"log"
	"net/http"
	"time"
)

type Config struct {
	Address      string
	CertPath     string
	KeyPath      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func main() {

	config := Config{":8080", "cert.pem", "key.pem", 2 * time.Second, 5 * time.Second}
	log.Println(NewServer(config))
}

func NewServer(c Config) error {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Web Server\n"))
	})

	srv := &http.Server{
		Addr:         c.Address,
		Handler:      r,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
	return srv.ListenAndServeTLS(c.CertPath, c.KeyPath)
}
