#include <stdio.h>

int strend(char* s, char* t);
int strlen2(char* s);

int main() {
    char* a = "superman";
    char* b = "man";
    char* c = "supermbn";
    char* d = "mbn";

    int res = -1;

    res = strend(a, b); // 1
    printf("res: %d\n", res);
    res = strend(a, c); // 0
    printf("res: %d\n", res);
    res = strend(b, d); // 0
    printf("res: %d\n", res);
    res = strend(a, d); // 0
    printf("res: %d\n", res);
    res = strend(c, d); // 1
    printf("res: %d\n", res);

    // edge case
    res = strend(b, a); // 0
    printf("res: %d\n", res);
}

int strlen2(char* s) {
    int l = 0;
    while (*s++ != '\0')
      l++;
    return l;
}

int strend(char* s, char* t) {
    // handle edge case where length t > length s
    if (strlen2(t) > strlen2(s))
        return 0;

    // get length of s
    int ls = strlen2(s);
    int lt = strlen2(t);

    // get to the end of s and subtract t
    s = s + ls - lt;

    // check through all characters
    for (; *s != '\0'; s++, t++) {
        if (*s != *t)
            return 0;
    }

    return 1;
}
