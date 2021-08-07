#include <stdio.h>

int AddTheseNumbers(int num1, int num2);

int main(int argc, const char * argv[]) {
    int sum = AddTheseNumbers(5, 6);
    printf("The sum is %d.", sum);
    return 0;
}

int AddTheseNumbers(int num1, int num2){
    return (num1 + num2);
}

