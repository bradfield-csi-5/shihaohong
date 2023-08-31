## AverageAge
### Initial Results

```sh
go test -bench=.
goos: darwin
goarch: amd64
pkg: metrics.com/m/v2
cpu: VirtualApple @ 2.50GHz
BenchmarkMetrics/Average_age-8         	     970	   1241247 ns/op
```

### remove use of sets, use slices

```sh
go test -bench=.
goos: darwin
goarch: amd64
pkg: metrics.com/m/v2
cpu: VirtualApple @ 2.50GHz
BenchmarkMetrics/Average_age-8         	   12778	     93183 ns/op
PASS
```

### Loop unrolling (4 accumulators)

```sh
go test -bench=.
goos: darwin
goarch: amd64
pkg: metrics.com/m/v2
cpu: VirtualApple @ 2.50GHz
BenchmarkMetrics/Average_age-8         	   44096	     27215 ns/op
```
