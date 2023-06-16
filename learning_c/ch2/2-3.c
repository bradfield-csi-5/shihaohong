#include <stdio.h>

int htoi(char s[]);

int main() {
    int i;
    i = htoi("555"); // 1365
    printf("val of i for \"555\": %d\n", i);
    i = htoi("fff"); // 4095
    printf("val of i for \"fff\": %d\n", i);
    i = htoi("abc"); // 2748
    printf("val of i for \"abd\": %d\n", i);
    i = htoi("AbC"); // 2748
    printf("val of i for \"AbC\": %d\n", i);
    i = htoi("a34"); // 2612
    printf("val of i for \"a34\": %d\n", i);
}

// TODO: consider optional 0x or 0X
int htoi(char s[]) {
    int l;
    l = 0;

    // get char len
    char c;
    c = s[l];
    while (c != '\0') {
        l++;
        c = s[l];
    }

    // compute hex to dec
    int res;
    int n;
    res = 0;
    n = 1;
    char h;
    int i;
    for (i = l - 1; i >= 0; i--) {
        // each char, mult n by 16 for each hex up
        h = s[i];

        if (h >= '0' && h <= '9') {
            res = res + (n * (h - '0'));
        } else if (h >= 'a' && h <= 'f') {
            res = res + (n * (h - 'a' + 10));
        } else if (h >= 'A' && h <= 'F') {
            res = res + (n * (h - 'A' + 10));
        }

        n = n * 16;
    }

    return res;
}
