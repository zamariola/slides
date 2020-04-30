package main

import (
	"fmt"
	"net/http"
)

//ENTITIES
type Response http.Response

func main() {

	r, _ := http.Get("http://www.google.com")
	fmt.Println(r)

	response := Response(*r)

	fmt.Printf("Success: %t", response.isSuccess())
}

func (r *Response) isSuccess() bool {
	return r.StatusCode == 200
}
