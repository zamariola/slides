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
	sv, _ := NewServer(":8080")
	//10 method calls (200 lines)
	sv.setTimeout(2*time.Second, 3*time.Second)
	//10 method calls (200 lines)
	log.Println(sv.Serve())
}
