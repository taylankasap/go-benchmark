# What

> In Go 1.22, each iteration of the loop creates new variables

As soon as I saw this statement I wanted to check the performance impact.

Source: https://tip.golang.org/doc/go1.22#language

# Benchmark Result

In an empty loop scenario 1.22rc1 is 2x "worse" than 1.21.5.

1.21.5:

```bash
$ go test -bench=. -count 5
goos: linux
goarch: amd64
pkg: loop-1.21
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkForLoop-24      5322812               222.7 ns/op
BenchmarkForLoop-24      5487280               220.9 ns/op
BenchmarkForLoop-24      5446502               219.0 ns/op
BenchmarkForLoop-24      5357581               220.5 ns/op
BenchmarkForLoop-24      5469464               220.2 ns/op
PASS
ok      loop-1.21       7.104s
```

1.22rc1:

```bash
$ go test -bench=. -count 5
goos: linux
goarch: amd64
pkg: loop-1.22
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkForLoop-24      2769760               430.8 ns/op
BenchmarkForLoop-24      2747745               433.6 ns/op
BenchmarkForLoop-24      2767976               433.0 ns/op
BenchmarkForLoop-24      2731050               439.8 ns/op
BenchmarkForLoop-24      2723104               445.4 ns/op
PASS
ok      loop-1.22       8.212s
```

# Comment

Exactly 2x difference is interesting, maybe this is not a coincidence.
