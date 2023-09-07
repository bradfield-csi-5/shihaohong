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

### DollarAmount without time (not very effective)
```sh
BenchmarkMetrics/Average_payment-8     	    1218	    944587 ns/op
```

### Calculate dollar vs cents separately (slightly effective)
```sh
BenchmarkMetrics/Average_payment-8     	    1284	    924113 ns/op
```

### Using uint32
```sh
BenchmarkMetrics/Average_payment-8     	    1786	    667641 ns/op
```

### Using cents
```sh
BenchmarkMetrics/Average_payment-8     	    3050	    400566 ns/op
```

## StdDevPaymentAmount
### Initial Results
```sh
BenchmarkMetrics/Payment_stddev-8      	      32	  37313669 ns/op
```

### Using Payments slice -- AveragePaymentAmount
```sh
BenchmarkMetrics/Payment_stddev-8      	      56	  20556248 ns/op
```

### Loop unrolling (4x4) -- AveragePaymentAmount
```sh
BenchmarkMetrics/Payment_stddev-8      	      56	  19661426 ns/op
```

### DollarAmount without time (mildly effective here) -- AveragePaymentAmount
```sh
BenchmarkMetrics/Payment_stddev-8      	      58	  19286717 ns/op
```

### Optimized stddev calculation
```sh
BenchmarkMetrics/Payment_stddev-8      	     781	   1602426 ns/op
```

### Using uint32
```sh
BenchmarkMetrics/Payment_stddev-8      	    1215	    979863 ns/op
```

### Using cents
```sh
BenchmarkMetrics/Payment_stddev-8      	    1252	    956452 ns/op
```

### Loop unrolling (2x2)
```sh
BenchmarkMetrics/Payment_stddev-8      	    1930	    608944 ns/op
```

