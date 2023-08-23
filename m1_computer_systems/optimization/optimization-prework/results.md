## Benchmarking Dot Product

Longs

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 3.53s to run 20000 tests (176684.55ns per test)
func: dotproduct_reduce_len_call -> 2.92s to run 20000 tests (145761.05ns per test)
func: dotproduct_reduce_proc_call -> 0.61s to run 20000 tests (30405.05ns per test)
func: dotproduct_unrolled_2_1 -> 0.53s to run 20000 tests (26314.95ns per test)
func: dotproduct_unrolled_2_2 -> 0.52s to run 20000 tests (26088.00ns per test)
func: dotproduct_unrolled_6_6 -> 0.49s to run 20000 tests (24350.05ns per test)

optimization-prework shihao$ cc -O1 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.29s to run 20000 tests (114347.25ns per test)
func: dotproduct_reduce_len_call -> 1.54s to run 20000 tests (76895.55ns per test)
func: dotproduct_reduce_proc_call -> 0.28s to run 20000 tests (14112.90ns per test)
func: dotproduct_unrolled_2_1 -> 0.21s to run 20000 tests (10436.40ns per test)
func: dotproduct_unrolled_2_2 -> 0.17s to run 20000 tests (8504.50ns per test)
func: dotproduct_unrolled_6_6 -> 0.16s to run 20000 tests (7955.20ns per test)

optimization-prework shihao$ cc -O2 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.29s to run 20000 tests (114482.90ns per test)
func: dotproduct_reduce_len_call -> 1.54s to run 20000 tests (77159.65ns per test)
func: dotproduct_reduce_proc_call -> 0.32s to run 20000 tests (15882.70ns per test)
func: dotproduct_unrolled_2_1 -> 0.20s to run 20000 tests (9849.10ns per test)
func: dotproduct_unrolled_2_2 -> 0.16s to run 20000 tests (7958.00ns per test)
func: dotproduct_unrolled_6_6 -> 0.16s to run 20000 tests (7949.70ns per test)

optimization-prework shihao$ cc -O3 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.28s to run 20000 tests (114086.20ns per test)
func: dotproduct_reduce_len_call -> 1.53s to run 20000 tests (76620.75ns per test)
func: dotproduct_reduce_proc_call -> 0.32s to run 20000 tests (15875.70ns per test)
func: dotproduct_unrolled_2_1 -> 0.20s to run 20000 tests (9856.10ns per test)
func: dotproduct_unrolled_2_2 -> 0.16s to run 20000 tests (7971.75ns per test)
func: dotproduct_unrolled_6_6 -> 0.16s to run 20000 tests (7943.90ns per test)
```

Doubles

```sh
optimization-prework shihao$ cc benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 2.74s to run 20000 tests (137134.10ns per test)
func: dotproduct_reduce_len_call -> 2.26s to run 20000 tests (112990.85ns per test)
func: dotproduct_reduce_proc_call -> 1.40s to run 20000 tests (70175.05ns per test)
func: dotproduct_unrolled_2_1 -> 1.38s to run 20000 tests (69172.75ns per test)
func: dotproduct_unrolled_2_2 -> 0.69s to run 20000 tests (34394.30ns per test)
func: dotproduct_unrolled_6_6 -> 0.36s to run 20000 tests (18200.30ns per test)
func: dotproduct_unrolled_10_10 -> 0.58s to run 20000 tests (29154.80ns per test)

optimization-prework shihao$ cc -O1 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.86s to run 20000 tests (92999.65ns per test)
func: dotproduct_reduce_len_call -> 2.08s to run 20000 tests (103979.25ns per test)
func: dotproduct_reduce_proc_call -> 0.37s to run 20000 tests (18657.85ns per test)
func: dotproduct_unrolled_2_1 -> 0.37s to run 20000 tests (18596.40ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9338.95ns per test)
func: dotproduct_unrolled_6_6 -> 0.11s to run 20000 tests (5297.30ns per test)
func: dotproduct_unrolled_10_10 -> 0.17s to run 20000 tests (8401.90ns per test)

optimization-prework shihao$ cc -O2 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.84s to run 20000 tests (91921.05ns per test)
func: dotproduct_reduce_len_call -> 2.04s to run 20000 tests (102115.10ns per test)
func: dotproduct_reduce_proc_call -> 0.38s to run 20000 tests (18936.35ns per test)
func: dotproduct_unrolled_2_1 -> 0.38s to run 20000 tests (18961.40ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9568.25ns per test)
func: dotproduct_unrolled_6_6 -> 0.06s to run 20000 tests (3239.05ns per test)
func: dotproduct_unrolled_10_10 -> 0.13s to run 20000 tests (6388.95ns per test)

cc -O3 benchmark.c vec/vec.c && ./a.out
func: dotproduct_raw -> 1.82s to run 20000 tests (90855.25ns per test)
func: dotproduct_reduce_len_call -> 2.03s to run 20000 tests (101503.25ns per test)
func: dotproduct_reduce_proc_call -> 0.38s to run 20000 tests (18963.50ns per test)
func: dotproduct_unrolled_2_1 -> 0.38s to run 20000 tests (19119.30ns per test)
func: dotproduct_unrolled_2_2 -> 0.19s to run 20000 tests (9529.65ns per test)
func: dotproduct_unrolled_6_6 -> 0.06s to run 20000 tests (3230.40ns per test)
func: dotproduct_unrolled_10_10 -> 0.12s to run 20000 tests (5993.05ns per test)
```
