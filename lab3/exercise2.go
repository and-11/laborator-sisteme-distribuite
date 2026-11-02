package main

type LamportClock struct {
	v int
}

func (lc *LamportClock) Tick() {
	lc.v++
}

func (lc *LamportClock) Send() int {
	lc.v++
	return lc.v
}

func (lc *LamportClock) Receive(ts int) {
	if ts > lc.v {
		lc.v = ts
	}
	lc.v++
}

type Message struct {
	from int
	ts   int
}

type Event struct {
	pid  int
	tick int
	typ  string
}

func less(x, y Event) bool {
	if x.tick == y.tick {
		return x.pid < y.pid
	}
	return x.tick < y.tick
}