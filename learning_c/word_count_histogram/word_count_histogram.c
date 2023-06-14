#include <stdio.h>
#include <stdbool.h>

// arbitrary constant value
//
// apparently "Pneumonoultramicroscopicsilicovolcanoconiosis" is the longest
// english word
//
// also assuming english input
#define MAXLENGTH 50;

void tally_word_count(int count[]);
bool is_special_char(int c);

int main() {
    // TODO: why is this necessary?
    // cannot use symblic constant to init array
    const int max_length = MAXLENGTH;
    // TODO: learn why C arrays initialize to zero globally, but not locally
    int count_arr[max_length] = {0};

    tally_word_count(count_arr);

    printf("************************\n");
    printf("WORD LENGTH HISTOGRAM\n");
    printf("************************\n");
    int i, j;

    // for histogram label, since no length 0 words)
    printf("0\n");

    for (i = 1; i < max_length; i++) {
        // print labels
        if (i % 5 == 0) {
            printf("%d", i);
        } else {
            printf(" ");
        }
        printf("\t");

        // generate histogram values
        int word_count;
        word_count = count_arr[i];
        for (j = 0; j < word_count; j++) {
            printf("*");
        }
        printf("\n");
    }

    printf("************************\n");
    printf("END OF HISTOGRAM\n");
    printf("************************\n");

}

bool is_special_char(int c) {
  return c == ' ' || c == '\n' || c == '\t';
}

void tally_word_count(int count[]) {
    int c, p;
    int char_count;

    char_count = 0;

    c = getchar();

    while (c != EOF) {
        // if c is a space or a newline, it's a new word
        if (!is_special_char(c) && is_special_char(p)) {
            char_count = 1;
        } else if (!is_special_char(c) && !is_special_char(p)) {
            char_count++;
        } else {
            count[char_count]++;
        }

        p = c;
        c = getchar();
    }
}
