package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ex 1:")
	
	srv := &Clock{offset: 0}
	cli := &Clock{offset: 100 * time.Millisecond}

	t, d := measure(cli, srv)
	fmt.Printf("t: %v, d: %v\n", t, d)

	b_theta, m_delta := simulate(cli, srv, 10)
	fmt.Printf("best_t: %v, min_d: %v\n", b_theta, m_delta)

	fmt.Println("\nex 2:")
}