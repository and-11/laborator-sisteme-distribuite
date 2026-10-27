package main

import "fmt"

func main() {
	ch := make(chan string)
	done := make(chan bool)

	go func() {
		mesaj := <-ch
		fmt.Println(mesaj)
		done <- true
	}()

	ch <- "salut"
	<-done
}