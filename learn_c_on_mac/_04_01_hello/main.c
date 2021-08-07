#include <stdio.h>

void SayHello(void);

int main(int agrc, const char * argv[]) {
    SayHello();
    return 0;
}

void SayHello(void) {
    printf("Hello, world!\n");
}