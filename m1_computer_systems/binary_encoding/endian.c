#include <stdio.h>

int main()
{
	int x = 1;
	char *cp = &x;
	if (*cp == 1)
		printf("Little endian");
	else
		printf("Big endian");
	return 0;
}
