#include <stdio.h>

// Assumes little endian
// printBits(sizeof(val), &val);
void printBits(size_t const size, void const *const ptr)
{
    unsigned char *b = (unsigned char *)ptr;
    unsigned char byte;
    int i, j;

    for (i = size - 1; i >= 0; i--)
    {
        for (j = 7; j >= 0; j--)
        {
            byte = (b[i] >> j) & 1;
            printf("%u", byte);
        }
    }
    puts("");
}

int main()
{
    unsigned int const v = 15; // Round this 32-bit value to the next highest power of 2
    unsigned int r;            // Put the result here. (So v=3 -> r=4; v=8 -> r=8)

    if (v > 1)
    {
        float f = (float)v;
        unsigned int const t = 1U << ((*(unsigned int *)&f >> 23) - 0x7f);
        printf("E = %d\n", (*(unsigned int *)&f >> 23) - 0x7f);
        printf("t = %d\n", t);
        r = t << (t < v);
    }
    else
    {
        r = 1;
    }
    printf("r = %d\n", r);
}
