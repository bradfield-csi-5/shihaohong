## Benchmarking Dot Product

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.74s to run 20000 tests (137019.15ns per test)
func: dotproduct_reduce_len_call -> 2.16s to run 20000 tests (108060.80ns per test)
func: dotproduct_reduce_proc_call -> 0.46s to run 20000 tests (22903.35ns per test)
func: dotproduct_unrolled_2_1 -> 0.40s to run 20000 tests (19910.80ns per test)
func: dotproduct_unrolled_2_2 -> 0.40s to run 20000 tests (20031.70ns per test)

optimization-prework shihao$ cc -O1 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.74s to run 20000 tests (87071.55ns per test)
func: dotproduct_reduce_len_call -> 1.17s to run 20000 tests (58582.60ns per test)
func: dotproduct_reduce_proc_call -> 0.22s to run 20000 tests (10815.55ns per test)
func: dotproduct_unrolled_2_1 -> 0.16s to run 20000 tests (7993.35ns per test)
func: dotproduct_unrolled_2_2 -> 0.13s to run 20000 tests (6527.70ns per test)

optimization-prework shihao$ cc -O2 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.76s to run 20000 tests (87977.00ns per test)
func: dotproduct_reduce_len_call -> 1.19s to run 20000 tests (59407.75ns per test)
func: dotproduct_reduce_proc_call -> 0.24s to run 20000 tests (12204.65ns per test)
func: dotproduct_unrolled_2_1 -> 0.15s to run 20000 tests (7572.80ns per test)
func: dotproduct_unrolled_2_2 -> 0.13s to run 20000 tests (6281.50ns per test)

optimization-prework shihao$ cc -O3 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.74s to run 20000 tests (86951.80ns per test)
func: dotproduct_reduce_len_call -> 1.17s to run 20000 tests (58664.10ns per test)
func: dotproduct_reduce_proc_call -> 0.24s to run 20000 tests (12126.10ns per test)
func: dotproduct_unrolled_2_1 -> 0.15s to run 20000 tests (7624.30ns per test)
func: dotproduct_unrolled_2_2 -> 0.13s to run 20000 tests (6283.35ns per test)
```
