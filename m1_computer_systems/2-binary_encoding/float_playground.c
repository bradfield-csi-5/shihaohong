#include <stdio.h>

int main() {
  float x = 0;
  float y = 100000000;

  for (int i = 0; i < 100000000; i++) {
    x += 1;
    if (i % 100 == 0) {
        printf("calculating === i : %i\n", i);
        printf("%f\n", x);
    }
  }

  printf("%f\n", x);
  printf("%f\n", y);
}
