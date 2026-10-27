package main

import "fmt"

func send(ch chan<- string) {
	ch <- "salut"
}

func receive(in <-chan string, out chan<- string) {
	mesaj := <-in
	out <- mesaj
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1)
	go receive(ch1, ch2)

	fmt.Println(<-ch2)
}