package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	bastao := make(chan int)
	
	var wg sync.WaitGroup
	wg.Add(4)

	go corredor(bastao, 1, &wg)
	go corredor(bastao, 2, &wg)
	go corredor(bastao, 3, &wg)
	go corredor(bastao, 4, &wg)

	wg.Wait()
	fmt.Printf("Fim de corrida\n")

}

func corredor(bastao chan int, numero int, wg *sync.WaitGroup) {

	bastao <- 1

	for {
		vez := <-bastao

		if vez == numero {
			fmt.Printf("Corredor %d começou a correr\n", numero)
			time.Sleep(time.Second * 1)
			fmt.Printf("Corredor %d concluiu a corrida\n", numero)
				
			bastao <- numero + 1

			if vez == 4 {
				close(bastao)
			}
			
			wg.Done()
			return

		}

	}
	
}