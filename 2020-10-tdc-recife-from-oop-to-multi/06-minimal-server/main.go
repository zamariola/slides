package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

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

func (s *Server) setTimeout(read, write time.Duration) {
	s.ReadTimeout = read
	s.WriteTimeout = write
}

func (s *Server) Serve() error {

	srv := &http.Server{
		Handler:      s.Handler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return srv.Serve(s.Ltn)
}

func main() {

	sv, err := NewServer(":8080")

	if err != nil {
		panic(err)
	}

	sv.setTimeout(2*time.Second, 3*time.Second)

	log.Println(sv.Serve())
}
