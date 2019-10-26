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
	n     = flag.Int("n", 1, "count of dices")
)

func randInterval(l, r int) int {
	if l > r {
		fmt.Println("Введённый интервал некорректен")
		return -1
	} else {
		return rand.Intn(r-l+1) + l
	}
}

func main() {
	flag.Parse()
	if *seed == -1 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	for i := 0; i < *n; i++ {
		res := randInterval(*start, *end)
		if res != -1 {
			fmt.Printf("%d ", res)
		}
	}
}
