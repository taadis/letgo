#include <stdio.h>

int main(int argc, const char * argv[]) {
    int i = 0;

    while(i++ < 4) {
        printf("while: i=%d\n", i);
    }
    printf("After while loop, i=%d,\n\n", i);

    for(i = 0; i < 4; i++) {
        printf("first for: i=%d\n", i);
    }
    printf("After first for loop, i=%d\n\n", i);

    for(i = 1; i <= 4; i++) {
        printf("second for: i=%d\n", i);
    }
    printf("After second for loop, i=%d\n", i);

    return 0;
}