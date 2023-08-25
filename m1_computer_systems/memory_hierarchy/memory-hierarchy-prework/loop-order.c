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

<-- THESE ARE THE SMALLER CORES -->
memory-hierarchy-prework shihao$ sysctl -a | grep hw.l
hw.l1icachesize: 131072
hw.l1dcachesize: 65536
hw.l2cachesize: 4194304

<-- LOOKING AT WIKIPEDIA -->
L1 cache of high-performance cores:
l1i: 192KB
l1d: 128KB
L2 cache of high-performance cores:
shared 12MB

L1 cache of efficient cores:
l1d stores data, l1i stores instructions
l1d: ~65.5KB
l2: ~4.2MB

l1 cache latency: 3-4 cycles
l2 cache latency: 18 cycles
src: https://www.7-cpu.com/cpu/Apple_M1.html

matrix: 4000 * 4000 * 4 bytes (int)
matrix size: ~64MB
row size: 16KB

Option 1:
assuming efficient core:
l1 cache is reset at least ~1000 times
l2 cache is reset at least 16 times

assuming high-perf core:
l1 cache is reset at least ~500 times
l2 cache is reset at least 5 times

Option 2:
assuming efficient core:
l1 cache is reset at least 4,000,000 times
l2 cache is reset at least ~60k times

assuming high-perf core:
l1 cache is reset at least 2,000,000 times
l2 cache is reset at least ~30k times

Difference in clock cycles (using worst case):
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

Instrument results:
-- Random observation, the work is split across multiple cores

Option 1:
L1d miss, L2 miss, elements
212k, 26k, one element
339k, 78k, 4kx4k elements
DIFF
127k, 52k

Option 2:
176k, 22k, one element
200k, 14.7M, 4kx4k elements
DIFF
24k, ~14M
*/

int x[1][1];

void option_one() {
  int i, j;
  for (i = 0; i < 1; i++) {
    for (j = 0; j < 1; j++) {
      x[i][j] = i + j;
    }
  }
}

void option_two() {
  int i, j;
  for (i = 0; i < 1; i++) {
    for (j = 0; j < 1; j++) {
      x[j][i] = i + j;
    }
  }
}

int main() {
  // option_one();
  option_two();
  // int v = x[0][0];

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
