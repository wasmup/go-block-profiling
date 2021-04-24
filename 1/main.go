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
	var m *my
	t0 := time.Now()
	for i := 0; i < max; i++ {
		a = m.shuffle(a)
	}
	fmt.Println(time.Since(t0))
}

func (*my) shuffle(a []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	result := make([]string, len(a))
	n := len(a)
	for i := 0; i < n; i++ {
		j := r.Intn(len(a))
		result[i] = a[j]
		a = append(a[:j], a[j+1:]...)
	}
	return result
}

type my struct{}
