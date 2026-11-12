package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Punct struct {
	X float64
	Y float64
}

func distanta(p1, p2 Punct) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

func worker(id int, pozitii []Punct, raza float64, startChan chan int, doneChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		cmd := <-startChan
		if cmd == 0 {
			return
		}

		sumaX := pozitii[id].X
		sumaY := pozitii[id].Y
		vecini := 0

		for i := 0; i < len(pozitii); i++ {
			if i != id {
				if distanta(pozitii[id], pozitii[i]) <= raza {
					sumaX += pozitii[i].X
					sumaY += pozitii[i].Y
					vecini++
				}
			}
		}

		newX := sumaX / float64(1+vecini)
		newY := sumaY / float64(1+vecini)

		doneChan <- 1

		cmd = <-startChan
		if cmd == 0 {
			return
		}

		pozitii[id].X = newX
		pozitii[id].Y = newY

		doneChan <- 1
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	tokens := strings.Fields(string(data))
	if len(tokens) == 0 {
		return
	}

	numarRoboti, _ := strconv.Atoi(tokens[0])
	var wg sync.WaitGroup

	pozitii := make([]Punct, numarRoboti)
	startChan := make([]chan int, numarRoboti)
	doneChan := make(chan int, numarRoboti)

	idx := 1
	for i := 0; i < numarRoboti; i++ {
		x, _ := strconv.ParseFloat(tokens[idx], 64)
		y, _ := strconv.ParseFloat(tokens[idx+1], 64)

		pozitii[i] = Punct{X: x, Y: y}
		startChan[i] = make(chan int)

		idx += 2
	}

	raza := 10.0

	for i := 0; i < numarRoboti; i++ {
		wg.Add(1)
		go worker(i, pozitii, raza, startChan[i], doneChan, &wg)
	}

	maxIter := 50

	for iter := 1; iter <= maxIter; iter++ {
		for i := 0; i < numarRoboti; i++ {
			startChan[i] <- 1
		}

		for i := 0; i < numarRoboti; i++ {
			<-doneChan
		}

		for i := 0; i < numarRoboti; i++ {
			startChan[i] <- 1
		}

		for i := 0; i < numarRoboti; i++ {
			<-doneChan
		}

		if(iter % 2 == 0){
			fmt.Println("\n\nPozitiile la pas %d:", iter)
			for i, p := range pozitii {
				fmt.Printf("Robot %d: (%.4f, %.4f)\n", i, p.X, p.Y)
			}
		}

	}

	for i := 0; i < numarRoboti; i++ {
		startChan[i] <- 0
	}

	wg.Wait()

	fmt.Println("Pozitiile finale:")
	for i, p := range pozitii {
		fmt.Printf("Robot %d: (%.4f, %.4f)\n", i, p.X, p.Y)
	}
}