// #include <stdio.h>
// #include <time.h>

  /*

Two different ways to loop over an array of arrays.

Spotted at:
http://stackoverflow.com/questions/9936132/why-does-the-order-of-the-loops-affect-performance-when-iterating-over-a-2d-arra

*/

/*
Should be faster than option_two
They should execute the same number of instructions,
but due to spatial locality, option_one should run
faster due to its cache utilization.

On Apple M1 Pro:
memory-hierarchy-prework shihao$ sysctl -a | grep hw.l
hw.l1icachesize: 131072
hw.l1dcachesize: 65536
hw.l2cachesize: 4194304

l1d stores data, l1i stores instructions
l1d: ~65.5KB
l2: ~4.2MB
l3: 8MB (shared for all clusters)

l1 cache latency: 3-4 cycles
l2 cache latency: 18 cycles
src: https://www.7-cpu.com/cpu/Apple_M1.html

matrix: 4000 * 4000 * 4 bytes (int)
matrix size: ~64MB
row size: 16KB

Option 1:
l1 cache is reset at least ~1000 times
l2 cache is reset at least 16 times

Option 2:
l1 cache can store about 4 rows, so it'll be utilized every 4 reads.
l1 cache is reset at least 4,000,000 times
l2 cache can store about ~260 rows, so it'll be utilized every 260 reads
l2 cache is reset at least 61,538 times

Difference in clock cycles:
4M * 3-4 = 12M cycles
61K * 18 = 1M cycles
13M clock cycles

2-3 GHz core,
13M / 2GHz = 6.5ms
13M / 3GHz = 4.3ms
So, at least a few ms slower.

Using benchmarking code, can see it's about an order of magnitude slower than predicted:
Time taken (option_one): 0.028s
Time taken (option_one): 0.083s
Diff: 0.055s = 55ms
*/

void option_one() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[i][j] = i + j;
    }
  }
}

void option_two() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[j][i] = i + j;
    }
  }
}

int main() {
  // option_one();
  option_two();

  // clock_t start, stop;
  // double time;
  // start = clock();
  // option_one();
  // stop = clock();
  // time = (stop - start) / (double)CLOCKS_PER_SEC;
  // printf("Time taken (option_one): %.3fs\n", time);

  // start = clock();
  // option_two();
  // stop = clock();
  // time = (stop - start) / (double)CLOCKS_PER_SEC;
  // printf("Time taken (option_two): %.3fs\n", time);

  return 0;
}
