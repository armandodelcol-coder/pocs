#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv) {
    if (argc < 3) {
        puts("Usage is ./add 42 7");
        return 0;
    }

    double result = 0.0;
    for (int i = 1; i < argc; i++) {
        printf("Arg pos %d is %s ... Adding with %.3g \n", i, argv[i], result);
        result += strtod(argv[i], NULL);
    }

    printf("\n\nResult is %.3g\n", result);
    return 0;
}