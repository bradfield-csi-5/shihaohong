#include <ctype.h>
#include <stdio.h>

#define NUMBER '0'
#define COMMAND '1'
#define BUFSIZE 100
#define MAXLENGTH 100

int getch(void);
void ungetch(int);
int getop(char s[]);

char buf[BUFSIZE];  // buffer for ungetch
int bufp = 0;       // next free position in buffer

int main(){
    int i = 0;
    char ca[MAXLENGTH];
    int i2 = 0;
    char ca2[MAXLENGTH];
    int i3 = 0;
    char ca3[MAXLENGTH];
    i = getop(ca);
    printf("resulting ascii int val: %d\n", i);
    printf("resulting string: %s\n", ca);
    i2 = getop(ca2);
    printf("resulting ascii int val: %d\n", i2);
    printf("resulting string2: %s\n", ca2);
    i3 = getop(ca3);
    printf("resulting ascii int val3: %d\n", i3);
    printf("resulting string3: %s\n", ca3);
}

/* getop: get next operator or numeric operand */
int getop(char s[])
{
    int i;
    char c;

    while ((s[0] = c = getch()) == ' ' || c == '\t')
        ;
    s[1] = '\0';

    // single char operator
    if (c == '%' || c == '+' || c == '/' || c == '\n' || c == '*') {
        return c;
    }

    if (!isdigit(c) && c != '.')
        return c;
    i = 0;
    if (isdigit(c)) { // collect integer
        while (isdigit(s[++i] = c = getch()))
            ;
    }

    if (c == '.') { // collect fraction
        while (isdigit(s[++i] = c = getch()))
            ;
    }
    s[i] = '\0';
    if (c != EOF)
        ungetch(c);

    return NUMBER;
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
