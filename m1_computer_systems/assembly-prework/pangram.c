#include <stdio.h>

int pangram(char*);

/*
nasm -f macho64 --prefix _ -L/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib pangram.asm \
    && gcc -arch x86_64 pangram.c pangram.o \
    && ./a.out
*/
int main() {
    printf("res: %d\n", pangram("abc"));
    printf("res: %d\n", pangram("abcdefghijklmnopqrstuvwxyz"));
    return 0;
}
