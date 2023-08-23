## Benchmarking Dot Product

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.74s to run 20000 tests (137019.15ns per test)
func: dotproduct_reduce_len_call -> 2.16s to run 20000 tests (108060.80ns per test)
func: dotproduct_reduce_proc_call -> 0.46s to run 20000 tests (22903.35ns per test)
func: dotproduct_unrolled_2_1 -> 0.40s to run 20000 tests (19910.80ns per test)
func: dotproduct_unrolled_2_2 -> 0.40s to run 20000 tests (20031.70ns per test)
```
