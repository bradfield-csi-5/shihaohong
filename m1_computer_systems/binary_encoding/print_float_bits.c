#include <stdio.h>
#include <stdint.h>

int main() {
    float f = 1;
    uint32_t *ptr = &f;
    printf("0x%x\n", *ptr);
}
