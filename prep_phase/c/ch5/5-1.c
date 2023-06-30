#include <stdio.h>
#include <ctype.h>

#define BUFSIZE 100

char buf[BUFSIZE];  // buffer for ungetch
int bufp = 0;       // next free position in buffer

int getch(void);
void ungetch(int);
int getint(int *pn);

/*
output:
a val before: -1
+asdf1234-
a val after: -1
check buf: a+ // expected, because buffer should be in reverse order.
check next input val: s
*/
int main() {
    int a = -1;
    printf("a val before: %d\n", a);
    getint(&a);
    // check that nothing was pushed to [a].
    printf("a val after: %d\n", a);
    // check buffer value is correct
    printf("check buf: %s\n", buf);
    // verify value of next char in input stream
    int c = getchar();
    printf("check next input val: %c\n", c);
}

int getint(int *pn) {
    int c, sign;

    while (isspace(c = getch()))
        ;

    if (!isdigit(c) && c != EOF && c != '+' && c != '-') {
        ungetch(c);
        return 0;
    }

    sign = (c == '-') ? -1 : 1;
    if (c == '+' || c == '-')
        c = getch();

    if (!isdigit(c)) {
        ungetch(c);
        sign == 1 ? ungetch('+') : ungetch('-');
        return 0;
    }

    for (*pn = 0; isdigit(c); c = getch())
        *pn = 10 * *pn + (c- '0');
    *pn *= sign;
    if (c != EOF)
        ungetch(c);
    return c;
}

/* getch: get a (possibly pushed back) character */
int getch(void)
{
    return (bufp > 0) ? buf[--bufp] : getchar();
}

/* push char back on input */
void ungetch(int c)
{
    if (bufp >= BUFSIZE)
        printf("ungetch: too many characters\n");
    else if (c != EOF)
        buf[bufp++] = c;
}
