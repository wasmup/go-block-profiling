
all:
	go test -bench=. -benchmem -memprofile mem.out -cpuprofile cpu.out -blockprofile block.out -mutexprofile mutex.out
	go tool pprof -png -output cpu.png cpu.out
	go tool pprof -png -output mem.png mem.out
	go tool pprof -png -output block.png block.out
	go tool pprof -png -output mutex.png mutex.out
    
    # will print a list of the hottest functions.
	go tool pprof --text cpu.out

