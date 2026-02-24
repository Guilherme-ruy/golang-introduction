//Closure: Função que chama variavél que está em outra função
// Ex: Contadores, encapsular comportamento

package main

import "fmt"

func main() {
	x := 0

	numero := func() int {
		x++
		return x
	}
	fmt.Println(numero())
}
