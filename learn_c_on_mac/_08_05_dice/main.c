#include <stdio.h>
#include <time.h>
#include <stdlib.h>

int RollOne(void);
void PrintRolls(int rolls[]);
void PrintX(int howMany);

int main(int argc, const char * argv[]) {
    int rolls[13];
    int twoDice;
    int i;

    srand(clock());
    for(i = 0; i <= 12; i++) {
        rolls[i] = 0;
    }
    for(i = 1; i <= 1000; i++) {
        twoDice = RollOne() + RollOne();
        ++rolls[twoDice];
    }
    PrintRolls(rolls);

    return 0;
}

int RollOne(void) {
    return (rand() % 6) + 1;
}

void PrintRolls(int rolls[]) {
    for(int i = 2; i <= 12; i++) {
        printf("%2d (%3d): ", i, rolls[i]);
        PrintX(rolls[i] / 10);
        printf("\n");
   }
}

void PrintX(int howMany) {
    for(int i = 1; i <= howMany; i++) {
        printf("x");
    }
}

