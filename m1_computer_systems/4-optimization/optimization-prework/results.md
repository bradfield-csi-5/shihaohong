## Benchmarking Dot Product

Longs

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.77s to run 20000 tests (138672.05ns per test)
func: dotproduct_reduce_len_call -> 2.18s to run 20000 tests (109073.20ns per test)
func: dotproduct_reduce_proc_call -> 0.46s to run 20000 tests (22836.35ns per test)
func: dotproduct_unrolled_2_1 -> 0.40s to run 20000 tests (19925.70ns per test)
func: dotproduct_unrolled_2_2 -> 0.40s to run 20000 tests (19835.95ns per test)
func: dotproduct_unrolled_6_6 -> 0.37s to run 20000 tests (18485.30ns per test)
func: dotproduct_unrolled_8_8 -> 0.36s to run 20000 tests (17906.95ns per test)
func: dotproduct_unrolled_10_10 -> 0.36s to run 20000 tests (17931.70ns per test)

optimization-prework shihao$ cc -O1 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.76s to run 20000 tests (88211.95ns per test)
func: dotproduct_reduce_len_call -> 1.19s to run 20000 tests (59638.20ns per test)
func: dotproduct_reduce_proc_call -> 0.22s to run 20000 tests (11158.85ns per test)
func: dotproduct_unrolled_2_1 -> 0.16s to run 20000 tests (8151.95ns per test)
func: dotproduct_unrolled_2_2 -> 0.13s to run 20000 tests (6711.15ns per test)
func: dotproduct_unrolled_6_6 -> 0.12s to run 20000 tests (6114.25ns per test)
func: dotproduct_unrolled_8_8 -> 0.13s to run 20000 tests (6511.00ns per test)
func: dotproduct_unrolled_10_10 -> 0.13s to run 20000 tests (6447.75ns per test)

optimization-prework shihao$ cc -O2 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.76s to run 20000 tests (88096.55ns per test)
func: dotproduct_reduce_len_call -> 1.18s to run 20000 tests (59205.00ns per test)
func: dotproduct_reduce_proc_call -> 0.24s to run 20000 tests (12068.45ns per test)
func: dotproduct_unrolled_2_1 -> 0.15s to run 20000 tests (7497.15ns per test)
func: dotproduct_unrolled_2_2 -> 0.12s to run 20000 tests (6064.50ns per test)
func: dotproduct_unrolled_6_6 -> 0.12s to run 20000 tests (6146.30ns per test)
func: dotproduct_unrolled_8_8 -> 0.13s to run 20000 tests (6546.40ns per test)
func: dotproduct_unrolled_10_10 -> 0.13s to run 20000 tests (6694.90ns per test)

optimization-prework shihao$ cc -O3 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.74s to run 20000 tests (86872.60ns per test)
func: dotproduct_reduce_len_call -> 1.17s to run 20000 tests (58374.15ns per test)
func: dotproduct_reduce_proc_call -> 0.25s to run 20000 tests (12262.90ns per test)
func: dotproduct_unrolled_2_1 -> 0.15s to run 20000 tests (7501.55ns per test)
func: dotproduct_unrolled_2_2 -> 0.12s to run 20000 tests (6120.40ns per test)
func: dotproduct_unrolled_6_6 -> 0.12s to run 20000 tests (6170.20ns per test)
func: dotproduct_unrolled_8_8 -> 0.13s to run 20000 tests (6250.65ns per test)
func: dotproduct_unrolled_10_10 -> 0.13s to run 20000 tests (6579.00ns per test)
```

Doubles

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.79s to run 20000 tests (139521.35ns per test)
func: dotproduct_reduce_len_call -> 2.20s to run 20000 tests (109877.95ns per test)
func: dotproduct_reduce_proc_call -> 1.37s to run 20000 tests (68525.55ns per test)
func: dotproduct_unrolled_2_1 -> 1.37s to run 20000 tests (68522.50ns per test)
func: dotproduct_unrolled_2_2 -> 0.69s to run 20000 tests (34397.80ns per test)
func: dotproduct_unrolled_6_6 -> 0.38s to run 20000 tests (18811.85ns per test)
func: dotproduct_unrolled_8_8 -> 0.36s to run 20000 tests (17854.80ns per test)
func: dotproduct_unrolled_10_10 -> 0.35s to run 20000 tests (17508.20ns per test)

optimization-prework shihao$ cc -O1 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.83s to run 20000 tests (91330.05ns per test)
func: dotproduct_reduce_len_call -> 2.02s to run 20000 tests (100830.05ns per test)
func: dotproduct_reduce_proc_call -> 0.37s to run 20000 tests (18608.85ns per test)
func: dotproduct_unrolled_2_1 -> 0.37s to run 20000 tests (18588.00ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9335.05ns per test)
func: dotproduct_unrolled_6_6 -> 0.11s to run 20000 tests (5286.40ns per test)
func: dotproduct_unrolled_8_8 -> 0.11s to run 20000 tests (5298.05ns per test)
func: dotproduct_unrolled_10_10 -> 0.11s to run 20000 tests (5293.30ns per test)

optimization-prework shihao$ cc -O2 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.85s to run 20000 tests (92736.60ns per test)
func: dotproduct_reduce_len_call -> 2.03s to run 20000 tests (101503.10ns per test)
func: dotproduct_reduce_proc_call -> 0.37s to run 20000 tests (18593.55ns per test)
func: dotproduct_unrolled_2_1 -> 0.37s to run 20000 tests (18592.55ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9314.55ns per test)
func: dotproduct_unrolled_6_6 -> 0.06s to run 20000 tests (3169.20ns per test)
func: dotproduct_unrolled_8_8 -> 0.05s to run 20000 tests (2380.40ns per test)
func: dotproduct_unrolled_10_10 -> 0.05s to run 20000 tests (2289.80ns per test)

optimization-prework shihao$ cc -O3 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.84s to run 20000 tests (92207.45ns per test)
func: dotproduct_reduce_len_call -> 2.06s to run 20000 tests (103049.45ns per test)
func: dotproduct_reduce_proc_call -> 0.37s to run 20000 tests (18596.35ns per test)
func: dotproduct_unrolled_2_1 -> 0.37s to run 20000 tests (18639.10ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9353.65ns per test)
func: dotproduct_unrolled_6_6 -> 0.06s to run 20000 tests (3176.50ns per test)
func: dotproduct_unrolled_8_8 -> 0.05s to run 20000 tests (2379.75ns per test)
func: dotproduct_unrolled_10_10 -> 0.05s to run 20000 tests (2684.05ns per test)
```
