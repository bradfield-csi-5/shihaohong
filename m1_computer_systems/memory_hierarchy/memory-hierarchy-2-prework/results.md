## AverageAge
### Initial Results

```sh
BenchmarkMetrics/Average_age-8         	     970	   1241247 ns/op
```

### Use slices instead of maps

- (cache technique) Removes iteration over bigger struct data type (can fit more into one cache line)

```sh
BenchmarkMetrics/Average_age-8         	   12778	     93183 ns/op
```

### Loop unrolling (4x4)

- (not cache technique) Takes advantage of CPU's functional units

```sh
BenchmarkMetrics/Average_age-8         	   44096	     27215 ns/op
```

## AveragePaymentAmount
### Initial Results
```sh
BenchmarkMetrics/Average_payment-8     	      60	  18061158 ns/op
```

### Using Payments slice
```sh
BenchmarkMetrics/Average_payment-8     	     913	   1288673 ns/op
```

### Loop unrolling (4x4)
```sh
BenchmarkMetrics/Average_payment-8     	    1267	    942685 ns/op
```

### Payment slice instead of User map

## StdDevPaymentAmount
### Initial Results
```sh
BenchmarkMetrics/Payment_stddev-8      	      32	  37313669 ns/op
```
