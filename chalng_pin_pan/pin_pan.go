/*Desafio: Imprimir de 1 a 100
- Quando for múltiplo de 3, printar "Pin"
- Quando for múltiplo de 5, printar "Pan"*/

package main

import "fmt"

func main() {
	i := 0
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i, "- Pin")
		}
		if i%5 == 0 {
			fmt.Println(i, "- Pan")
		}
		i++
	}
}
