#include <stdio.h>
#include <stdbool.h>

void print_one_word_per_line();
bool is_special_char(int c);

int main() {
    print_one_word_per_line();
}

bool is_special_char(int c) {
  return c == ' ' || c == '\n' || c == '\t';
}

void print_one_word_per_line() {
    // while not EOF
    int c, p;

    c = getchar();

    while (c != EOF) {
        // if c is a space or a newline, it's a new word
        if (is_special_char(c)) {
            // do not create a newline if the previous character is a non-word char
            if (!is_special_char(p)) {
                printf("\n");
            }
        } else {
            putchar(c);
        }

        p = c;
        c = getchar();
    }
}
