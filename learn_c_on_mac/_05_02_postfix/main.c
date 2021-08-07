#include <stdio.h>

int main(int argc, const char * argv[]) {
    int myInt = 5;
    printf("myInt ---> %d\n", myInt++);
    printf("myInt ---> %d\n", ++myInt);
    return 0;
}