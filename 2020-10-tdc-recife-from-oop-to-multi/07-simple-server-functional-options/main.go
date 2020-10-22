package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

func WithRoutes() func(*Server) {

	return func(s *Server) {

		r := http.NewServeMux()

		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("My Web Server\n"))
		})

		s.Handler = r
	}
}

func WithTimeout(read, write time.Duration) func(*Server) {

	return func(s *Server) {
		s.ReadTimeout = read
		s.WriteTimeout = write
	}
}

func WithCertificate(cert, key string) func(*Server) {

	return func(s *Server) {
		cert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			panic(err)
		}
		config := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		//Close old listener
		if err := s.Ltn.Close(); err != nil {
			panic(err)
		}

		//Keep old address
		originalAddr := s.Ltn.Addr().String()
		ltn, err := tls.Listen("tcp", originalAddr, config)
		if err != nil {
			panic(err)
		}
		s.Ltn = ltn
	}
}

type Server struct {
	Ltn          net.Listener //address + certificates
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Handler      http.Handler
}

func NewServer(addr string) (*Server, error) {

	dftLtn, err := net.Listen("tcp", addr)

	return &Server{
		Ltn:          dftLtn,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
	}, err
}

func (s *Server) Serve(options ...func(*Server)) error {

	for _, opt := range options {
		opt(s)
	}

	srv := &http.Server{
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return srv.Serve(s.Ltn)
}

func main() {

	sv, err := NewServer(":8080")

	log.Println(sv.Serve(
		WithTimeout(2*time.Second, 3*time.Second),
		WithRoutes(),
		WithCertificate("cert.pem", "key.pem"),
	))
}
