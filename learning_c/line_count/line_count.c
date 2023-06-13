#include <stdio.h>

void count_blanks_tabs_newlines() {
    int c, count;
    count = 0;

    while (c != EOF) {
        printf("current char:");
        putchar(c);
        printf("\ncurrent count: %d\n", count);

        c = getchar();
        if (c == ' ' || c == '\t' || c == '\n') {
            count++;
        }
    }
}

int main() {
  count_blanks_tabs_newlines();
}
