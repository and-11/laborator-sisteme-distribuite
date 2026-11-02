package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

func estimate(t1, t2, t3, t4 time.Time) (time.Duration, time.Duration) {
	t := (t2.Sub(t1) + t3.Sub(t4)) / 2
	d := t4.Sub(t1) - t3.Sub(t2)
	return t, d
}

func measure(cli, srv *Clock) (time.Duration, time.Duration) {
	t1 := cli.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	
	t2 := srv.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	
	t3 := srv.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	
	t4 := cli.Now()

	return estimate(t1, t2, t3, t4)
}

func simulate(cli, srv *Clock, n int) (time.Duration, time.Duration) {
	var bestT, minD time.Duration

	for i := 0; i < n; i++ {
		t, d := measure(cli, srv)
		fmt.Printf("[%d] t: %v, d: %v\n", i, t, d)

		if i == 0 || d < minD {
			minD = d
			bestT = t
		}
	}

	return bestT, minD
}