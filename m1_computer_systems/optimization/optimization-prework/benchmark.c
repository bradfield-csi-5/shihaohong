#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "vec/dotproduct.c"

void benchmark(long n, data_t (*f)(vec_ptr, vec_ptr), char* func_name) {
  clock_t test_start, test_end;
  double clocks_elapsed, time_elapsed;
  long i;
  vec_ptr u = new_vec(n);
  vec_ptr v = new_vec(n);

  for (i = 0; i < n; i++) {
    set_vec_element(u, i, i + 1);
    set_vec_element(v, i, i + 1);
  }

  test_start = clock();
  for (i = 0; i < n; i++) {
    (*f)(u, v);
  }
  test_end = clock();

  clocks_elapsed = test_end - test_start;
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("func: %s -> %.2fs to run %ld tests (%.2fns per test)\n", func_name, time_elapsed, n,
         time_elapsed * 1e9 / n);
}

// sample use
int main() {
    benchmark(20000, dotproduct, "dotproduct_raw");
    benchmark(20000, dotproduct_reduce_len_call, "dotproduct_reduce_len_call");
    benchmark(20000, dotproduct_reduce_proc_call, "dotproduct_reduce_proc_call");
    benchmark(20000, dotproduct_unrolled_2_1, "dotproduct_unrolled_2_1");
    benchmark(20000, dotproduct_unrolled_2_2, "dotproduct_unrolled_2_2");
    benchmark(20000, dotproduct_unrolled_4_4, "dotproduct_unrolled_4_4");
    benchmark(20000, dotproduct_unrolled_6_6, "dotproduct_unrolled_6_6");
    benchmark(20000, dotproduct_unrolled_8_8, "dotproduct_unrolled_8_8");
    benchmark(20000, dotproduct_unrolled_10_10, "dotproduct_unrolled_10_10");
}
