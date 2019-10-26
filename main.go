package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	seed     = flag.Int("seed", -1, "seed for random generator. unix(now) be default")
	start    = flag.Int("start", 1, "min random value")
	end      = flag.Int("end", 6, "max random value")
	n        = flag.Int("n", 1, "count of dices")
	norepeat = flag.Bool("norepeat", false, "unique values")
)

func randInterval(l, r int) int {
	return rand.Intn(r-l+1) + l
}

func check_Errors() int {
	if *start > *end {
		fmt.Printf("Введённый отрезок [%d - %d] некорректен\n", *start, *end)
		return 1
	}
	if *norepeat && *end-*start+1 < *n {
		fmt.Printf("Невозможно вывести %d уникальных элементов на отрезке [%d - %d]", *n, *start, *end)
		return 2
	}
	return 0
}

func main() {
	flag.Parse()
	if *seed == -1 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	er := check_Errors()
	if er != 0 {
		os.Exit(0)
	}
	set := make(map[int]bool)
	for i := 0; i < *n; i++ {
		value := randInterval(*start, *end)
		if *norepeat {
			for set[value] {
				value = randInterval(*start, *end)
			}
		}
		fmt.Printf("%d ", value)
		set[value] = true
	}
	fmt.Println()
}
