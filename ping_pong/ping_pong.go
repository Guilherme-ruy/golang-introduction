/*
Desafio: Um programa que exiba, em alternância, as palavras ping e pong
*/

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" //usado "<-" para enviar e receber mensagem pelo chan
	}
}

func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}
