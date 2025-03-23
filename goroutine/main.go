package main

import (
	"fmt"
	"time"
)

func goroutine(stop chan bool) {
	for {
		select {
		//Menerima sinyal stop dari channel
		case <-stop:
			fmt.Println("Goroutine dihentikan")
			return
		default:
			fmt.Println("Goroutine berjalan...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go goroutine(stop)

	time.Sleep(5 * time.Second)
	stop <- true // Mengirim sinyal untuk menghentikan goroutine

	time.Sleep(1 * time.Second) // Menunggu goroutine selesai
	fmt.Println("Program selesai")
}
