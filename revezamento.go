package main

import (
	"fmt"
	"time"
)

// Na função main é iniciada a corrida de revezamento com quatro corredores
func main() {
	bastao := make(chan int)

	go corredor(bastao, 1)
	time.Sleep(time.Second)
	go corredor(bastao, 2)
	time.Sleep(time.Second)
	go corredor(bastao, 3)
	time.Sleep(time.Second)
	go corredor(bastao, 4)
	time.Sleep(time.Second)

	bastao <- 1

	<-bastao

	fmt.Printf("Fim de corrida\n")
}

// A função corredor define um corredor. Ele recebe um bastão e um número que o identifica. Quando o bastão é passado pra ele, ele inicia a sua corrida.
func corredor(bastao chan int, numero int) {
	
	<-bastao
	fmt.Printf("Corredor %d começou a correr\n", numero)
	time.Sleep(time.Second * 3)
	fmt.Printf("Corredor %d concluiu a corrida\n", numero)
	bastao <- 1
	if numero == 4 {
		close(bastao)
		return
	}
			
}