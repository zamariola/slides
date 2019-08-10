package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(NewServer(":8080"))
}

func NewServer(address string) error {
	r := http.NewServeMux()

	r.HandleFunc("/", Authenticate(TraceRequest(GetMainResponse())))

	srv := &http.Server{
		Addr:    address,
		Handler: r,
	}
	return srv.ListenAndServe()
}

func GetMainResponse() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Web Server\n"))
	}
}

func TraceRequest(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Start request: %s", time.Now())
		f(w, r)
		log.Printf("End request: %s", time.Now())
	}
}

func Authenticate(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 { //Check token
			w.WriteHeader(401)
		} else {
			f(w, r)
		}
	}
}
