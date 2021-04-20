
## Predefined profiles provided by the runtime/pprof package


https://golang.org/doc/diagnostics

cpu: CPU profile determines where a program spends its time while actively consuming CPU cycles (as opposed to while sleeping or waiting for I/O).

heap: Heap profile reports memory allocation samples; used to monitor current and historical memory usage, and to check for memory leaks.
threadcreate: Thread creation profile reports the sections of the program that lead the creation of new OS threads.
goroutine: Goroutine profile reports the stack traces of all current goroutines.

block: Block profile shows where goroutines block waiting on synchronization primitives (including timer channels). Block profile is not enabled by default; use runtime.SetBlockProfileRate to enable it.

mutex: Mutex profile reports the lock contentions. When you think your CPU is not fully utilized due to a mutex contention, use this profile. Mutex profile is not enabled by default, see runtime.SetMutexProfileFraction to enable it.


```sh
make
# BenchmarkShuffle0-8     13761372                87.82 ns/op            0 B/op          0 allocs/op
# BenchmarkShuffle1-8        86557             12098 ns/op               0 B/op          0 allocs/op
# BenchmarkShuffle2-8        72139             22491 ns/op            5376 B/op          1 allocs/op
# BenchmarkShuffle3-8        56299             23308 ns/op            5440 B/op          2 allocs/op
# BenchmarkShuffle4-8      1000000              4180 ns/op               2 B/op          0 allocs/op
```

https://github.com/google/pprof/blob/master/doc/README.md#interpreting-the-callgraph


