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
    // gets only one operator/number
    i = getop(ca);
    // print the int value of the next char
    i = printf("resulting val: %d\n", i);
    // print the string stored
    printf("resulting string: %s\n", ca);
}

/* getop: get next operator or numeric operand */
int getop(char s[])
{
    int i, c;

    // skip whitespace
    while ((s[0] = c = getch()) == ' ' || c == '\t')
        ;
    s[1] = '\0';

    // single char operator
    if (c == '%' || c == '+' || c == '/' || c == '\n' || c == '*') {
        return c;
    }

    i = 0;
    // '-' is a special case since it can signify negative number
    if (c == '-') {
        // need to determine if neg number or subtraction
        if ((c = getch()) == ' ' || c == '\t' || c == EOF || c == '\n') {
            // subtraction
            ungetch(c);
            return '-';
        } else // neg number
            ungetch(c);
    }

    if (!isdigit(c) && c != '.' && c != '-') {
        // variable or command so get entire thing
        while ((s[++i] = c = getch()) != ' ' && c != EOF && c != '\n' && c != '\t')
            ;
        s[i] = '\0';
        ungetch(c);
        return COMMAND;
    }

    if (isdigit(c))     // collect integer part
        while (isdigit(s[++i] = c = getch()))
            ;
    if (c == '.') {     // collect fraction part
        while (isdigit(s[++i] = c = getch()))
            ;
    }
    s[i] = '\0';
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
