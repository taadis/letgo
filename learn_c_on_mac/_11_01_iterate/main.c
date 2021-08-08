#include <stdio.h>

int main(int argc, const char * argv[]) {
    int i;
    int num;
    long fac;

    num = 5;
    fac = 1;

    for(i = 1; i <= num; i++) {
        fac *= i;
    }

    printf("%d factorial is %ld.", num, fac);

    return 0;
}

