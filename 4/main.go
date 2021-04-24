package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

const max = 3_000

func main() {
	{
		runtime.SetBlockProfileRate(1)

		f, err := os.Create("cpu.out")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	a := make([]string, max)
	for i := range a {
		a[i] = strings.Repeat(string(rune('a'+rand.Intn(26))), 1+rand.Intn(10))
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	t0 := time.Now()
	for i := 0; i < max; i++ {
		shuffle(r, a)
	}
	fmt.Println(time.Since(t0))
}

func shuffle(r *rand.Rand, a []string) {
	r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
