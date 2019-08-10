package main

import "fmt"

//type alias
type SomaFunc func(int, int) int

//High-Order: retorna uma função
func myFunction() SomaFunc {
	return func(a, b int) int {
		return a + b
	}
}

func main() {

	//first class citizen: atribuir a uma variável
	f := myFunction()

	//first class citizen: chamar através da variável
	soma := f(2, 3)

	fmt.Println(soma)

}
