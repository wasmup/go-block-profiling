
all:
	go build -ldflags=-s -o /dev/shm/main
	time /dev/shm/main
	go tool pprof -text -output cpu.txt cpu.out
	go tool pprof -png -output cpu.png cpu.out

pprof:
	go test -bench=. -benchmem -memprofile mem.out -cpuprofile cpu.out -blockprofile block.out -mutexprofile mutex.out
	
	go tool pprof -text -output cpu.txt cpu.out
	go tool pprof -png -output cpu.png cpu.out

	go tool pprof -png -output mem.png mem.out
	go tool pprof -png -output block.png block.out
	go tool pprof -png -output mutex.png mutex.out

rmpprof:
	rm cpu.out
	rm cpu.png
	rm cpu.txt
	rm mem.out
	rm mem.png
	rm block.png 
	rm block.out
	rm mutex.png 
	rm mutex.out
