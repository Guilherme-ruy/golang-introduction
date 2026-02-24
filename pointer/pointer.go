/* Ponteiro:
- Variável guarda valor
- Ponteiro guarda endereço
- "*" acessa o valor do endereço
- "&" pega o endereço
*/

package main

import "fmt"

func inicial(x *int) {
	*x = 0
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}
