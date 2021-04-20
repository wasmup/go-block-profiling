package main

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// var n int // We use TestMain to set up the done channel.
	// flag.IntVar(&n, "n", 1_000_000, "chan cap")
	// flag.Parse()
	runtime.SetBlockProfileRate(1)

	os.Exit(m.Run())
}
func BenchmarkShuffle0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffle0(a)
	}
}

func BenchmarkShuffle1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffle1(a)
	}
}

func BenchmarkShuffle2(b *testing.B) {
	var ds *my
	for i := 0; i < b.N; i++ {
		a = ds.shuffle2(a)
	}
}

func BenchmarkShuffle3(b *testing.B) {
	var ds *my
	for i := 0; i < b.N; i++ {
		a = ds.shuffle3(a)
	}
}

func BenchmarkShuffle4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go shuffle0(a)
		shuffle0(a)
	}
}

var a = []string{"a", "b", "c", "d"}
