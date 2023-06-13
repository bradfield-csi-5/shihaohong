#include <stdio.h>

main() {
    int c;
    c = getchar();

    /*
    Test EOF char value (I see "-1")
    */
    printf("eof char: %d\n", EOF);

    /*
    Print keyboard char stream
    */
    while (c != EOF) {
        printf("char:");
        putchar(c);
        printf("\t intval:%d\n", c);
        c = getchar();
    }
}
