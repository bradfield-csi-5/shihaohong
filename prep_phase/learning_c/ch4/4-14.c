#include <stdio.h>

// swap fn cannot be declared after #define, else it will
// be replaced by the macro and the program won't compile.
void swap(int x, int y);
void swap_ptr(int* x, int* y);

#define swap(t,x,y) { t tmp; tmp = x; x = y; y = tmp; }

int main() {
    int a, b;
    a = 1;
    b = 2;

    printf("==== USING A MACRO ====\n");
    printf("before --> a: %d, b: %d\n", a, b);
    // works because it's not a function. it's simply a macro.
    swap(int, a, b);
    printf("after --> a: %d, b: %d\n", a, b);

    // reset
    a = 1;
    b = 2;
    #undef swap
    printf("==== USING A FN (as expected, doesnt work)====\n");
    printf("before --> a: %d, b: %d\n", a, b);
    swap(a, b);
    printf("after --> a: %d, b: %d\n", a, b);

    printf("==== USING PTR FN (as expected, works)====\n");
    printf("before --> a: %d, b: %d\n", a, b);
    swap_ptr(&a, &b);
    printf("after --> a: %d, b: %d\n", a, b);
}

// doesnt work because the variables are copied (pass by value)
void swap(int x, int y) {
    int tmp;
    tmp = x;
    x = y;
    y = tmp;
}

// works because fn takes pointers and this fn updates
// the values in each address directly
// x and y are pointer copies
void swap_ptr(int* x, int* y) {
    int tmp = *x;
    *x = *y;
    *y = tmp;
}
