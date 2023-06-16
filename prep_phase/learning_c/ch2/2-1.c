#include <stdio.h>
#include <limits.h>

// had to look up the source code for limits.h
// source code link: https://www.gnu.org/software/libc/sources.html
// alternatively, cmd+click limits.h once C extensions installed for IDE.
int main() {
    // char
    // %d is for signed decimal int
    printf("char lowest val: %d\n", CHAR_MIN);
    printf("char lowest val: %d\n", CHAR_MAX);
    printf("unsigned char range: %d\n", UCHAR_MAX);

    // short
    // %hi for signed short
    // %hu for unsigned short
    printf("short lowest val: %hi\n", SHRT_MIN);
    printf("short lowest val: %hi\n", SHRT_MAX);
    printf("unsigned short range: %hu\n", USHRT_MAX);

    // int
    // %i for signed int. %i will take base value as entered.
    // %u for unsigned int
    printf("int lowest val: %i\n", INT_MIN);
    printf("int lowest val: %i\n", INT_MAX);
    printf("unsigned int range: %u\n", UINT_MAX);

    // long
    // %ld for long
    // %lu is for unsigned long values
    printf("long lowest val: %ld\n", LONG_MIN);
    printf("long lowest val: %ld\n", LONG_MAX);
    printf("unsigned long range: %lu\n", ULONG_MAX);
}
