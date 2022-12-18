package main

import (
	"fmt"
	"time"
)

func main() {
	bastao := make(chan int)

	go corredor(bastao, 1)
	go corredor(bastao, 2)
	go corredor(bastao, 3)
	go corredor(bastao, 4)
	bastao <- 1

	<-bastao
}

func corredor(bastao chan int, numero int) {
	for {
		<-bastao
		fmt.Printf("Corredor %d começou a correr\n", numero)
		time.Sleep(time.Second)
		fmt.Printf("Corredor %d concluiu a corrida\n", numero)

		bastao <- 1

		if numero == 4 {
			close(bastao)
			return
		}
	}
}