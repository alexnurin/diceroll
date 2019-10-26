package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed  = flag.Int("seed", -1, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "min random value")
	end   = flag.Int("end", 6, "max random value")
)

func randInterval(l, r int) int {
	return rand.Intn(r-l+1) + l
}

func main() {
	flag.Parse()
	if *seed == -1 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	fmt.Println(randInterval(*start, *end))
}
