#include <stdio.h>

int htoi(char s[]);

int main() {
    int i;
    i = htoi("555"); // 1365
    printf("val of i: %d", i);
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
    int i;
    int res;
    int n;
    n = 1;
    char h;
    for (i = l - 1; i >= 0; i--) {
        // each char, mult n by 16 for each hex up
        h = s[i];
        printf("val: %c\n", h);
        printf("i: %d\n", i);
        printf("n: %d\n", n);

        // if 0-9
        if (h >= '0' && h <= '9') {
            res = res + (n * (h - '0'));
        } else if (h >= 'a' && h <= 'f') {
            // TODO: finish A-F implementations
            ;
        } // TODO: ignore illegal characters for now

        n = n * 16;

        printf("result: %d\n", i);
    }

    return res;
}
