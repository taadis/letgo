#include <stdio.h>

int main (int argc, const char * argv[]) {
    int number;
    int sum = 0;
    for (int number = 1; number <= 10; number++) {
        sum += number;
    }
    printf("The sum of the numbers from 1 to 10 is %d.\n", sum);
    return 0;
}