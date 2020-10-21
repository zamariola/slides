package main

import "net/http"

func main() {

	//first class citizen: atribuir a uma variável
	f := handlerFunction()

	http.HandleFunc("/", f)
	http.ListenAndServe(":8080", nil)
}

//High-Order: retorna uma função
func handlerFunction() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Content\n"))
	}
}
