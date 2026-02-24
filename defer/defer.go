// Defer: Adia a execução de uma função, até o fim da função atual
// Usado para Close(), Unlock(), defer fmt.Println()
// 👉 Regra mental: “executa no final”

package main

import "fmt"

func dayOne() {
	fmt.Println("Domingo")
}

func dayTwo() {
	fmt.Println(("Segunda-feira"))
}

func main() {
	defer dayOne()
	dayTwo()
}
