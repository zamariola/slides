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

type Config struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Address      string
}

type ConfigBuilder struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Address      string
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (cb *ConfigBuilder) WithReadTimeOut(rto time.Duration) *ConfigBuilder {
	cb.ReadTimeout = rto
	return cb
}

func (cb *ConfigBuilder) WithWriteTimeOut(wto time.Duration) *ConfigBuilder {
	cb.WriteTimeout = wto
	return cb
}

func (cb *ConfigBuilder) WithAddress(address string) *ConfigBuilder {
	cb.Address = address
	return cb
}

func (cb *ConfigBuilder) Build() *Config {
	return &Config{cb.ReadTimeout, cb.WriteTimeout, cb.Address}
}

func NewServer(cfg Config) (*Server, error) {

	dftLtn, err := net.Listen("tcp", cfg.Address)

	return &Server{
		Ltn:          dftLtn,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}, err
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
	cfg := NewConfigBuilder().
		WithAddress(":8080").
		WithReadTimeOut(2 * time.Second).
		WithWriteTimeOut(3 * time.Second).
		Build()

	//10 method calls (200 lines)
	sv, _ := NewServer(*cfg)
	//10 method calls (200 lines)
	log.Println(sv.Serve())
}
