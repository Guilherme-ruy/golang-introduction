//Defer: Escalona as funções

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
