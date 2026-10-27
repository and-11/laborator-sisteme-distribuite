package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "mesaj 1"
		ch <- "mesaj 2"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}