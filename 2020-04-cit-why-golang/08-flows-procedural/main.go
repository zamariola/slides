package main

import (
	"fmt"
	"net/http"
)
//PROCEDURES
func main() {

	r, _ := http.Get("http://www.google.com")
	fmt.Println(r)

	success := IsSuccess(r)
	fmt.Printf("Success: %t", success)
}

func IsSuccess(response *http.Response) bool {
	return response.StatusCode == 200
}