```sh
File: go-block-profiling.test
Type: cpu
Time: Apr 20, 2021 at 3:09am (MST)
Duration: 11.13s, Total samples = 14.37s (129.10%)
Showing nodes accounting for 12.87s, 89.56% of 14.37s total
Dropped 142 nodes (cum <= 0.07s)
      flat  flat%   sum%        cum   cum%
     3.41s 23.73% 23.73%      3.41s 23.73%  math/rand.seedrand
     1.46s 10.16% 33.89%      1.46s 10.16%  runtime.futex
     0.84s  5.85% 39.74%      6.09s 42.38%  math/rand.(*lockedSource).Int63
     0.83s  5.78% 45.51%      0.88s  6.12%  runtime.step
     0.82s  5.71% 51.22%      3.25s 22.62%  sync.(*Mutex).Unlock (inline)
     0.78s  5.43% 56.65%      4.19s 29.16%  math/rand.(*rngSource).Seed
     0.69s  4.80% 61.45%      1.77s 12.32%  runtime.pcvalue
     0.41s  2.85% 64.30%      2.47s 17.19%  runtime.gentraceback
     0.34s  2.37% 66.67%      0.34s  2.37%  go-block-profiling.shuffle0.func1
     0.30s  2.09% 68.75%      1.76s 12.25%  sync.(*Mutex).lockSlow
     0.25s  1.74% 70.49%      1.65s 11.48%  runtime.findrunnable
     0.23s  1.60% 72.09%      6.88s 47.88%  math/rand.(*Rand).Shuffle
     0.21s  1.46% 73.56%      0.21s  1.46%  math/rand.(*rngSource).Uint64 (inline)
     0.19s  1.32% 74.88%      0.19s  1.32%  runtime.findfunc
     0.17s  1.18% 76.06%      0.17s  1.18%  runtime.fastrand (inline)
     0.14s  0.97% 77.04%      0.14s  0.97%  runtime.memclrNoHeapPointers
     0.14s  0.97% 78.01%      0.14s  0.97%  runtime.procyield
     0.13s   0.9% 78.91%      6.31s 43.91%  math/rand.(*Rand).int31n
     0.13s   0.9% 79.82%      0.13s   0.9%  runtime.casgstatus
     0.12s  0.84% 80.65%      0.20s  1.39%  runtime.lock2
     0.09s  0.63% 81.28%      6.18s 43.01%  math/rand.(*Rand).Int63 (inline)
     0.09s  0.63% 81.91%      0.09s  0.63%  runtime.(*gList).pop (inline)
     0.08s  0.56% 82.46%      0.16s  1.11%  runtime.runqgrab
     0.08s  0.56% 83.02%      0.08s  0.56%  runtime.usleep
     0.07s  0.49% 83.51%      4.03s 28.04%  runtime.systemstack
     0.06s  0.42% 83.92%      6.93s 48.23%  go-block-profiling.shuffle0
     0.06s  0.42% 84.34%      2.35s 16.35%  runtime.semrelease1
     0.06s  0.42% 84.76%      0.93s  6.47%  runtime.wakep
     0.06s  0.42% 85.18%      2.43s 16.91%  sync.(*Mutex).unlockSlow
     0.05s  0.35% 85.53%      0.14s  0.97%  runtime.gfget
     0.05s  0.35% 85.87%      2.18s 15.17%  runtime.mcall
     0.04s  0.28% 86.15%      0.78s  5.43%  runtime.funcspdelta
     0.04s  0.28% 86.43%      0.38s  2.64%  runtime.newproc1
     0.04s  0.28% 86.71%      2.64s 18.37%  runtime.saveblockevent
     0.03s  0.21% 86.92%      0.64s  4.45%  runtime.notesleep
     0.03s  0.21% 87.13%      0.80s  5.57%  runtime.ready
     0.03s  0.21% 87.33%      0.09s  0.63%  runtime.runqput
     0.03s  0.21% 87.54%      0.19s  1.32%  runtime.runqsteal
     0.03s  0.21% 87.75%      1.96s 13.64%  runtime.schedule
     0.03s  0.21% 87.96%      1.37s  9.53%  runtime.semacquire1
     0.03s  0.21% 88.17%      1.79s 12.46%  sync.(*Mutex).Lock
     0.02s  0.14% 88.31%      2.24s 15.59%  go-block-profiling.BenchmarkShuffle0
     0.02s  0.14% 88.45%      1.15s  8.00%  go-block-profiling.BenchmarkShuffle4
     0.02s  0.14% 88.59%      0.21s  1.46%  runtime.isSystemGoroutine
     0.02s  0.14% 88.73%      0.87s  6.05%  runtime.notewakeup
     0.02s  0.14% 88.87%      0.76s  5.29%  runtime.stopm
     0.02s  0.14% 89.00%      2.37s 16.49%  sync.runtime_Semrelease
     0.01s  0.07% 89.07%      0.18s  1.25%  runtime.(*mcache).refill
     0.01s  0.07% 89.14%      2.49s 17.33%  runtime.callers
     0.01s  0.07% 89.21%      2.48s 17.26%  runtime.callers.func1
     0.01s  0.07% 89.28%      0.62s  4.31%  runtime.futexsleep
     0.01s  0.07% 89.35%      1.34s  9.32%  runtime.goexit0
     0.01s  0.07% 89.42%      0.81s  5.64%  runtime.goready
     0.01s  0.07% 89.49%      0.08s  0.56%  runtime.gostringnocopy (inline)
     0.01s  0.07% 89.56%      1.04s  7.24%  runtime.pcdatavalue
         0     0% 89.56%      1.78s 12.39%  go-block-profiling.(*my).shuffle2
         0     0% 89.56%      1.49s 10.37%  go-block-profiling.(*my).shuffle3
         0     0% 89.56%      1.19s  8.28%  go-block-profiling.BenchmarkShuffle1
         0     0% 89.56%      1.78s 12.39%  go-block-profiling.BenchmarkShuffle2
         0     0% 89.56%      1.49s 10.37%  go-block-profiling.BenchmarkShuffle3
         0     0% 89.56%      1.19s  8.28%  go-block-profiling.shuffle1
         0     0% 89.56%      1.19s  8.28%  math/rand.(*Rand).Seed
         0     0% 89.56%      6.18s 43.01%  math/rand.(*Rand).Uint32 (inline)
         0     0% 89.56%      1.19s  8.28%  math/rand.(*lockedSource).seedPos
         0     0% 89.56%      0.21s  1.46%  math/rand.(*rngSource).Int63 (inline)
         0     0% 89.56%      3.23s 22.48%  math/rand.NewSource (inline)
         0     0% 89.56%      1.19s  8.28%  math/rand.Seed (inline)
         0     0% 89.56%      6.87s 47.81%  math/rand.Shuffle (inline)
         0     0% 89.56%      0.18s  1.25%  runtime.(*mcache).nextFree
         0     0% 89.56%      0.17s  1.18%  runtime.(*mcentral).cacheSpan
         0     0% 89.56%      0.17s  1.18%  runtime.(*mcentral).grow
         0     0% 89.56%      0.16s  1.11%  runtime.(*mheap).alloc
         0     0% 89.56%      0.16s  1.11%  runtime._System
         0     0% 89.56%      1.23s  8.56%  runtime.blockevent
         0     0% 89.56%      0.08s  0.56%  runtime.funcname
         0     0% 89.56%      0.85s  5.92%  runtime.futexwakeup
         0     0% 89.56%      0.11s  0.77%  runtime.gcBgMarkWorker
         0     0% 89.56%      0.09s  0.63%  runtime.gcBgMarkWorker.func2
         0     0% 89.56%      0.09s  0.63%  runtime.gcDrain
         0     0% 89.56%      0.80s  5.57%  runtime.goready.func1
         0     0% 89.56%      0.14s  0.97%  runtime.lock (inline)
         0     0% 89.56%      0.20s  1.39%  runtime.lockWithRank (inline)
         0     0% 89.56%      0.69s  4.80%  runtime.mPark
         0     0% 89.56%      0.25s  1.74%  runtime.mallocgc
         0     0% 89.56%      0.23s  1.60%  runtime.newobject
         0     0% 89.56%      0.44s  3.06%  runtime.newproc
         0     0% 89.56%      0.44s  3.06%  runtime.newproc.func1
         0     0% 89.56%      0.72s  5.01%  runtime.park_m
         0     0% 89.56%      0.82s  5.71%  runtime.readyWithTime
         0     0% 89.56%      0.21s  1.46%  runtime.resetspinning
         0     0% 89.56%      0.88s  6.12%  runtime.startm
         0     0% 89.56%      1.43s  9.95%  sync.event
         0     0% 89.56%      1.37s  9.53%  sync.runtime_SemacquireMutex
         0     0% 89.56%      7.86s 54.70%  testing.(*B).launch
         0     0% 89.56%      7.86s 54.70%  testing.(*B).runN
```


<img src="cpu.png">
<img src="mem.png">
<img src="block.png">
<img src="mutex.png">
