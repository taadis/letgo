#include <stdio.h>

void SayHello(void);

int main(int argc, const char * argv[]) {
    SayHello();
    SayHello();
    SayHello();
    return 0;
}

void SayHello(void) {
    printf("Hello, world!\n");
}