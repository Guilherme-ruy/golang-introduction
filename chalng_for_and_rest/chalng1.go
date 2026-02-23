/*Desafio 1: Criar um código que exiba todos os números de 1 a 100
que são divisíveis por 3*/

package main

import "fmt"

func main() {
	i := 0
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println("Esse número é divisível por 3: ", i)
		}
		i = i + 1
	}
}
