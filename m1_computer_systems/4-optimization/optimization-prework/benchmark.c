#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "vec/dotproduct.c"

#define ELEMENTS 1000000

void benchmark(long n, data_t (*f)(vec_ptr, vec_ptr), char* func_name) {
  clock_t test_start, test_end;
  double clocks_elapsed, time_elapsed;
  long i;
  vec_ptr u = new_vec(n);
  vec_ptr v = new_vec(n);

  for (i = 0; i < ELEMENTS; i++) {
    set_vec_element(u, i, i + 1);
    set_vec_element(v, i, i + 1);
  }

  test_start = clock();
  (*f)(u, v);
  test_end = clock();

  clocks_elapsed = test_end - test_start;
  time_elapsed = clocks_elapsed / CLOCKS_PER_SEC;

  printf("func: %s -> %.2fns per element (total time: %.2fms)\n", func_name, time_elapsed * 1e9 / n, time_elapsed * 1e3);
}

// sample use
int main() {
    benchmark(ELEMENTS, dotproduct, "dotproduct_raw");
    benchmark(ELEMENTS, dotproduct_reduce_len_call, "dotproduct_reduce_len_call");
    benchmark(ELEMENTS, dotproduct_reduce_proc_call, "dotproduct_reduce_proc_call");
    benchmark(ELEMENTS, dotproduct_unrolled_2_1, "dotproduct_unrolled_2_1");
    benchmark(ELEMENTS, dotproduct_unrolled_2_2, "dotproduct_unrolled_2_2");
    benchmark(ELEMENTS, dotproduct_unrolled_4_4, "dotproduct_unrolled_4_4");
    benchmark(ELEMENTS, dotproduct_unrolled_6_6, "dotproduct_unrolled_6_6");
    benchmark(ELEMENTS, dotproduct_unrolled_8_8, "dotproduct_unrolled_8_8");
    benchmark(ELEMENTS, dotproduct_unrolled_10_10, "dotproduct_unrolled_10_10");
}
