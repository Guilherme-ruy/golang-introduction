package main

import "fmt"

//variavel const do ponto de ebulição da água em fahrenheit
const ebulicaoF float64 = 212

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9
	fmt.Println("Temperatura de ebulição da água em Fahrenheit = ", tempF)
	fmt.Println("Temperatura de ebulição da água em Celsius é = ", tempC)
}
