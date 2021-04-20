package main

import (
	"math/rand"
	"time"
)

func main() {
}

func shuffle0(a []string) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func shuffle1(a []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func (*my) shuffle2(a []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func (*my) shuffle3(a []string) []string {
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
